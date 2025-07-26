package main

import (
	"context"
	"log/slog"
)

type MQClient interface {
	Publish(ctx context.Context, event DomainEvent) error
}

type MockMQClient struct{}

func (m *MockMQClient) Publish(ctx context.Context, event DomainEvent) error {
	// Simulate publishing to a message queue
	slog.Info("MockMQClient published event",
		"eventType", event.EventType(),
	)
	return nil
}
