package unitusk

type EndpointProvider interface {
	GetMessages() ([]Message, error)
	Queue() SendQueue
	Errors() []error
	Send(Message) error
}