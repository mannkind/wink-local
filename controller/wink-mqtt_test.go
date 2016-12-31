package controller

import (
	"testing"
	"time"
)

var testWinkMQTT = func() *winkMQTT {
	myCtrl := winkMQTT{
		settings: struct {
			ClientID  string
			Broker    string
			Username  string
			Password  string
			TopicBase string
			Retain    bool
		}{
			ClientID:  "WinkLocal",
			Broker:    "tcp://test.mosquitto.org:1883",
			TopicBase: "winkhub",
			Retain:    false,
		},
	}

	myCtrl.start()
	return &myCtrl
}()

func TestMqttAddDevice(t *testing.T) {
	var tests = []struct {
		Topic   string
		Payload string
		Cmd     string
	}{
		{"/device/add", "zigbee", "aprontest -a -r zigbee"},
		{"/device/add", "zwave", "aprontest -a -r zwave"},
		{"/device/add", "lutron", "aprontest -a -r lutron"},
		{"/device/add", "totes", ""},
	}

	for _, v := range tests {
		msg := &mockMessage{
			topic:   v.Topic,
			payload: []byte(v.Payload),
		}

		testWinkMQTT.apron.LastRun = ""
		testWinkMQTT.addDevice(testWinkMQTT.client, msg)
		if testWinkMQTT.apron.LastRun != v.Cmd {
			t.Errorf("Wrong cmd - Actual: %s, Expected: %s", testWinkMQTT.apron.LastRun, v.Cmd)
		}
	}
}

func TestMqttAddGroup(t *testing.T) {
	var tests = []struct {
		Topic   string
		Payload string
		Cmd     string
	}{
		{"/group/add", "Group1", "aprontest -a -s Group1"},
		{"/group/add", "", ""},
	}

	for _, v := range tests {
		msg := &mockMessage{
			topic:   v.Topic,
			payload: []byte(v.Payload),
		}

		testWinkMQTT.apron.LastRun = ""
		testWinkMQTT.addGroup(testWinkMQTT.client, msg)

		if testWinkMQTT.apron.LastRun != v.Cmd {
			t.Errorf("Wrong cmd - Actual: %s, Expected: %s", testWinkMQTT.apron.LastRun, v.Cmd)
		}
	}
}

func TestMqttDeleteGroup(t *testing.T) {
	var tests = []struct {
		Topic   string
		Payload string
		Cmd     string
	}{
		{"/group/delete", "1", "aprontest -d -w 1"},
		{"/group/delete", "A", ""},
	}

	for _, v := range tests {
		msg := &mockMessage{
			topic:   v.Topic,
			payload: []byte(v.Payload),
		}

		testWinkMQTT.apron.LastRun = ""
		testWinkMQTT.deleteGroup(testWinkMQTT.client, msg)

		if testWinkMQTT.apron.LastRun != v.Cmd {
			t.Errorf("Wrong cmd - Actual: %s, Expected: %s", testWinkMQTT.apron.LastRun, v.Cmd)
		}
	}
}

func TestMqttAddDeviceToGroup(t *testing.T) {
	var tests = []struct {
		Topic   string
		Payload string
		Cmd     string
	}{
		{"/group/2/add", "1", "aprontest -a -x 2 -m 1"},
		{"/group/2/add", "", ""},
		{"/group/A/add", "", ""},
	}

	for _, v := range tests {
		msg := &mockMessage{
			topic:   v.Topic,
			payload: []byte(v.Payload),
		}

		testWinkMQTT.apron.LastRun = ""
		testWinkMQTT.addDeviceToGroup(testWinkMQTT.client, msg)

		if testWinkMQTT.apron.LastRun != v.Cmd {
			t.Errorf("Wrong cmd - Actual: %s, Expected: %s", testWinkMQTT.apron.LastRun, v.Cmd)
		}
	}
}

func TestMqttDeleteDeviceFromGroup(t *testing.T) {
	var tests = []struct {
		Topic   string
		Payload string
		Cmd     string
	}{
		{"/group/2/delete", "1", "aprontest -d -x 2 -m 1"},
		{"/group/2/delete", "", ""},
		{"/group/A/delete", "", ""},
	}

	for _, v := range tests {
		msg := &mockMessage{
			topic:   v.Topic,
			payload: []byte(v.Payload),
		}

		testWinkMQTT.apron.LastRun = ""
		testWinkMQTT.deleteDeviceFromGroup(testWinkMQTT.client, msg)

		if testWinkMQTT.apron.LastRun != v.Cmd {
			t.Errorf("Wrong cmd - Actual: %s, Expected: %s", testWinkMQTT.apron.LastRun, v.Cmd)
		}
	}
}

