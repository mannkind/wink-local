package handlers

import (
	"errors"
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"
)

// RGB - RGB all the things!
type RGB struct {
	LastRun     string
	LastRunTime time.Time
}

// Flash - Add a device to the hub
func (t *RGB) Flash(color string, alternative string, microseconds string) {
	r, g, b, err := t.colorParse(color)
	r2, g2, b2, err2 := t.colorParse(alternative)
	if err != nil || err2 != nil {
		log.Printf("Cannot flash RGB to %s/%s/%s - Not valid colors", color, alternative, microseconds)
		return
	}

	t.run([]string{fmt.Sprintf("%d", r), fmt.Sprintf("%d", g), fmt.Sprintf("%d", b), fmt.Sprintf("%d", r2), fmt.Sprintf("%d", g2), fmt.Sprintf("%d", b2), "flash", microseconds})
}

// Update - Add a device to the hub
func (t *RGB) Update(color string) bool {
	if time.Since(t.LastRunTime) < time.Millisecond*1000 {
		return false
	}

	t.LastRunTime = time.Now()

	r, g, b, err := t.colorParse(color)
	if err != nil {
		log.Printf("Cannot update RGB to %s - Not valid colors", color)
		return false
	}

	t.run([]string{fmt.Sprintf("%d", r), fmt.Sprintf("%d", g), fmt.Sprintf("%d", b)})

	return true
}

func (t *RGB) colorParse(color string) (uint8, uint8, uint8, error) {
	if len(color) == 0 {
		return 0, 0, 0, errors.New("No color available")
	}
	format := "%d,%d,%d"
	if color[0] == "#"[0] {
		format = "#%02x%02x%02x"
		if len(color) == 4 {
			format = "#%1x%1x%1x"
		}
	}

	var r, g, b uint8
	_, err := fmt.Sscanf(color, format, &r, &g, &b)
	if err != nil {
		log.Printf("Error parsing color %s", color)
		return 0, 0, 0, err
	}

	return r, g, b, nil
}

func (t *RGB) run(args []string) {
	runnable := exec.Command("set_rgb", args...)
	t.LastRun = fmt.Sprintf("set_rgb %s", strings.Join(args, " "))
	log.Printf("Running %s", t.LastRun)

	go func(runnable *exec.Cmd) {
		if err := runnable.Run(); err != nil {
			log.Printf("... %s", err)
		}
	}(runnable)
}
