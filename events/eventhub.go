package events

import (
	"fmt"
)

type Event struct {
	message string
}

type Eventhub struct {
	listeners []Listener
}

type Listener interface {
	Process(e *Event)
}

type ConsoleListener struct {
}

func (n *ConsoleListener) Process(e *Event) {
	fmt.Println(e.message)
}

func (eh *Eventhub) RegisterListener(l Listener) {
	eh.listeners = append(eh.listeners, l)
}

func (eh *Eventhub) RegisterEvent(e *Event) {
	for index := 0; index < len(eh.listeners); index++ {
		l := eh.listeners[index]
		l.Process(e)
	}
}