func TestMqttDeleteDevice(t *testing.T) {
	var tests = []struct {
		Topic   string
		Payload string
		Cmd     string
	}{
		{"/device/delete", "1", "aprontest -d -m 1"},
		{"/device/delete", "A", ""},
	}

	for _, v := range tests {
		msg := &mockMessage{
			topic:   v.Topic,
			payload: []byte(v.Payload),
		}

		testWinkMQTT.apron.LastRun = ""
		testWinkMQTT.deleteDevice(testWinkMQTT.client, msg)
		if testWinkMQTT.apron.LastRun != v.Cmd {
			t.Errorf("Wrong cmd - Actual: %s, Expected: %s", testWinkMQTT.apron.LastRun, v.Cmd)
		}
	}
}

func TestMqttUpdateGroup(t *testing.T) {
	var tests = []struct {
		Topic   string
		Payload string
		Cmd     string
	}{
		{"/group/1/2/update", "5", "aprontest -u -x 1 -t 2 -v 5"},
		{"/group/1/A/update", "", ""},
		{"/group/A/A/update", "", ""},
	}

	for _, v := range tests {
		msg := &mockMessage{
			topic:   v.Topic,
			payload: []byte(v.Payload),
		}

		testWinkMQTT.apron.LastRun = ""
		testWinkMQTT.updateGroup(testWinkMQTT.client, msg)

		if testWinkMQTT.apron.LastRun != v.Cmd {
			t.Errorf("Wrong cmd - Actual: %s, Expected: %s", testWinkMQTT.apron.LastRun, v.Cmd)
		}
	}
}

func TestMqttUpdateDevice(t *testing.T) {
	var tests = []struct {
		Topic   string
		Payload string
		Cmd     string
	}{
		{"/device/1/2/update", "5", "aprontest -u -m 1 -t 2 -v 5"},
		{"/device/1/A/update", "5", ""},
		{"/device/A/A/update", "A", ""},
	}

	for _, v := range tests {
		msg := &mockMessage{
			topic:   v.Topic,
			payload: []byte(v.Payload),
		}

		testWinkMQTT.apron.LastRun = ""
		testWinkMQTT.updateDevice(testWinkMQTT.client, msg)

		if testWinkMQTT.apron.LastRun != v.Cmd {
			t.Errorf("Wrong cmd - Actual: %s, Expected: %s", testWinkMQTT.apron.LastRun, v.Cmd)
		}
	}
}

func TestMqttUpdateDeviceName(t *testing.T) {
	var tests = []struct {
		Topic   string
		Payload string
		Cmd     string
	}{
		{"/device/1/update", "TestingName", "aprontest -m 1 --set-name TestingName"},
		{"/device/1/update", "", ""},
		{"/device/A/update", "", ""},
	}

	for _, v := range tests {
		msg := &mockMessage{
			topic:   v.Topic,
			payload: []byte(v.Payload),
		}

		testWinkMQTT.apron.LastRun = ""
		testWinkMQTT.updateDeviceName(testWinkMQTT.client, msg)

		if testWinkMQTT.apron.LastRun != v.Cmd {
			t.Errorf("Wrong cmd - Actual: %s, Expected: %s", testWinkMQTT.apron.LastRun, v.Cmd)
		}
	}
}

func TestMqttUpdateStatuslightState(t *testing.T) {
	var tests = []struct {
		Topic   string
		Payload string
		Cmd     string
	}{
		{"/status_light/state/update", "ON", "set_rgb 255 255 255"},
		{"/status_light/state/update", "", ""},
	}

	for _, v := range tests {
		msg := &mockMessage{
			topic:   v.Topic,
			payload: []byte(v.Payload),
		}

		testWinkMQTT.statuslight.LastRun = ""
		testWinkMQTT.updateStatuslightState(testWinkMQTT.client, msg)

		if testWinkMQTT.statuslight.LastRun != v.Cmd {
			t.Errorf("Wrong cmd - Actual: %s, Expected: %s", testWinkMQTT.statuslight.LastRun, v.Cmd)
		}

		time.Sleep(time.Millisecond * 1000)
	}
}

func TestMqttUpdateStatuslightRGB(t *testing.T) {

	var tests = []struct {
		Topic   string
		Payload string
		Cmd     string
	}{
		{"/status_light/rgb/update", "0,255,0", "set_rgb 0 255 0"},
		{"/status_light/rgb/update", "", ""},
	}

	for _, v := range tests {
		msg := &mockMessage{
			topic:   v.Topic,
			payload: []byte(v.Payload),
		}

		testWinkMQTT.statuslight.LastRun = ""
		testWinkMQTT.updateStatuslightRGB(testWinkMQTT.client, msg)

		if testWinkMQTT.statuslight.LastRun != v.Cmd {
			t.Errorf("Wrong cmd - Actual: %s, Expected: %s", testWinkMQTT.statuslight.LastRun, v.Cmd)
		}

		time.Sleep(time.Millisecond * 1000)
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
