package unitusk

import (
	"io"
	"log"
	"github.com/wallnutkraken/unitusk/gomarkov"
)


// New creates an instance of the Hivemind
func New(logWriter io.Writer, chainLength int, sendIntervalSeconds int, maxGenWords int) Hivemind {
	hm := new(bot)
	hm.logger = log.New(logWriter, "[UniTusk]", log.Ltime)
	hm.brain = gomarkov.NewChain(chainLength)
	man := new(manager)
	if sendIntervalSeconds < 1 {
		sendIntervalSeconds = 30
	}
	go man.sendQueueLoop(sendIntervalSeconds)
	hm.manager = man
	hm.maxLen = maxGenWords
	return hm
}