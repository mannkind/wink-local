package main

import (
	"fmt"
	MQTT "git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git"
	"math"
	"os/exec"
	"strconv"
	"strings"
)

func runCmd(cmd string) {
	fmt.Printf("Command: %s\n", cmd)
	err := exec.Command("sh", "-c", cmd).Run()
	check(err, "Unable to run the command", false)
}

func initMQTT() {
	fmt.Println("Connecting to MQTT")
	opts := MQTT.NewClientOptions().AddBroker("tcp://localhost:1883")
	opts.SetClientID("WinkLights")
	opts.SetDefaultPublishHandler(func(client *MQTT.Client, msg MQTT.Message) {
		fmt.Printf("Received topic '%s' with payload '%s'\n", msg.Topic(), msg.Payload())

		// Make sure the SSH connection is up
		cmd := "ssh winkhub -O check || { ssh winkhub -N & }"
		runCmd(cmd)

		s := strings.Split(msg.Topic(), "/")
		if s[1] == "light_bulb" {
			go handleLights(msg.Topic(), string(msg.Payload()))
		}
	})

	// Create and start a client using the above ClientOptions
	mqtt_client := MQTT.NewClient(opts)
	connect_token := mqtt_client.Connect()
	connect_token.Wait()
	check(connect_token.Error(), "Unable to connect to MQTT broker", true)

	// Subscribe to topic
	subscribe_token := mqtt_client.Subscribe("winknet/#", 0, nil)
	subscribe_token.Wait()
	check(subscribe_token.Error(), "Unable to subscribe to MQTT topic", true)
	fmt.Println("Connected to MQTT")
}

func handleLights(topic string, payload string) {
	var onoff = func(device_id string, device_action string) {
		device_action = strings.ToUpper(device_action)

		cmd := fmt.Sprintf("ssh winkhub \"/usr/sbin/aprontest -m %s -t 1 -v %s -u\"", device_id, device_action)
		runCmd(cmd)
	}

	var dim_to_percent = func(device_id string, payload string) {
		percent, err := strconv.ParseFloat(payload, 10)
		if check(err, "Unable to parse the value from dim_to_percent into a percentage", false) {
			return
		}

		device_value := fmt.Sprintf("%.0f", math.Ceil(255.0*(percent/100.0)))
		cmd := fmt.Sprintf("ssh winkhub \"/usr/sbin/aprontest -m %s -t 2 -v %s -u\"", device_id, device_value)
		runCmd(cmd)
	}

	s := strings.Split(topic, "/")
	device_id := s[2]
	device_action := s[3]

	if device_action == "on" || (device_action == "dim_to_percent" && payload == "ON") {
		go onoff(device_id, "ON")
		go dim_to_percent(device_id, "100")
	} else if device_action == "off" || (device_action == "dim_to_percent" && payload == "OFF") {
		go onoff(device_id, "OFF")
	} else if device_action == "dim_to_percent" {
		go dim_to_percent(device_id, payload)
	}
}

// Check, log, panic!
func check(e error, msg string, shouldPanic bool) bool {
	if e != nil {
		fmt.Printf("ERROR: %s\n", msg)
		if shouldPanic {
			panic(e)
		}
	}

	return e != nil
}

func main() {
	initMQTT()

	// Wait forever
	select {}
}
