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
	LastRun     string
	LastRunTime map[string]time.Time
}

// ApronDeviceGroup - A device/group listed in the Apron database
type ApronDeviceGroup struct {
	ID    int64
	Name  string
	Nodes []ApronDeviceGroup
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

// DeleteDeviceFromGroup - Delete a device from a group
func (t *Apron) DeleteDeviceFromGroup(id string, groudID string) {
	if _, err := strconv.Atoi(id); err != nil {
		log.Print("Can't add device to group; id is not a valid number")
		return
	}

	if _, err := strconv.Atoi(groudID); err != nil {
		log.Print("Can't add device to group; groudID is not a valid number")
		return
	}

	t.run([]string{"-d", "-x", groudID, "-m", id})
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
		log.Print("Can't update device; attr is not a valid number")
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

// ListDevices - Gets a list of devices
func (t *Apron) ListDevices() []ApronDeviceGroup {
	args := []string{"-l"}
	runnable := exec.Command("aprontest", args...)
	stdout, err := runnable.Output()
	if err != nil {
		log.Println(err)
		stdout = []byte{}
	}

	devices := []ApronDeviceGroup{}
	lines := strings.Split(string(stdout), "\n")
	for i, line := range lines {
		// Skip the first two lines of the output
		if i < 2 {
			continue
		}

		// Check for correct number of fields
		deviceInfo := strings.Split(line, "|")
		if len(deviceInfo) < 2 {
			continue
		}

		// Parse the id, and select the name
		id, err := strconv.ParseInt(strings.TrimSpace(deviceInfo[0]), 10, 64)
		if err != nil {
			break
		}
		name := strings.TrimSpace(deviceInfo[2])
		devices = append(devices, ApronDeviceGroup{ID: id, Name: name})
	}

	return devices
}

// ListGroups - Gets a list of groups
func (t *Apron) ListGroups() []ApronDeviceGroup {
	args := []string{"-l"}
	runnable := exec.Command("aprontest", args...)
	stdout, err := runnable.Output()
	if err != nil {
		log.Println(err)
		stdout = []byte{}
	}

	found := false
	groups := []ApronDeviceGroup{}
	lines := strings.Split(string(stdout), "\n")
	for _, line := range lines {
		// Wait until we've found the groups
		if strings.Contains(line, "master groups") {
			found = true
			continue
		} else if !found {
			continue
		}

		// Check for correct number of fields
		deviceInfo := strings.Split(line, "|")
		if len(deviceInfo) < 2 {
			break
		}

		// Skip GROUP ID line
		deviceID := strings.TrimSpace(deviceInfo[0])
		if deviceID == "GROUP ID" {
			continue
		}

		// Parse the id, and select the name
		id, err := strconv.ParseInt(deviceID, 10, 64)
		if err != nil {
			break
		}
		name := strings.TrimSpace(deviceInfo[1])

		// Remember to get the nodes for the group
		group := ApronDeviceGroup{
			ID:    id,
			Name:  name,
			Nodes: t.ListGroupNodes(deviceID),
		}
		groups = append(groups, group)
	}

	return groups
}

// ListGroupNodes - Gets a list of nodes in a group
func (t *Apron) ListGroupNodes(id string) []ApronDeviceGroup {
	nodes := []ApronDeviceGroup{}
	args := []string{"-l", "-x", id}
	runnable := exec.Command("aprontest", args...)

	stdout, err := runnable.Output()
	if err != nil {
		log.Println(err)
		stdout = []byte{}
	}

	devices := t.ListDevices()
	found := false
	lines := strings.Split(string(stdout), "\n")
	for _, line := range lines {
		// Wait until we've found the nodes for the group
		if strings.Contains(line, "nodes for master group") {
			found = true
			continue
		} else if !found {
			continue
		}

		// Check for correct number of fields
		deviceInfo := strings.Split(line, "|")
		if len(deviceInfo) < 2 {
			break
		}

		// Skip GROUP ID line
		deviceID := strings.TrimSpace(deviceInfo[0])
		if deviceID == "GROUP ID" {
			continue
		}

		// Parse the id
		id, err := strconv.ParseInt(strings.TrimSpace(deviceInfo[1]), 10, 64)
		if err != nil {
			break
		}

		// Find the device name in known devices
		name := ""
		for _, device := range devices {
			if device.ID == id {
				name = device.Name
				break
			}
		}
		nodes = append(nodes, ApronDeviceGroup{ID: id, Name: name})
	}

	return nodes
}

func (t *Apron) deleteBoth(isGroup bool, id string) {
	argXW := "-m"
	if isGroup {
		argXW = "-w"
	}
	t.run([]string{"-d", argXW, id})
}

func (t *Apron) updateBoth(isGroup bool, id string, attr string, value string) bool {
	topic := fmt.Sprintf("%t/%s/%s", isGroup, id, attr)
	if t.LastRunTime == nil {
		t.LastRunTime = make(map[string]time.Time)
	}
	if time.Since(t.LastRunTime[topic]) < time.Millisecond*1000 {
		return false
	}

	t.LastRunTime[topic] = time.Now()

	argXM := "-m"
	if isGroup {
		argXM = "-x"
	}

	t.run([]string{"-u", argXM, id, "-t", attr, "-v", value})

	return true
}

func (t *Apron) run(args []string) {
	runnable := exec.Command("aprontest", args...)
	t.LastRun = fmt.Sprintf("aprontest %s", strings.Join(args, " "))
	log.Printf("Running %s", t.LastRun)

	go func(runnable *exec.Cmd) {
		if err := runnable.Run(); err != nil {
			log.Printf("... %s", err)
		}
	}(runnable)
}
