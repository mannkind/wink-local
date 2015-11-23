package controller

import (
	"testing"

	"github.com/eclipse/paho.mqtt.golang"
	"gopkg.in/yaml.v2"
)

var testMqttClient = mqtt.NewClient(mqtt.NewClientOptions())

func defaultTestController() *WinkController {
	var testConfig = `
        mqtt:
          clientid: 'WinkMQTT'
          broker: "tcp://test.mosquitto.org:1883"
          basetopic: 'winkhub'
          retain: true
    `

	myCtrl := WinkController{}
	err := yaml.Unmarshal([]byte(testConfig), &myCtrl)
	if err != nil {
		panic(err)
	}
	return &myCtrl
}

func TestMqttStart(t *testing.T) {
	myCtrl := defaultTestController()
	if err := myCtrl.Start(); err != nil {
		t.Error("Something went wrong starting!")
	}
}

func TestMqttConnect(t *testing.T) {
	myCtrl := defaultTestController()
	myCtrl.onConnect(testMqttClient)
}

func TestMqttAddDevice(t *testing.T) {
	myCtrl := defaultTestController()
	var tests = []struct {
		Topic   string
		Payload string
		Cmd     string
	}{
		{"winkhub/device/add", "zigbee", "aprontest -a -r zigbee"},
		{"winkhub/device/add", "zwave", "aprontest -a -r zwave"},
		{"winkhub/device/add", "lutron", "aprontest -a -r lutron"},
		{"winkhub/device/add", "totes", ""},
	}

	for _, v := range tests {
		msg := &mockMessage{
			topic:   v.Topic,
			payload: []byte(v.Payload),
		}

		expected := v.Cmd

		myCtrl.apron.LastRun = ""
		myCtrl.addDevice(testMqttClient, msg)
		if myCtrl.apron.LastRun != expected {
			t.Errorf("Wrong cmd - Actual: %s, Expected: %s", myCtrl.apron.LastRun, expected)
		}
	}
}

func TestMqttAddGroup(t *testing.T) {
	myCtrl := defaultTestController()
	var tests = []struct {
		Topic   string
		Payload string
		Cmd     string
	}{
		{"winkhub/group/add", "Group1", "aprontest -a -s Group1"},
		{"winkhub/group/add", "", ""},
	}

	for _, v := range tests {
		msg := &mockMessage{
			topic:   v.Topic,
			payload: []byte(v.Payload),
		}

		expected := v.Cmd

		myCtrl.apron.LastRun = ""
		myCtrl.addGroup(testMqttClient, msg)

		if myCtrl.apron.LastRun != expected {
			t.Errorf("Wrong cmd - Actual: %s, Expected: %s", myCtrl.apron.LastRun, expected)
		}
	}
}

func TestMqttDeleteGroup(t *testing.T) {
	myCtrl := defaultTestController()
	var tests = []struct {
		Topic   string
		Payload string
		Cmd     string
	}{
		{"winkhub/group/1/delete", "", "aprontest -d -w 1"},
		{"winkhub/group/A/delete", "", ""},
	}

	for _, v := range tests {
		msg := &mockMessage{
			topic:   v.Topic,
			payload: []byte(v.Payload),
		}

		expected := v.Cmd

		myCtrl.apron.LastRun = ""
		myCtrl.deleteGroup(testMqttClient, msg)

		if myCtrl.apron.LastRun != expected {
			t.Errorf("Wrong cmd - Actual: %s, Expected: %s", myCtrl.apron.LastRun, expected)
		}
	}
}

func TestMqttAddDeviceToGroup(t *testing.T) {
	myCtrl := defaultTestController()
	var tests = []struct {
		Topic   string
		Payload string
		Cmd     string
	}{
		{"winkhub/group/2/add", "1", "aprontest -a -x 2 -m 1"},
		{"winkhub/group/2/add", "", ""},
		{"winkhub/group/A/add", "", ""},
	}

	for _, v := range tests {
		msg := &mockMessage{
			topic:   v.Topic,
			payload: []byte(v.Payload),
		}

		expected := v.Cmd

		myCtrl.apron.LastRun = ""
		myCtrl.addDeviceToGroup(testMqttClient, msg)

		if myCtrl.apron.LastRun != expected {
			t.Errorf("Wrong cmd - Actual: %s, Expected: %s", myCtrl.apron.LastRun, expected)
		}
	}
}

func TestMqttDeleteDevice(t *testing.T) {
	myCtrl := defaultTestController()
	var tests = []struct {
		Topic   string
		Payload string
		Cmd     string
	}{
		{"winkhub/device/1/delete", "", "aprontest -d -m 1"},
		{"winkhub/device/A/delete", "", ""},
	}

	for _, v := range tests {
		msg := &mockMessage{
			topic:   v.Topic,
			payload: []byte(v.Payload),
		}

		expected := v.Cmd

		myCtrl.apron.LastRun = ""
		myCtrl.deleteDevice(testMqttClient, msg)
		if myCtrl.apron.LastRun != expected {
			t.Errorf("Wrong cmd - Actual: %s, Expected: %s", myCtrl.apron.LastRun, expected)
		}
	}
}

