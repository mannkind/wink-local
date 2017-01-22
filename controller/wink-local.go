package controller

import (
	"fmt"
	"log"
	"strings"

	"github.com/mannkind/wink-local/handlers"
)

// WinkController - WinkController all the things!
type WinkController struct {
	HTTP struct {
		Port     int16
		Username string
		Password string
	}
	MQTT struct {
		ClientID  string
		Broker    string
		Username  string
		Password  string
		TopicBase string
		Retain    bool
	}
	People []struct {
		BT    string
		Topic string
	}
	RadioAttrs struct {
		Zigbee string
		Zwave  string
		Lutron string
	}

	winkMQTT     winkMQTT
	winkHTTP     winkHTTP
	database     handlers.Sqlite3
	watcher      handlers.Tail
	knownDevices map[string]string
}

// Start - Start WinkController
func (t *WinkController) Start() error {
	t.winkMQTT.settings = t.MQTT
	t.winkMQTT.people = t.People
	t.winkHTTP.settings = t.HTTP
	t.database.Zigbee = t.RadioAttrs.Zigbee
	t.database.Zwave = t.RadioAttrs.Zwave
	t.database.Lutron = t.RadioAttrs.Lutron

	t.knownDevices = make(map[string]string)

	t.winkMQTT.start()
	t.winkHTTP.start()

	t.waitForDeviceUpdate()

	return nil
}

func (t *WinkController) waitForDeviceUpdate() {
	// Refresh from DB every so often
	deviceUpdated := make(chan bool)
	go t.watcher.Start(deviceUpdated)

	go func(deviceUpdated chan bool) {
		for _ = range deviceUpdated {
			log.Print("Running DB Comparison")
			t.databaseComparison()
		}
	}(deviceUpdated)
}

func (t *WinkController) databaseComparison() {
	stdOutPieces, err := t.database.States()
	if err != nil {
		log.Printf("Unable to process database results due to error; %s", err)
	}

	for _, device := range stdOutPieces {
		devicePieces := strings.Split(device, ",")

		if len(devicePieces) != 3 {
			continue
		}

		id := devicePieces[0]
		attr := devicePieces[1]

		topic := fmt.Sprintf("%s/device/%s/%s", t.MQTT.TopicBase, id, attr)
		value := devicePieces[2]

		if prevValue, ok := t.knownDevices[topic]; !ok || prevValue != value {
			log.Printf("Updating %s %s", topic, value)
			t.winkMQTT.publish(topic, value)
			t.knownDevices[topic] = value
		}
	}
}
