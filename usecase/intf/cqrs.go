package intf

// Command defines command pattern interface.
type Command interface {
	CommandID() string
}

// CommandHandler defines command handler pattern interface.
type CommandHandler interface {
	// Handle handles the given Command.
	Handle(Command) error
}

// Query defines query pattern ('Q' in CQRS) interface.
type Query interface {
	QueryID() string
}

// QueryHandler defines query handler pattern interface.
type QueryHandler interface {
	// Handle handles the given Query and returns the result.
	Handle(q Query, result interface{}) error
}

// Event defines an event.
type Event interface {
	EventID() string
}

// EventHandler defines an event handler.
type EventHandler interface {
	// Handle handles the given Event.
	Handle(Event) error
}