func TestMqttUpdateGroup(t *testing.T) {
	myCtrl := defaultTestController()
	myCtrl.Start()
	var tests = []struct {
		Topic   string
		Payload string
		Cmd     string
	}{
		{"winkhub/group/1/2/update", "5", "aprontest -u -x 1 -t 2 -v 5"},
		{"winkhub/group/1/A/update", "", ""},
		{"winkhub/group/A/A/update", "", ""},
	}

	for _, v := range tests {
		msg := &mockMessage{
			topic:   v.Topic,
			payload: []byte(v.Payload),
		}

		expected := v.Cmd

		myCtrl.apron.LastRun = ""
		myCtrl.updateGroup(testMqttClient, msg)

		if myCtrl.apron.LastRun != expected {
			t.Errorf("Wrong cmd - Actual: %s, Expected: %s", myCtrl.apron.LastRun, expected)
		}
	}
}

func TestMqttUpdateDevice(t *testing.T) {
	myCtrl := defaultTestController()
	myCtrl.Start()
	var tests = []struct {
		Topic   string
		Payload string
		Cmd     string
	}{
		{"winkhub/device/1/2/update", "5", "aprontest -u -m 1 -t 2 -v 5"},
		{"winkhub/device/1/A/update", "", ""},
		{"winkhub/device/A/A/update", "", ""},
	}

	for _, v := range tests {
		msg := &mockMessage{
			topic:   v.Topic,
			payload: []byte(v.Payload),
		}

		expected := v.Cmd

		myCtrl.apron.LastRun = ""
		myCtrl.updateDevice(testMqttClient, msg)

		if myCtrl.apron.LastRun != expected {
			t.Errorf("Wrong cmd - Actual: %s, Expected: %s", myCtrl.apron.LastRun, expected)
		}
	}
}

func TestMqttUpdateDeviceName(t *testing.T) {
	myCtrl := defaultTestController()
	var tests = []struct {
		Topic   string
		Payload string
		Cmd     string
	}{
		{"winkhub/device/1/updateName", "TestingName", "aprontest -m 1 --set-name TestingName"},
		{"winkhub/device/1/updateName", "", ""},
		{"winkhub/device/A/updateName", "", ""},
	}

	for _, v := range tests {
		msg := &mockMessage{
			topic:   v.Topic,
			payload: []byte(v.Payload),
		}

		expected := v.Cmd

		myCtrl.apron.LastRun = ""
		myCtrl.updateDeviceName(testMqttClient, msg)

		if myCtrl.apron.LastRun != expected {
			t.Errorf("Wrong cmd - Actual: %s, Expected: %s", myCtrl.apron.LastRun, expected)
		}
	}
}

func TestMqttUpdateStatuslightState(t *testing.T) {
	myCtrl := defaultTestController()
	myCtrl.Start()
	var tests = []struct {
		Topic   string
		Payload string
		Cmd     string
	}{
		{"winkhub/status_light/state/update", "ON", "set_rgb 255 255 255"},
		{"winkhub/status_light/state/update", "", ""},
	}

	for _, v := range tests {
		msg := &mockMessage{
			topic:   v.Topic,
			payload: []byte(v.Payload),
		}

		expected := v.Cmd

		myCtrl.statuslight.LastRun = ""
		myCtrl.updateStatuslightState(testMqttClient, msg)

		if myCtrl.statuslight.LastRun != expected {
			t.Errorf("Wrong cmd - Actual: %s, Expected: %s", myCtrl.statuslight.LastRun, expected)
		}
	}
}

func TestMqttUpdateStatuslightRGB(t *testing.T) {
	myCtrl := defaultTestController()
	myCtrl.Start()
	var tests = []struct {
		Topic   string
		Payload string
		Cmd     string
	}{
		{"winkhub/status_light/rgb/update", "0,255,0", "set_rgb 0 255 0"},
		{"winkhub/status_light/rgb/update", "", ""},
	}

	for _, v := range tests {
		msg := &mockMessage{
			topic:   v.Topic,
			payload: []byte(v.Payload),
		}

		expected := v.Cmd

		myCtrl.statuslight.LastRun = ""
		myCtrl.updateStatuslightRGB(testMqttClient, msg)

		if myCtrl.statuslight.LastRun != expected {
			t.Errorf("Wrong cmd - Actual: %s, Expected: %s", myCtrl.statuslight.LastRun, expected)
		}
	}
}

type mockMessage struct {
	topic   string
	payload []byte
}

func (m *mockMessage) Duplicate() bool {
	return true
}

func (m *mockMessage) Qos() byte {
	return 'a'
}

func (m *mockMessage) Retained() bool {
	return true
}

func (m *mockMessage) Topic() string {
	return m.topic
}

func (m *mockMessage) MessageID() uint16 {
	return 0
}

func (m *mockMessage) Payload() []byte {
	return m.payload
}
