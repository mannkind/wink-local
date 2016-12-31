package handlers

import (
	"testing"
)

func TestAddDevice(t *testing.T) {
	a := Apron{}
	var tests = []struct {
		Radio    string
		Expected string
	}{
		{"zigbee", "aprontest -a -r zigbee"},
		{"zwave", "aprontest -a -r zwave"},
		{"lutron", "aprontest -a -r lutron"},
		{"totes", ""},
	}

	for _, v := range tests {
		a.LastRun = ""
		a.AddDevice(v.Radio)

		if a.LastRun != v.Expected {
			t.Errorf("Wrong cmd - Actual: %s, Expected: %s", a.LastRun, v.Expected)
		}
	}
}

func TestAddGroup(t *testing.T) {
	a := Apron{}
	var tests = []struct {
		Name     string
		Expected string
	}{
		{"Group1", "aprontest -a -s Group1"},
		{"", ""},
	}

	for _, v := range tests {
		a.LastRun = ""
		a.AddGroup(v.Name)

		if a.LastRun != v.Expected {
			t.Errorf("Wrong cmd - Actual: %s, Expected: %s", a.LastRun, v.Expected)
		}
	}
}

func TestDeleteGroup(t *testing.T) {
	a := Apron{}
	var tests = []struct {
		ID       string
		Expected string
	}{
		{"1", "aprontest -d -w 1"},
		{"", ""},
	}

	for _, v := range tests {
		a.LastRun = ""
		a.DeleteGroup(v.ID)

		if a.LastRun != v.Expected {
			t.Errorf("Wrong cmd - Actual: %s, Expected: %s", a.LastRun, v.Expected)
		}
	}
}

func TestAddDeviceToGroup(t *testing.T) {
	a := Apron{}
	var tests = []struct {
		ID       string
		GroupID  string
		Expected string
	}{
		{"1", "2", "aprontest -a -x 2 -m 1"},
		{"2", "", ""},
		{"", "", ""},
	}

	for _, v := range tests {
		a.LastRun = ""
		a.AddDeviceToGroup(v.ID, v.GroupID)

		if a.LastRun != v.Expected {
			t.Errorf("Wrong cmd - Actual: %s, Expected: %s", a.LastRun, v.Expected)
		}
	}
}

func TestDeleteDeviceFromGroup(t *testing.T) {
	a := Apron{}
	var tests = []struct {
		ID       string
		GroupID  string
		Expected string
	}{
		{"1", "2", "aprontest -d -x 2 -m 1"},
		{"2", "", ""},
		{"", "", ""},
	}

	for _, v := range tests {
		a.LastRun = ""
		a.DeleteDeviceFromGroup(v.ID, v.GroupID)

		if a.LastRun != v.Expected {
			t.Errorf("Wrong cmd - Actual: %s, Expected: %s", a.LastRun, v.Expected)
		}
	}
}

func TestDeleteDevice(t *testing.T) {
	a := Apron{}
	var tests = []struct {
		Device   string
		Expected string
	}{
		{"1", "aprontest -d -m 1"},
		{"a", ""},
	}

	for _, v := range tests {
		a.LastRun = ""
		a.DeleteDevice(v.Device)

		if a.LastRun != v.Expected {
			t.Errorf("Wrong cmd - Actual: %s, Expected: %s", a.LastRun, v.Expected)
		}
	}
}

func TestUpdateGroup(t *testing.T) {
	a := Apron{}
	var tests = []struct {
		ID       string
		AttrID   string
		Value    string
		Expected string
	}{
		{"1", "2", "5", "aprontest -u -x 1 -t 2 -v 5"},
		{"2", "", "", ""},
		{"", "", "", ""},
	}

	for _, v := range tests {
		a.LastRun = ""
		a.UpdateGroup(v.ID, v.AttrID, v.Value)

		if a.LastRun != v.Expected {
			t.Errorf("Wrong cmd - Actual: %s, Expected: %s", a.LastRun, v.Expected)
		}
	}
}

func TestUpdateDevice(t *testing.T) {
	a := Apron{}
	var tests = []struct {
		ID       string
		AttrID   string
		Value    string
		Expected string
	}{
		{"1", "2", "5", "aprontest -u -m 1 -t 2 -v 5"},
		{"2", "", "", ""},
		{"", "", "", ""},
	}

	for _, v := range tests {
		a.LastRun = ""
		a.UpdateDevice(v.ID, v.AttrID, v.Value)

		if a.LastRun != v.Expected {
			t.Errorf("Wrong cmd - Actual: %s, Expected: %s", a.LastRun, v.Expected)
		}
	}
}

func TestUpdateDeviceName(t *testing.T) {
	a := Apron{}
	var tests = []struct {
		ID       string
		Name     string
		Expected string
	}{
		{"1", "TestingName", "aprontest -m 1 --set-name TestingName"},
		{"1", "", ""},
		{"", "", ""},
	}

	for _, v := range tests {
		a.LastRun = ""
		a.UpdateDeviceName(v.ID, v.Name)

		if a.LastRun != v.Expected {
			t.Errorf("Wrong cmd - Actual: %s, Expected: %s", a.LastRun, v.Expected)
		}
	}
}

func TestListDevices(t *testing.T) {
	a := Apron{}
	result := a.ListDevices()

	if len(result) != 0 {
		t.Errorf("Expected nothing without aprontest")
	}
}

func TestListGroups(t *testing.T) {
	a := Apron{}
	result := a.ListGroups()

	if len(result) != 0 {
		t.Errorf("Expected nothing without aprontest")
	}
}

func TestListGroupNodes(t *testing.T) {
	a := Apron{}
	result := a.ListGroupNodes("1")

	if len(result) != 0 {
		t.Errorf("Expected nothing without aprontest")
	}
}
