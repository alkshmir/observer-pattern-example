package main

import (
	"context"
	"sync"
)

type EventDispatcher struct {
	subscribers map[EventType][]EventSubscriber
	mu          sync.RWMutex
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		subscribers: make(map[EventType][]EventSubscriber),
	}
}

func (d *EventDispatcher) Register(eventType EventType, subscriber EventSubscriber) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.subscribers[eventType] = append(d.subscribers[eventType], subscriber)
}

func (d *EventDispatcher) Dispatch(ctx context.Context, event DomainEvent) error {
	d.mu.RLock()
	defer d.mu.RUnlock()

	subs := d.subscribers[event.EventType()]
	var errs []error

	for _, sub := range subs {
		if err := sub.Handle(ctx, event); err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		return SubscriberErrors{Errors: errs}
	}

	return nil
}

type SubscriberErrors struct {
	Errors []error
}

func (e SubscriberErrors) Error() string {
	if len(e.Errors) == 0 {
		return "no errors"
	}

	var msg string
	for _, err := range e.Errors {
		msg += err.Error() + "; "
	}
	return msg[:len(msg)-2] // Remove trailing "; "
}
