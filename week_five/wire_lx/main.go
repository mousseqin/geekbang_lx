package wire_lx

import (
	"errors"
	"fmt"
	"os"
	"time"
)

// https://github.com/google/wire/blob/main/_tutorial/README.md

type Message string

type Greeter struct {
	Message Message // <- adding a Message field
	Grumpy bool
}

type Event struct {
	Greeter Greeter // <- adding a Greeter field
}

// NewMessage 消息类型只是封装了一个字符串。现在，我们将创建一个简单的初始化器，总是返回一个硬编码的消息。
func NewMessage() Message {
	return Message("Hi there!")
}

// NewGreeter 需要对信息的引用。所以让我们也为我们的Greeter创建一个初始化器。
func NewGreeter(m Message) Greeter {
	var grumpy bool
	if time.Now().Unix()%2 == 0 {
		grumpy = true
	}
	return Greeter{Message: m, Grumpy: grumpy}
}

// Greet 在初始化器中，我们给Greeter分配了一个Message字段。
// 现在，当我们在Greeter上创建一个问候方法时，我们可以使用这个消息。
func (g *Greeter) Greet() Message {
	if g.Grumpy {
		return Message("Go away!")
	}
	return g.Message
}

// NewEvent 接下来，我们需要我们的事件有一个Greeter，所以我们也将为它创建一个初始化器
func NewEvent(g Greeter) (Event, error) {
	if g.Grumpy {
		return Event{}, errors.New("could not create event: event greeter is grumpy")
	}
	return Event{Greeter: g}, nil
}

func (e *Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}

func main() {
	e, err := InitializeEvent()
	if err != nil {
		fmt.Printf("failed to create event: %s\n", err)
		os.Exit(2)
	}
	e.Start()
	// 使用wire依赖注入
	//e := InitializeEvent()
	//e.Start()

	// 不使用WIRE
	//message := NewMessage()
	//greeter := NewGreeter(message)
	//event := NewEvent(greeter)
	//
	//event.Start()
}























