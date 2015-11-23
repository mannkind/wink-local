package handlers

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// Sqlite3 - Query the aprondb
type Sqlite3 struct {
	Zigbee string
	Zwave  string
	Lutron string

	db     string
	locked bool
}

// States - Return the current state of the aprondb
func (t *Sqlite3) States() ([]string, error) {
	t.checkDefaults()

	root := "/tmp/database"
	localDB := t.db
	filepath.Walk(root, func(path string, f os.FileInfo, err error) error {
		if path == root || f == nil {
			return nil
		}

		dbStat, err := os.Stat(localDB)
		if f.ModTime().Sub(dbStat.ModTime()) > 0 {
			localDB = path
		}

		return err
	})

	if localDB != t.db {
		log.Printf("Using %s as the database", localDB)
		t.db = localDB
	}

	if !t.locked {
		t.locked = true

		sqlQuery := fmt.Sprintf("select d.masterId, s.attributeId, s.value_GET FROM zigbeeDeviceState AS s,zigbeeDevice AS d WHERE d.globalId=s.globalId AND s.attributeId IN (%s) UNION select d.masterId, s.attributeId, s.value_SET FROM zwaveDeviceState AS s,zwaveDevice AS d WHERE d.nodeId=s.nodeId AND s.attributeId IN (%s) UNION select d.masterId, s.attributeId, s.value_SET FROM lutronDeviceState AS s,lutronDevice AS d WHERE d.lNodeId=s.lNodeId AND s.attributeId IN (%s)", t.Zigbee, t.Zwave, t.Lutron)
		args := []string{"-csv", t.db, sqlQuery}
		runnable := exec.Command("sqlite3", args...)
		stdout, err := runnable.Output()
		if err != nil {
			log.Println(err)
			return []string{}, err
		}
		t.locked = false
		return strings.Split(string(stdout), "\n"), nil
	}
	return []string{}, nil
}

func (t *Sqlite3) checkDefaults() {
	if t.db == "" {
		t.db = "/database/apron.db"
	}

	if t.Zigbee == "" {
		t.Zigbee = "1,2"
	}

	if t.Zwave == "" {
		t.Zwave = "2,3,7,8"
	}

	if t.Lutron == "" {
		t.Lutron = "1"
	}
}
