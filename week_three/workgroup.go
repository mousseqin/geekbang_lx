// 包工作组为相关的goroutines组提供同步
package main

// RunFunc 在它自己的goroutine中与其他相关函数一起执行
// 传递给RunFunc的通道的关闭应该触发返回
type RunFunc func( <-chan struct{} ) error

// Group 是一组相关的Goroutines
// 一个组的零值是完全可以使用的，无需初始化
type Group struct {
	fns []RunFunc
}

// Add 添加一个函数
// 当Run被调用时，该函数将在它自己的goroutine中执行
// Add必须在Run之前调用
func (g *Group) Add(fn RunFunc) {
	g.fns = append(g.fns,fn)
}

// Run 在它自己的goroutine中执行通过Add注册的每个函数
// Run 阻塞直到所有的函数都返回，然后从它们中返回第一个非零的错误（如果有的话）
// 第一个返回的函数将触发传递给每个函数的通道的关闭，反过来，它也应该返回
func (g *Group) Run() error {
	if len(g.fns) == 0 {
		return nil
	}

	stop := make(chan struct{})
	done := make(chan error, len(g.fns))
	defer close(done)

	for _, fn := range g.fns {
		go func(fn RunFunc) {
			done <- fn(stop)
		}(fn)
	}

	var err error
	for i := 0; i < cap(done); i++ {
		if err == nil {
			err = <-done
		} else {
			<-done
		}
		if i == 0 {
			close(stop)
		}
	}
	return err
}


