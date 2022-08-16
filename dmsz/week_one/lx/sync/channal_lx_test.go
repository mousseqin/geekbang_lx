package sync

import (
	"github.com/davecgh/go-spew/spew"
	"testing"
)

func TestBroker(t *testing.T) {
	b := &Broker{
		consumer: make([]*Consumer, 0, 10),
	}

	ch1 := &Consumer{ch: make(chan string, 1)}
	ch2 := &Consumer{ch: make(chan string, 1)}

	b.Subscribe(ch1)
	b.Subscribe(ch2)

	b.Produce("hello world")
	spew.Dump(<-ch1.ch)
	spew.Dump(<-ch2.ch)
}

type Consumer struct {
	ch chan string
}

type Broker struct {
	consumer []*Consumer
}

func (b *Broker) Produce(msg string) {
	for _, x := range b.consumer {
		x.ch <- msg
	}
}

func (b *Broker) Subscribe(c *Consumer) {
	b.consumer = append(b.consumer, c)
}

// ---------------------------  broker1 ----------------------------- //

func TestBroker1(t *testing.T) {
	b := NewBroker1()
	b.Produce("test123")
	b.Subscribe(func(s string) {
		spew.Dump(s)
	})

	b.Start()
}

type Broker1 struct {
	ch        chan string
	consumers []func(s string)
}

func (b *Broker1) Produce(msg string) {
	b.ch <- msg
}

func (b *Broker1) Subscribe(consume func(s string)) {
	b.consumers = append(b.consumers, consume)
}

func (b *Broker1) Start() {
	go func() {
		s := <-b.ch
		for _, x := range b.consumers {
			x(s)
		}
	}()
}

func NewBroker1() *Broker1 {
	b := &Broker1{ch: make(chan string, 10), consumers: make([]func(s string), 0, 10)}
	go func() {
		s := <-b.ch
		for _, c := range b.consumers {
			c(s)
		}
	}()
	return b
}
