package gitpool_lx

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
)

//channelPool 存放连接信息
type channelPool struct {
	mu           sync.RWMutex
	conns        chan *idleConn
	factory      func() (interface{}, error)
	close        func(interface{}) error
	ping         func(interface{}) error
	idleTimeout  time.Duration
	waitTimeout  time.Duration
	maxActive    int
	openingConns int
	connReqs     []chan connReq
}

type Config struct {
	//连接池中拥有的最小连接数
	InitialCap int
	//最大并发存活连接数
	MaxCap int
	//最大空闲连接
	MaxIdle int
	//生成连接的方法
	Factory func() (interface{}, error)
	//关闭连接的方法
	Close func(interface{}) error
	//检查连接是否有效的方法
	Ping func(interface{}) error
	//连接最大空闲时间，超过该事件则将失效
	IdleTimeout time.Duration
}

// NewChannelPool 初始化连接
func NewChannelPool(p *Config) (Pool, error) {
	//   最小连接数  <= 最大空闲连接  最大并发存活连接数>=最大空闲连接  连接池中拥有的最小连接数 > 0
	if !(p.InitialCap <= p.MaxIdle && p.MaxCap >= p.MaxIdle && p.InitialCap >= 0) {
		return nil, errors.New("invalid capacity settings")
	}
	if p.Factory == nil {
		return nil, errors.New("invalid capacity settings")
	}
	if p.Close == nil {
		return nil, errors.New("invalid capacity settings")
	}

	c := &channelPool{
		conns:        make(chan *idleConn, p.MaxIdle),
		factory:      p.Factory,
		close:        p.Close,
		idleTimeout:  p.IdleTimeout,
		maxActive:    p.MaxCap,
		openingConns: p.InitialCap,
	}

	if p.Ping != nil {
		c.ping = p.Ping
	}

	for i := 0; i < p.InitialCap; i++ {
		conn, err := c.factory()
		if err != nil {
			c.Release()
			return nil, fmt.Errorf("factory is not able to fill the pool: %s", err)
		}
		c.conns <- &idleConn{conn: conn, t: time.Now()}
	}
	return c, nil
}

func (c *channelPool) Get() (interface{}, error) {
	conns := c.getConns()
	if conns == nil {
		return nil, ErrClosed
	}

	for {
		select {
		case wrapConn := <-conns:
			if wrapConn == nil {
				return nil, ErrClosed
			}
			//判断是否超时，超时则丢弃
			if timeOut := c.idleTimeout; timeOut > 0 {
				if wrapConn.t.Add(timeOut).Before(time.Now()) {
					//丢弃并关闭该连接
					_ = c.Close(wrapConn.conn)
					continue
				}
			}
			//判断是否失效，失效则丢弃，如果用户没有设定 ping 方法，就不检查
			if c.ping != nil {
				if err := c.Ping(wrapConn.conn); err != nil {
					_ = c.Close(wrapConn.conn)
					continue
				}
			}
			return wrapConn.conn, nil
		default:
			c.mu.Lock()
			log.Debugf("openConn %v %v", c.openingConns, c.maxActive)
			if c.openingConns >= c.maxActive {
				req := make(chan connReq, 1)
				c.connReqs = append(c.connReqs, req)
				c.mu.Unlock()
				ret, ok := <-req
				if !ok {
					return nil, ErrMaxActiveConnReached
				}
				if timeOut := c.idleTimeout; timeOut > 0 {
					if ret.idleConn.t.Add(timeOut).Before(time.Now()) {
						_ = c.Close(ret.idleConn.conn)
						continue
					}
				}
				return ret.idleConn.conn, nil
			}
			if c.factory == nil {
				c.mu.Unlock()
				return nil, ErrClosed
			}
			conn, err := c.factory()
			if err != nil {
				c.mu.Unlock()
				return nil, err
			}
			c.openingConns++
			c.mu.Unlock()
			return conn, nil
		}
	}
}

func (c *channelPool) Put(conn interface{}) error {
	if conn == nil {
		return errors.New("connection is nil. rejecting")
	}

	c.mu.Lock()

	if c.conns == nil {
		c.mu.Unlock()
		return c.Close(conn)
	}

	if l := len(c.connReqs); l > 0 {
		req := c.connReqs[0]
		copy(c.connReqs, c.connReqs[1:])
		c.connReqs = c.connReqs[:l-1]
		req <- connReq{
			idleConn: &idleConn{conn: conn, t: time.Now()},
		}
		c.mu.Unlock()
		return nil
	} else {
		select {
		case c.conns <- &idleConn{conn: conn, t: time.Now()}:
			c.mu.Unlock()
			return nil
		default:
			c.mu.Unlock()
			//连接池已满，直接关闭该连接
			return c.Close(conn)
		}
	}
}

// Close 关闭单条连接
func (c *channelPool) Close(conn interface{}) error {
	if conn == nil {
		return errors.New("connection is nil. rejecting")
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.close == nil {
		return nil
	}
	c.openingConns--
	return c.close(conn)
}

// Ping 检查单条连接是否有效
func (c *channelPool) Ping(conn interface{}) error {
	if conn == nil {
		return errors.New("connection is nil. rejecting")
	}
	return c.ping(conn)
}

// Release 释放连接池中所有连接
func (c *channelPool) Release() {
	c.mu.Lock()
	conns := c.conns
	c.conns = nil
	c.factory = nil
	c.ping = nil
	closeFun := c.close
	c.close = nil
	c.mu.Unlock()

	if conns == nil {
		return
	}

	close(conns)
	for wrapConn := range conns {
		closeFun(wrapConn.conn)
	}
}

func (c *channelPool) Len() int {
	return len(c.getConns())
}

// getConns 获取所有连接
func (c *channelPool) getConns() chan *idleConn {
	c.mu.Lock()
	conns := c.conns
	c.mu.Unlock()
	return conns
}

type connReq struct {
	idleConn *idleConn
}

type idleConn struct {
	conn interface{}
	t    time.Time
}
