package handlers

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

// Apron - Apron all the things!
type Apron struct {
	LastRun  string
	Debounce map[string]bool
}

// AddDevice - Add a device to the hub
func (t *Apron) AddDevice(radio string) {
	radios := map[string]bool{"lutron": true, "zwave": true, "zigbee": true, "kidde": true}
	if _, ok := radios[radio]; !ok {
		log.Print("Can't add device; not a valid radio")
		return
	}

	t.run([]string{"-a", "-r", radio})
}

// AddGroup - Add a group to the hub
func (t *Apron) AddGroup(name string) {
	if len(name) == 0 {
		return
	}

	t.run([]string{"-a", "-s", name})
}

// DeleteGroup - Delete a group from the hub
func (t *Apron) DeleteGroup(id string) {
	if _, err := strconv.Atoi(id); err != nil {
		log.Print("Can't delete group; id is not a valid number")
		return
	}
	t.deleteBoth(true, id)
}

// AddDeviceToGroup - Add a device to a group
func (t *Apron) AddDeviceToGroup(id string, groudID string) {
	if _, err := strconv.Atoi(id); err != nil {
		log.Print("Can't add device to group; id is not a valid number")
		return
	}

	if _, err := strconv.Atoi(groudID); err != nil {
		log.Print("Can't add device to group; groudID is not a valid number")
		return
	}

	t.run([]string{"-a", "-x", groudID, "-m", id})
}

// DeleteDevice - Delete a device from the hub
func (t *Apron) DeleteDevice(id string) {
	if _, err := strconv.Atoi(id); err != nil {
		log.Print("Can't delete device; id is not a valid number")
		return
	}
	t.deleteBoth(false, id)
}

// UpdateGroup - Update a group on the hub
func (t *Apron) UpdateGroup(id string, attr string, value string) bool {
	if _, err := strconv.Atoi(id); err != nil {
		log.Print("Can't update group; id is not a valid number")
		return false
	}

	if _, err := strconv.Atoi(attr); err != nil {
		log.Print("Can't update group; attr is not a valid number")
		return false
	}

	return t.updateBoth(true, id, attr, value)
}

// UpdateDevice - Update a device on the hub
func (t *Apron) UpdateDevice(id string, attr string, value string) bool {
	if _, err := strconv.Atoi(id); err != nil {
		log.Print("Can't update device; id is not a valid number")
		return false
	}

	if _, err := strconv.Atoi(attr); err != nil {
		log.Print("Can't update drvice; attr is not a valid number")
		return false
	}

	return t.updateBoth(false, id, attr, value)
}

// UpdateDeviceName - Update the name of a device
func (t *Apron) UpdateDeviceName(id string, name string) {
	if _, err := strconv.Atoi(id); err != nil {
		log.Print("Can't update device name; id is not a valid number")
		return
	}

	if len(name) == 0 {
		log.Print("Can't update device name; no name provided")
		return
	}

	t.run([]string{"-m", id, "--set-name", name})
}

func (t *Apron) deleteBoth(isGroup bool, id string) {
	argXW := "-m"
	if isGroup {
		argXW = "-w"
	}
	t.run([]string{"-d", argXW, id})
}

func (t *Apron) updateBoth(isGroup bool, id string, attr string, value string) bool {
	topic := fmt.Sprintf("%s/%s", id, attr)
	if t.Debounce == nil {
		t.Debounce = make(map[string]bool)
	}
	if t.Debounce[topic] {
		return false
	}

	argXM := "-m"
	if isGroup {
		argXM = "-x"
	}

	t.run([]string{"-u", argXM, id, "-t", attr, "-v", value})

	t.Debounce[topic] = true
	debounceTimer := time.NewTimer(time.Millisecond * 1000)
	go func() {
		<-debounceTimer.C
		t.Debounce[topic] = false
	}()

	return true
}

func (t *Apron) run(args []string) {
	runnable := exec.Command("aprontest", args...)
	t.LastRun = fmt.Sprintf("aprontest %s", strings.Join(args, " "))
	log.Printf("Running %s", t.LastRun)

	go func(cmd *exec.Cmd) {
		if err := cmd.Run(); err != nil {
			log.Printf("... %s", err)
		}
	}(runnable)
}
