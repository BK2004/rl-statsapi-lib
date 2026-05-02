package publisher

import (
	"sync"
)

const (
	SUBSCRIBER_BUFFER_SIZE = 8
)

type Publisher[T any] interface {
	Subscribe() chan T
	Publish(data T)
}

type publisher[T any] struct {
	subscribers      []chan T
	subscribersMutex sync.RWMutex
	Topic            string
}

func New[T any](topic string) Publisher[T] {
	return &publisher[T]{
		subscribers:      make([]chan T, 0),
		subscribersMutex: sync.RWMutex{},
		Topic:            topic,
	}
}

func (p *publisher[T]) Subscribe() chan T {
	p.subscribersMutex.Lock()
	defer p.subscribersMutex.Unlock()

	ch := make(chan T, SUBSCRIBER_BUFFER_SIZE)
	p.subscribers = append(p.subscribers, ch)
	return ch
}

func (p *publisher[T]) Publish(data T) {
	p.subscribersMutex.RLock()
	defer p.subscribersMutex.RUnlock()

	for _, ch := range p.subscribers {
		select {
		case ch <- data:
		default:
		}
	}
}
