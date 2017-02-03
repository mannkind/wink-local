package handlers

import (
	"os/exec"
)

// BT - BT all the things!
type BT struct {
}

// Up - Bring hci0 up (the wink hub seems to turn it off sometimes)
func (t *BT) Up() {
	runnable := exec.Command("hciconfig", "hci0", "up")
	runnable.Output()
}

// FindPerson - Find a device w/l2ping
func (t *BT) FindPerson(mac string) string {
	state := "home"
	runnable := exec.Command("l2ping", "-f", "-c", "1", "-s", "1", mac)
	_, err := runnable.Output()
	if err != nil {
		state = "not_home"
	}

	return state
}
