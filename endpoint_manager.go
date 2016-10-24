package unitusk

import "time"

type EndpointManager interface {
	Endpoints() []EndpointProvider
	QueueAll(Message)
	AddEndpoint(EndpointProvider)
	RemoveEndpoint(EndpointProvider)
	GatherUpdates() ([]Message, []error)
}

type manager struct {
	endpoints []EndpointProvider
}

func (man *manager) Endpoints() []EndpointProvider {
	return man.endpoints
}

func (man *manager) QueueAll(msg Message) {
	for _, endpoint := range man.endpoints {
		endpoint.Queue().Add(msg)
	}
}

func (man *manager) AddEndpoint(newProvider EndpointProvider) {
	man.endpoints = append(man.endpoints, newProvider)
}

func (man *manager) RemoveEndpoint(provider EndpointProvider) {
	for index := 0; index < len(man.endpoints); index++ {
		if man.endpoints[index] == provider {
			/* Move last to current spot */
			man.endpoints[index] = man.endpoints[len(man.endpoints) - 1]
			/* Remove last now */
			man.endpoints = man.endpoints[:len(man.endpoints)-1]
			return
		}
	}
}

func (man *manager) GatherUpdates() ([]Message, []error) {
	var allMsgs = make([]Message, 0)
	var allErrs = make([]error, 0)
	for _, endpoint := range man.endpoints {
		msgs, err := endpoint.GetMessages()
		if err != nil {
			allErrs = append(allErrs, err)
		} else {
			allMsgs = append(allMsgs, msgs...)
		}
	}

	return allMsgs, allErrs
}

func (man *manager) sendQueueLoop(interval int) {
	for {
		time.Sleep(time.Second * time.Duration(interval))
		for _, endp := range man.endpoints {
			msg := endp.Queue().Take()
			if msg != nil {
				endp.Send(msg)
			}
		}
	}
}