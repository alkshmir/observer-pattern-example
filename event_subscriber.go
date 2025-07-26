package main

import (
	"context"
	"log/slog"
)

type EventSubscriber interface {
	Handle(ctx context.Context, event DomainEvent) error
}

type InternalHandler struct {
}

func (h *InternalHandler) Handle(ctx context.Context, event DomainEvent) error {
	slog.Info("InternalHandler handling event",
		"eventType", event.EventType(),
	)
	return nil
}

type MQHandler struct {
	mqClient MQClient
}

func (h *MQHandler) Handle(ctx context.Context, event DomainEvent) error {
	slog.Info("MQHandler handling event",
		"eventType", event.EventType(),
	)

	if err := h.mqClient.Publish(ctx, event); err != nil {
		slog.Error("Failed to publish event to MQ", "error", err)
		return err
	}

	return nil
}
