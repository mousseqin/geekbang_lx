package main

type Once chan struct {}

func NewOnce() Once {
	o := make(Once,1)
	o <- struct {}{}
	return o
}

func (o Once) Do(f func()){
	// 从一个封闭的通道中读取总是成功的。
	// 这只是在初始化过程中的阻塞
	_, ok := <-o
	if !ok {
		// 通道已关闭，提前返回。
		// f必须已经返回。
		return
	}
	// 调用f。
	// 只有一个goroutine会到这里，因为只有一个值
	// 在通道中。
	f()
	// 这就释放了所有等待的goroutines和未来的goroutines。
	close(o)
}
