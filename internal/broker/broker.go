package broker

import (
	"net"
	"sync"
)

type Broker struct {
	mu          sync.RWMutex
	subscribers map[string][]net.Conn
}

func New() *Broker {
	return &Broker{
		subscribers: make(map[string][]net.Conn),
	}
}

func (b *Broker) Subscribe(topic string, conn net.Conn) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.subscribers[topic] = append(b.subscribers[topic], conn)
}

func (b *Broker) Publish(topic string, msg string, sender net.Conn) {
	b.mu.RLock()
	defer b.mu.RUnlock()
	for _, conn := range b.subscribers[topic] {
		if sender == conn {
			continue
		}
		conn.Write([]byte(msg))
	}
}
