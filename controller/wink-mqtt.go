package controller

import (
	"fmt"
	"github.com/eclipse/paho.mqtt.golang"
	"github.com/mannkind/wink-mqtt/handlers"
	"log"
	"strings"
)

// WinkController - WinkController all the things!
type WinkController struct {
	Client mqtt.Client
	MQTT   struct {
		ClientID  string
		Broker    string
		Username  string
		Password  string
		TopicBase string
		Retain    bool
	}
	RadioAttrs struct {
		Zigbee string
		Zwave  string
		Lutron string
	}
	apron        handlers.Apron
	statuslight  handlers.RGB
	database     handlers.Sqlite3
	watcher      handlers.Tail
	knownDevices map[string]string
}

// Start - Connect and Subscribe
func (t *WinkController) Start() error {
	t.apron = handlers.Apron{}
	t.statuslight = handlers.RGB{}
	t.database = handlers.Sqlite3{
		Zigbee: t.RadioAttrs.Zigbee,
		Zwave:  t.RadioAttrs.Zwave,
		Lutron: t.RadioAttrs.Lutron,
	}
	t.watcher = handlers.Tail{}
	t.knownDevices = make(map[string]string)

	log.Println("Connecting to MQTT: ", t.MQTT.Broker)
	opts := mqtt.NewClientOptions().
		AddBroker(t.MQTT.Broker).
		SetClientID(t.MQTT.ClientID).
		SetOnConnectHandler(t.onConnect).
		SetConnectionLostHandler(func(client mqtt.Client, err error) {
			log.Printf("Disconnected from MQTT: %s.", err)
		}).
		SetUsername(t.MQTT.Username).
		SetPassword(t.MQTT.Password)

	t.Client = mqtt.NewClient(opts)
	if token := t.Client.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}

	t.waitForDeviceUpdate()

	return nil
}

func (t *WinkController) onConnect(client mqtt.Client) {
	log.Println("Connected to MQTT")

	// Subscribe to topics
	subscriptions := map[string]mqtt.MessageHandler{
		fmt.Sprintf("%s/device/add", t.MQTT.TopicBase):                t.addDevice,
		fmt.Sprintf("%s/device/+/delete", t.MQTT.TopicBase):           t.deleteDevice,
		fmt.Sprintf("%s/device/+/updateName", t.MQTT.TopicBase):       t.updateDeviceName,
		fmt.Sprintf("%s/device/+/+/update", t.MQTT.TopicBase):         t.updateDevice,
		fmt.Sprintf("%s/group/add", t.MQTT.TopicBase):                 t.addGroup,
		fmt.Sprintf("%s/group/+/add", t.MQTT.TopicBase):               t.addDeviceToGroup,
		fmt.Sprintf("%s/group/+/delete", t.MQTT.TopicBase):            t.deleteGroup,
		fmt.Sprintf("%s/group/+/+/update", t.MQTT.TopicBase):          t.updateGroup,
		fmt.Sprintf("%s/status_light/state/update", t.MQTT.TopicBase): t.updateStatuslightState,
		fmt.Sprintf("%s/status_light/rgb/update", t.MQTT.TopicBase):   t.updateStatuslightRGB,
	}

	if !client.IsConnected() {
		log.Print("Subscribe Error: Not Connected (Reloading Config?)")
		return
	}

	// Subscribe to the topics
	for topic, handler := range subscriptions {
		if token := client.Subscribe(topic, 0, handler); token.Wait() && token.Error() != nil {
			log.Printf("Subscribe Error: %s", token.Error())
		}
	}
}

func (t *WinkController) addDevice(client mqtt.Client, msg mqtt.Message) {
	topic := msg.Topic()
	payload := string(msg.Payload())
	log.Printf("Received '%s' with payload '%s'", topic, payload)

	t.apron.AddDevice(payload)
}

func (t *WinkController) deleteDevice(client mqtt.Client, msg mqtt.Message) {
	topic := msg.Topic()
	payload := string(msg.Payload())
	log.Printf("Received '%s' with payload '%s'", topic, payload)

	id := strings.Split(topic, "/")[2]

	t.apron.DeleteDevice(id)
}

