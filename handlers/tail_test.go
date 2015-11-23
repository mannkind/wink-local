package handlers

import (
	"os/exec"
	"testing"
)

func TestTailDNE(t *testing.T) {
	runnable := exec.Command("rm", "-f", "/tmp/all.log")
	runnable.Run()

	c := make(chan bool)
	tl := Tail{}
	go tl.Start(c)

	v := <-c
	if v {
		t.Errorf("Tail didn't return false for a file that does not exist.")
	}
}

func TestTail(t *testing.T) {
	runnable := exec.Command("bash", "-c", "echo 'state changed in device' > /tmp/all.log")
	runnable.Run()

	c := make(chan bool)
	tl := Tail{}
	go tl.Start(c)

	v := <-c
	if !v {
		t.Errorf("Tail didn't communicate true")
	}
}
