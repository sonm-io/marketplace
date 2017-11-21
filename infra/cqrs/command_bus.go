package cqrs

import "fmt"

// Command defines command pattern interface.
type Command interface {
	CommandID() string
}

// Handler defines command handler pattern interface.
type Handler interface {
	Handle(Command) error
}

// CommandBus a command bus.
type CommandBus struct {
	handlers map[string]Handler
}

// NewCommandBus creates a new instance of CommandBus.
func NewCommandBus() *CommandBus {
	return &CommandBus{
		handlers: make(map[string]Handler),
	}
}

// Handle dispatches the given Command to its appropriate Handler and handles it.
func (ch *CommandBus) Handle(c Command) error {

	h, err := ch.resolveHandler(c.CommandID())
	if err != nil {
		return err
	}

	return h.Handle(c)
}

// RegisterHandler registers a Handler for the given CommandID.
func (ch *CommandBus) RegisterHandler(commandID string, h Handler) *CommandBus {
	ch.handlers[commandID] = h
	return ch
}

func (ch *CommandBus) resolveHandler(commandID string) (Handler, error) {

	h, ok := ch.handlers[commandID]
	if !ok {
		return nil, fmt.Errorf("handler for %q is not found", commandID)

	}
	return h, nil
}
