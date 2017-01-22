package controller

import (
	"fmt"
	"log"
	"strings"

	"github.com/eclipse/paho.mqtt.golang"
	"github.com/mannkind/wink-local/handlers"
)

type winkMQTT struct {
	client   mqtt.Client
	settings struct {
		ClientID  string
		Broker    string
		Username  string
		Password  string
		TopicBase string
		Retain    bool
	}
	people []struct {
		BT    string
		Topic string
	}
	apron       handlers.Apron
	statuslight handlers.RGB
}

func (t *winkMQTT) start() error {
	t.checkDefaults()

	log.Println("Connecting to MQTT: ", t.settings.Broker)
	opts := mqtt.NewClientOptions().
		AddBroker(t.settings.Broker).
		SetClientID(t.settings.ClientID).
		SetOnConnectHandler(t.onConnect).
		SetConnectionLostHandler(func(client mqtt.Client, err error) {
			log.Printf("Disconnected from MQTT: %s.", err)
		}).
		SetUsername(t.settings.Username).
		SetPassword(t.settings.Password)

	t.client = mqtt.NewClient(opts)
	if token := t.client.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}

	return nil
}

func (t *winkMQTT) publish(topic string, payload string) {
	if token := t.client.Publish(topic, 0, t.settings.Retain, payload); token.Wait() && token.Error() != nil {
		log.Printf("Publish Error: %s", token.Error())
		return
	}
}

func (t *winkMQTT) checkDefaults() {
	if len(t.settings.ClientID) == 0 {
		t.settings.ClientID = "WinkLocal"
	}
}

func (t *winkMQTT) onConnect(client mqtt.Client) {
	log.Println("Connected to MQTT")

	// Subscribe to topics
	subscriptions := map[string]mqtt.MessageHandler{
		fmt.Sprintf("%s/device/add", t.settings.TopicBase):                t.addDevice,
		fmt.Sprintf("%s/device/delete", t.settings.TopicBase):             t.deleteDevice,
		fmt.Sprintf("%s/device/+/update", t.settings.TopicBase):           t.updateDeviceName,
		fmt.Sprintf("%s/device/+/+/update", t.settings.TopicBase):         t.updateDevice,
		fmt.Sprintf("%s/group/add", t.settings.TopicBase):                 t.addGroup,
		fmt.Sprintf("%s/group/delete", t.settings.TopicBase):              t.deleteGroup,
		fmt.Sprintf("%s/group/+/add", t.settings.TopicBase):               t.addDeviceToGroup,
		fmt.Sprintf("%s/group/+/delete", t.settings.TopicBase):            t.deleteDeviceFromGroup,
		fmt.Sprintf("%s/group/+/+/update", t.settings.TopicBase):          t.updateGroup,
		fmt.Sprintf("%s/status_light/state/update", t.settings.TopicBase): t.updateStatuslightState,
		fmt.Sprintf("%s/status_light/rgb/update", t.settings.TopicBase):   t.updateStatuslightRGB,
		"home/people": t.searchForPeople,
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

func (t *winkMQTT) addDevice(client mqtt.Client, msg mqtt.Message) {
	payload := string(msg.Payload())
	radio := strings.ToLower(payload)

	t.apron.AddDevice(radio)
}

func (t *winkMQTT) deleteDevice(client mqtt.Client, msg mqtt.Message) {
	id := string(msg.Payload())

	t.apron.DeleteDevice(id)
}

func (t *winkMQTT) updateDeviceName(client mqtt.Client, msg mqtt.Message) {
	topic := msg.Topic()
	payload := string(msg.Payload())

	id := strings.Split(topic, "/")[2]
	name := payload

	t.apron.UpdateDeviceName(id, name)
}

func (t *winkMQTT) updateDevice(client mqtt.Client, msg mqtt.Message) {
	topic := msg.Topic()
	payload := string(msg.Payload())

	parts := strings.Split(topic, "/")
	pieces := parts[:len(parts)-1]

	id := pieces[2]
	attr := pieces[3]
	value := payload

	if t.apron.UpdateDevice(id, attr, value) {
		t.publish(strings.Join(pieces, "/"), value)
	}
}

func (t *winkMQTT) addGroup(client mqtt.Client, msg mqtt.Message) {
	name := string(msg.Payload())

	t.apron.AddGroup(name)
}

func (t *winkMQTT) addDeviceToGroup(client mqtt.Client, msg mqtt.Message) {
	topic := msg.Topic()
	payload := string(msg.Payload())

	groupID := strings.Split(topic, "/")[2]
	id := payload

	t.apron.AddDeviceToGroup(id, groupID)
}

func (t *winkMQTT) deleteDeviceFromGroup(client mqtt.Client, msg mqtt.Message) {
	topic := msg.Topic()
	payload := string(msg.Payload())

	groupID := strings.Split(topic, "/")[2]
	id := payload

	t.apron.DeleteDeviceFromGroup(id, groupID)
}

func (t *winkMQTT) deleteGroup(client mqtt.Client, msg mqtt.Message) {
	id := string(msg.Payload())

	t.apron.DeleteGroup(id)
}

func (t *winkMQTT) updateGroup(client mqtt.Client, msg mqtt.Message) {
	topic := msg.Topic()
	payload := string(msg.Payload())

	parts := strings.Split(topic, "/")
	pieces := parts[:len(parts)-1]

	id := pieces[2]
	attr := pieces[3]
	value := payload

	if t.apron.UpdateGroup(id, attr, value) {
		t.publish(strings.Join(pieces, "/"), value)
	}
}

func (t *winkMQTT) updateStatuslightState(client mqtt.Client, msg mqtt.Message) {
	topic := msg.Topic()
	payload := string(msg.Payload())

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

func (t *winkMQTT) updateStatuslightRGB(client mqtt.Client, msg mqtt.Message) {
	topic := msg.Topic()
	payload := string(msg.Payload())

	parts := strings.Split(payload, " ")
	if len(parts) == 3 {
		color := parts[0]
		alt := parts[1]
		ms := parts[2]
		t.statuslight.Flash(color, alt, ms)
	} else {
		color := parts[0]
		t.statuslight.Update(color)
	}

	parts = strings.Split(topic, "/")
	pieces := parts[:len(parts)-1]
	t.publish(strings.Join(pieces, "/"), payload)
}

func (t *winkMQTT) searchForPeople(client mqtt.Client, msg mqtt.Message) {
	payload := string(msg.Payload())
	if payload != "Find" {
		return
	}

	bt := handlers.BT{}
	bt.Up()

	for _, person := range t.people {
		log.Printf("Searching for %s", person.BT)
		state := bt.FindPerson(person.BT)

		log.Printf("Publishing %s %s", person.Topic, state)
		t.publish(person.Topic, state)
	}
}
