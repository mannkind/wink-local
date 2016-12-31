# Wink Local

[![Software
License](https://img.shields.io/badge/License-MIT-orange.svg?style=flat-square)](https://github.com/mannkind/wink-local/blob/master/LICENSE.md)
[![Travis CI](https://img.shields.io/travis/mannkind/wink-local/master.svg?style=flat-square)](https://travis-ci.org/mannkind/wink-local)
[![Coverage Status](https://img.shields.io/codecov/c/github/mannkind/wink-local/master.svg)](http://codecov.io/github/mannkind/wink-local?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/mannkind/wink-local)](https://goreportcard.com/report/github.com/mannkind/wink-local)

A local-control replacement for the Wink Hub that utilizes MQTT

# Build

* ./build.sh

# Installation

```
ssh winkhub "mkdir -p /opt/wink-local" 
scp -r web/dist winkhub:/opt/wink-local
cat wfs/etc/monitrc | ssh winkhub "cat >> /etc/monitrc"
cat wfs/etc/rc.d/init.d/wink-local | ssh winkhub "cat >> /etc/rc.d/init.d/wink-local"
scp wink-local winkhub:/opt/wink-local
scp wfs/opt/wink-local/wink-local.yaml winkhub:/opt/wink-local/
scp wfs/etc/rc.d/init.d/wink-local winkhub:/etc/rc.d/init.d/wink-local
ssh winkhub "/etc/rc.d/init.d/wink-local start"
```

# Configuration

Configuration happens in the /opt/wink-local/wink-local.yaml file. An example might look this:

```
http:
    port: 8080
mqtt:
    clientid: 'WinkLocal'
    broker:   'tcp://mosquitto:1883'
    topicbase: 'winkhub'
    retain: true
```

# Dependencies

* [Go](https://golang.org)
* [UPX](https://upx.github.io)
* An MQTT broker, such as [Mosquitto](https://mosquitto.org)
