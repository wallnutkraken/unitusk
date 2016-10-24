package unitusk

type Hivemind interface {
	UpdateAndFeed()
	QueueToAll()
	LogErrors()
}