package eventx

// EventHandler is the interface for handling events.
type EventHandler interface {
	Handle(event DomainEvent)
}