func (t *WinkController) updateDeviceName(client mqtt.Client, msg mqtt.Message) {
	topic := msg.Topic()
	payload := string(msg.Payload())
	log.Printf("Received '%s' with payload '%s'", topic, payload)

	id := strings.Split(topic, "/")[2]
	name := payload

	t.apron.UpdateDeviceName(id, name)
}

func (t *WinkController) updateDevice(client mqtt.Client, msg mqtt.Message) {
	topic := msg.Topic()
	payload := string(msg.Payload())
	log.Printf("Received '%s' with payload '%s'", topic, payload)

	parts := strings.Split(topic, "/")
	pieces := parts[:len(parts)-1]

	id := pieces[2]
	attr := pieces[3]
	value := payload

	if t.apron.UpdateDevice(id, attr, value) {
		t.publish(strings.Join(pieces, "/"), value)
	}
}

func (t *WinkController) addGroup(client mqtt.Client, msg mqtt.Message) {
	topic := msg.Topic()
	payload := string(msg.Payload())
	log.Printf("Received '%s' with payload '%s'", topic, payload)

	name := payload

	t.apron.AddGroup(name)
}

func (t *WinkController) addDeviceToGroup(client mqtt.Client, msg mqtt.Message) {
	topic := msg.Topic()
	payload := string(msg.Payload())
	log.Printf("Received '%s' with payload '%s'", topic, payload)

	id := payload
	groupID := strings.Split(topic, "/")[2]

	t.apron.AddDeviceToGroup(id, groupID)
}

func (t *WinkController) deleteGroup(client mqtt.Client, msg mqtt.Message) {
	topic := msg.Topic()
	payload := string(msg.Payload())
	log.Printf("Received '%s' with payload '%s'", topic, payload)

	id := strings.Split(topic, "/")[2]

	t.apron.DeleteGroup(id)
}

func (t *WinkController) updateGroup(client mqtt.Client, msg mqtt.Message) {
	topic := msg.Topic()
	payload := string(msg.Payload())
	log.Printf("Received '%s' with payload '%s'", topic, payload)

	parts := strings.Split(topic, "/")
	pieces := parts[:len(parts)-1]

	id := pieces[2]
	attr := pieces[3]
	value := payload

	if t.apron.UpdateGroup(id, attr, value) {
		t.publish(strings.Join(pieces, "/"), value)
	}
}

func (t *WinkController) updateStatuslightState(client mqtt.Client, msg mqtt.Message) {
	topic := msg.Topic()
	payload := string(msg.Payload())
	log.Printf("Received '%s' with payload '%s'", topic, payload)

	result := false
	if payload == "ON" {
		result = t.statuslight.Update("255,255,255")
	} else if payload == "OFF" {
		result = t.statuslight.Update("0,0,0")
	}

	if result {
		parts := strings.Split(topic, "/")
		pieces := parts[:len(parts)-1]
		t.publish(strings.Join(pieces, "/"), payload)
	}
}

func (t *WinkController) updateStatuslightRGB(client mqtt.Client, msg mqtt.Message) {
	topic := msg.Topic()
	payload := string(msg.Payload())
	log.Printf("Received '%s' with payload '%s'", topic, payload)

	parts := strings.Split(payload, " ")
	if len(parts) == 3 {
		t.statuslight.Flash(parts[0], parts[1], parts[2])
	} else {
		t.statuslight.Update(parts[0])
	}

	parts = strings.Split(topic, "/")
	pieces := parts[:len(parts)-1]
	t.publish(strings.Join(pieces, "/"), payload)
}

func (t *WinkController) publish(topic string, payload string) {
	if token := t.Client.Publish(topic, 0, t.MQTT.Retain, payload); token.Wait() && token.Error() != nil {
		log.Printf("Publish Error: %s", token.Error())
		return
	}

	t.knownDevices[topic] = payload
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
			t.publish(topic, value)
		}
	}
}
