# Wink MQTT

[![Software
License](https://img.shields.io/badge/License-MIT-orange.svg?style=flat-square)](https://github.com/mannkind/wink-mqtt/blob/master/LICENSE.md)
[![Travis CI](https://img.shields.io/travis/mannkind/wink-mqtt/master.svg?style=flat-square)](https://travis-ci.org/mannkind/wink-mqtt)
[![Coverage Status](https://img.shields.io/codecov/c/github/mannkind/wink-mqtt/master.svg)](http://codecov.io/github/mannkind/wink-mqtt?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/mannkind/wink-mqtt)](https://goreportcard.com/report/github.com/mannkind/wink-mqtt)

A local-control replacement for the Wink Hub that utilizes MQTT

# Build

* ./build.sh

# Installation

* ssh winkhub "mkdir -p /opt/wink-mqtt" 
* cat wfs/etc/monitrc | ssh winkhub "cat >> /etc/monitrc"
* scp wink-mqtt winkhub:/opt/wink-mqtt
* scp wfs/opt/wink-mqtt/wink-mqtt.yaml winkhub:/opt/wink-mqtt/
* scp wfs/etc/rc.d/init.d/wink-mqtt winkhub:/etc/rc.d/init.d/wink-mqtt
* ssh winkhub "/etc/rc.d/init.d/wink-mqtt start"

# Configuration

Configuration happens in the /opt/wink-mqtt/wink-mqtt.yaml file. An example might look this:

```
mqtt:
    clientid: 'WinkMQTT'
    broker:   'tcp://mosquitto:1883'
    topicbase: 'winkhub'
    retain: true
```

# Dependencies

* [Go](https://golang.org)
* [UPX](https://upx.github.io)
* An MQTT broker, such as [Mosquitto](https://mosquitto.org)
