# Wink MQTT

Control a rooted Winkhub via MQTT. Currently only works for lights.

## Installation

* `go get`
* `go build`
* setup an entry in your ~/.ssh/config, the ssh key used needs to be passwordless
r
    Host winkhub
        HostName <hostname or ip>
        User root
        ControlMaster auto
        ControlPath /tmp/ssh_mux_%h_%p_%r
        IdentityFile ~/.ssh/id_dsa
        IdentitiesOnly yes
