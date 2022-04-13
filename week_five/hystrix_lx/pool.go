package hystrix_lx

type executorPool struct {
	Name string
	Max int

	Tickets chan *struct{}
}
