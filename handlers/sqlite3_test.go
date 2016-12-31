package handlers

import (
	"os/exec"
	"testing"
)

func TestStatesFake(t *testing.T) {
	runnable := exec.Command("bash", "-c", "echo 1")
	runnable.Run()
}

/*
@TODO: Figure out how to make this work on TravisCI
func TestStatesError(t *testing.T) {
	runnable := exec.Command("bash", "-c", "rm -f /tmp/database && mkdir -p /tmp/database")
	runnable.Run()

	sqlite3 := Sqlite3{db: "/tmp/database/apron.db"}
	if results, err := sqlite3.States(); err != nil && len(results) > 0 {
		t.Errorf("Sqlite3 returned results unexpectedly")
	}
}

func TestStates(t *testing.T) {
	runnable := exec.Command("bash", "-c", "mkdir -p /tmp/database; cp $GOHOME/src/github.com/mannkind/wink-local/tests/apron.db /tmp/database")
	runnable.Run()

	sqlite3 := Sqlite3{db: "/tmp/database/apron.db"}
	if _, err := sqlite3.States(); err != nil {
		t.Errorf("Sqlite3 errored unexpectedly")
	}
}
*/
