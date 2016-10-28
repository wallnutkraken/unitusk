package unitusk

import (
	"github.com/wallnutkraken/unitusk/gomarkov"
	"log"
	"bytes"
	"strings"
	"math/rand"
)

type bot struct {
	manager EndpointManager
	brain *gomarkov.Chain
	logger *log.Logger
}

func (b *bot) UpdateAndFeed() {
	updates, err := b.manager.GatherUpdates()
	for _, e := range err {
		b.logger.Println(e.Error())
	}

	updatesStr := make([]string, len(updates))
	for index, update := range updates {
		updatesStr[index] = update.Text()
	}

	b.brain.Build(bytes.NewBufferString(strings.Join(updatesStr, " ")))
}

func (b *bot) QueueToAll() {
	genStr := b.brain.Generate(rand.Intn(30))
	msg := basicMessage{genStr}

	b.manager.QueueAll(msg)
}

func (b *bot) AddEndpoint(endpoint EndpointProvider) {
	b.manager.AddEndpoint(endpoint)
}

func (b *bot) LogErrors() {
	endpoints := b.manager.Endpoints()
	for _, endpoint := range endpoints {
		errs := endpoint.Errors()
		for _, err := range errs {
			b.logger.Println(err.Error())
		}
	}
}