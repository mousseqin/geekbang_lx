package kratos

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/google/uuid"
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// AppInfo is application context value.
type AppInfo interface {
	ID() string
	Name() string
	Version() string
	Metadata() map[string]string
	Endpoint() []string
}

func (a *App) ID() string {
	return a.opts.id
}

func (a *App) Name() string {
	return a.opts.name
}

func (a *App) Version() string {
	return a.opts.version
}

func (a *App) Metadata() map[string]string {
	return a.opts.metadata
}

func (a *App) Endpoint() []string {
	if a.instance == nil {
		return []string{}
	}
	return a.instance.Endpoints
}

type App struct {
	opts 	 options
	ctx      context.Context
	cancel   func()
	lk       sync.Mutex
	instance *registry.ServiceInstance
}

func (a *App) Run() error {
	instance, err := a.buildInstance()
	if err != nil {
		return err
	}
	eg, ctx := errgroup.WithContext(NewContext(a.ctx, a))
	wg := sync.WaitGroup{}

	for _,srv := range a.opts.servers {
		srv := srv // 闭包
		eg.Go(func() error {
			<- ctx.Done() // wait for stop signal
			stopCtx , cancel := context.WithTimeout(NewContext(a.opts.ctx,a),a.opts.stopTimeout)
			defer cancel()
			return srv.Stop(stopCtx)
		})
		wg.Add(1)
		eg.Go(func()error{
			wg.Done()
			return srv.Start(NewContext(a.opts.ctx,a))
		})
	}
	wg.Wait()
	if a.opts.registrar != nil {
		rctx, rcancel := context.WithTimeout(ctx, a.opts.registrarTimeout)
		defer rcancel()
		if err := a.opts.registrar.Register(rctx, instance); err != nil {
			return err
		}
		a.lk.Lock()
		a.instance = instance
		a.lk.Unlock()
	}

	// signal
	c := make(chan os.Signal,1)
	signal.Notify(c,a.opts.sigs...)
	eg.Go(func()error{
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-c:
				if err := a.Stop();err != nil {
					a.opts.logger.Errorf("failed to stop app: %v \n",err)
					return err
				}
			}
		}
	})
	if err := eg.Wait();err != nil && !errors.Is(err, context.Canceled) {
		return err
	}
	return nil
}

// Stop gracefully stops the application.
func (a *App) Stop() error {
	a.lk.Lock()
	instance := a.instance
	a.lk.Unlock()
	if a.opts.registrar != nil && instance != nil {
		ctx, cancel := context.WithTimeout(NewContext(a.ctx, a), a.opts.registrarTimeout)
		defer cancel()
		if err := a.opts.registrar.Deregister(ctx, instance); err != nil {
			return err
		}
	}
	if a.cancel != nil {
		a.cancel()
	}
	return nil
}

// ServiceInstance is an instance of a service in a discovery system.
type ServiceInstance struct {
	// ID is the unique instance ID as registered.
	ID string `json:"id"`
	// Name is the service name as registered.
	Name string `json:"name"`
	// Version is the version of the compiled.
	Version string `json:"version"`
	// Metadata is the kv pair metadata associated with the service instance.
	Metadata map[string]string `json:"metadata"`
	// Endpoints is endpoint addresses of the service instance.
	// schema:
	//   http://127.0.0.1:8000?isSecure=false
	//   grpc://127.0.0.1:9000?isSecure=false
	Endpoints []string `json:"endpoints"`
}

func (a *App) buildInstance() (*registry.ServiceInstance, error) {
	endpoints := make([]string, 0, len(a.opts.endpoints))
	for _, e := range a.opts.endpoints {
		endpoints = append(endpoints, e.String())
	}
	if len(endpoints) == 0 {
		for _, srv := range a.opts.servers {
			if r, ok := srv.(transport.Endpointer); ok {
				e, err := r.Endpoint()
				if err != nil {
					return nil, err
				}
				endpoints = append(endpoints, e.String())
			}
		}
	}
	return &registry.ServiceInstance{
		ID:        a.opts.id,
		Name:      a.opts.name,
		Version:   a.opts.version,
		Metadata:  a.opts.metadata,
		Endpoints: endpoints,
	}, nil
}

type appKey struct{}

// NewContext returns a new Context that carries value.
func NewContext(ctx context.Context, s AppInfo) context.Context {
	return context.WithValue(ctx, appKey{}, s)
}

// FromContext returns the Transport value stored in ctx, if any.
func FromContext(ctx context.Context) (s AppInfo, ok bool) {
	s, ok = ctx.Value(appKey{}).(AppInfo)
	return
}

func New(opts ...Option) *App {
	o := options{
		ctx:              context.Background(),
		logger:           log.NewHelper(log.GetLogger()),
		sigs:             []os.Signal{syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT},
		registrarTimeout: 10 * time.Second,
		stopTimeout:      10 * time.Second,
	}
	if id, err := uuid.NewUUID(); err == nil {
		o.id = id.String()
	}
	for _, opt := range opts {
		opt(&o)
	}
	ctx, cancel := context.WithCancel(o.ctx)
	return &App{
		ctx:    ctx,
		cancel: cancel,
		opts:   o,
	}
}
















