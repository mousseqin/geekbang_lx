package goleak

// go get -u go.uber.org/goleak

func leak() {
	ch := make(chan struct{})
	go func() {
		ch <- struct{}{}
	}()
}
