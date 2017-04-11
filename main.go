package main

import (
	"log"

	"github.com/mannkind/wink-local/cmd"
)

// Version - Set during compilation when using included Makefile
var Version = "X.X.X"

func main() {
	log.Printf("Wink Local Version: %s", Version)
	cmd.Execute()
}
