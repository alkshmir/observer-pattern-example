package main

import "time"

type DomainEvent interface {
	AggregateID() string
	AggregateVersion() int
	OccurredAt() time.Time
	EventType() EventType
}

type EventType string

const (
	UserCreatedEventType EventType = "UserCreated"
)

type UserRegistered struct {
	aggregateID string
	occurredAt  time.Time
}

func NewUserRegistered(aggregateID string) UserRegistered {
	return UserRegistered{
		aggregateID: aggregateID,
		occurredAt:  time.Now(),
	}
}

func (e UserRegistered) AggregateID() string {
	return e.aggregateID
}

func (e UserRegistered) OccurredAt() time.Time {
	return e.occurredAt
}

func (e UserRegistered) EventType() EventType {
	return UserCreatedEventType
}

func (e UserRegistered) AggregateVersion() int {
	return 1
}
