package unitusk

type Message interface {
	Text() string
}

type basicMessage struct {
	text string
}

func (msg basicMessage) Text() string {
	return msg.text
}