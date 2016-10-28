package unitusk

type SendQueue interface {
	Clear()
	Add(Message)
	Count() int
	Take() Message
}

type basicQueue struct {
	messages []Message
}

func (q *basicQueue) Clear() {
	q.messages = make([]Message, 0)
}

func (q *basicQueue) Add(msg Message) {
	q.messages = append(q.messages, msg)
}

func (q *basicQueue) Count() int {
	return len(q.messages)
}

func (q *basicQueue) Take() Message {
	if len(q.messages) == 0 {
		return nil
	}
	ret := q.messages[0]
	if len(q.messages) == 1 {
		q.Clear()
	} else {
		q.messages = q.messages[1:]
	}
	return ret;
}

func CreateQueue() SendQueue {
	return new(basicQueue)
}