package main

import (
	"context"
	"log/slog"
)

type UserAppService struct {
	dispatcher *EventDispatcher
}

func NewUserAppService(dispatcher *EventDispatcher) *UserAppService {
	return &UserAppService{
		dispatcher: dispatcher,
	}
}

func (s *UserAppService) RegisterUser(ctx context.Context, command string) error {
	slog.Info("Registering user", "command", command)

	// Register user here
	slog.Info("User registered successfully", "command", command)

	// Publish event
	event := NewUserRegistered(command)
	if err := s.dispatcher.Dispatch(ctx, event); err != nil {
		slog.Error("Failed to dispatch event", "error", err)
		return err
	}
	slog.Info("Event dispatched successfully", "eventType", event.EventType())
	return nil
}

func main() {
	// DI
	dispatcher := NewEventDispatcher()
	dispatcher.Register(UserCreatedEventType, &InternalHandler{})
	dispatcher.Register(UserCreatedEventType, &MQHandler{mqClient: &MockMQClient{}})

	ctx := context.Background()
	userService := NewUserAppService(dispatcher)

	err := userService.RegisterUser(ctx, "user123")
	if err != nil {
		slog.Error("Error registering user", "error", err)
		return
	}
	slog.Info("User registration process completed successfully")
}
