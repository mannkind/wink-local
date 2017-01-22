package handlers

import (
	"log"
	"strings"

	"github.com/hpcloud/tail"
)

// Tail - Tail the wink hub log file
type Tail struct {
	log string
}

// Start - Start the tailing process, send notifications over the channel
func (t *Tail) Start(deviceUpdated chan bool) {
	t.log = "/tmp/all.log"
	log.Printf("Watching %s for polled changes", t.log)
	out, err := tail.TailFile(t.log, tail.Config{Follow: true, ReOpen: true, Poll: true, MustExist: true})
	if err != nil {
		log.Printf("Failed to watch %s; %s", t.log, err)
		deviceUpdated <- false
		return
	}

	for line := range out.Lines {
		if strings.Contains(line.Text, "state changed in device") {
			log.Print("Polling State Change")
			deviceUpdated <- true
		}
	}
}
