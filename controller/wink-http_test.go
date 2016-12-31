package controller

import (
	"fmt"
	"github.com/abbot/go-http-auth"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var testAuthenticatorUsername = "root"
var testAuthenticatorPassword = "hello"
var testAuthenticator = func() *auth.BasicAuth {
	a := auth.NewBasicAuthenticator("winkhub", func(user, realm string) string {
		return "$1$dlPL2MqE$oQmn16q49SqdmhenQuNgs1" // hello
	})

	return a
}()

var testWinkHTTP = func() *winkHTTP {
	myCtrl := winkHTTP{
		settings: struct {
			Port     int16
			Username string
			Password string
		}{
			Port: 8080,
		},
	}

	myCtrl.start()
	return &myCtrl
}()

var testRouter = func() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/device/add", testAuthenticator.Wrap(testWinkHTTP.addDevice))
	r.HandleFunc("/device/delete", testAuthenticator.Wrap(testWinkHTTP.deleteDevice))
	r.HandleFunc("/device/list", testAuthenticator.Wrap(testWinkHTTP.listDevice))
	r.HandleFunc("/device/{id}/update", testAuthenticator.Wrap(testWinkHTTP.updateDeviceName))
	r.HandleFunc("/group/add", testAuthenticator.Wrap(testWinkHTTP.addGroup))
	r.HandleFunc("/group/delete", testAuthenticator.Wrap(testWinkHTTP.deleteGroup))
	r.HandleFunc("/group/list", testAuthenticator.Wrap(testWinkHTTP.listGroup))
	r.HandleFunc("/group/{id}/add", testAuthenticator.Wrap(testWinkHTTP.addDeviceToGroup))
	r.HandleFunc("/group/{id}/delete", testAuthenticator.Wrap(testWinkHTTP.deleteDeviceFromGroup))
	r.HandleFunc("/status_light/state/update", testAuthenticator.Wrap(testWinkHTTP.updateStatuslightState))
	r.HandleFunc("/status_light/rgb/update", testAuthenticator.Wrap(testWinkHTTP.updateStatuslightRGB))

	return r
}()

func TestAddDevice(t *testing.T) {
	var tests = []struct {
		URL  string
		Body string
		Cmd  string
	}{
		{"/device/add", "zigbee", "aprontest -a -r zigbee"},
		{"/device/add", "zwave", "aprontest -a -r zwave"},
		{"/device/add", "lutron", "aprontest -a -r lutron"},
		{"/device/add", "totes", ""},
	}

	for _, v := range tests {
		body := strings.NewReader(fmt.Sprintf("{ \"Radio\": \"%s\" }", v.Body))
		req, err := http.NewRequest("POST", v.URL, body)
		req.SetBasicAuth(testAuthenticatorUsername, testAuthenticatorPassword)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		testWinkHTTP.apron.LastRun = ""
		testRouter.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v, %s", status, http.StatusOK, rr.Body)
		}

		if testWinkHTTP.apron.LastRun != v.Cmd {
			t.Errorf("Wrong cmd - Actual: %s, Expected: %s", testWinkMQTT.apron.LastRun, v.Cmd)
		}
	}
}

func TestDeleteDevice(t *testing.T) {
	var tests = []struct {
		URL  string
		Body string
		Cmd  string
	}{
		{"/device/delete", "1", "aprontest -d -m 1"},
		{"/device/delete", "A", ""},
	}

	for _, v := range tests {
		body := strings.NewReader(fmt.Sprintf("{ \"ID\": \"%s\" }", v.Body))
		req, err := http.NewRequest("POST", v.URL, body)
		req.SetBasicAuth(testAuthenticatorUsername, testAuthenticatorPassword)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		testWinkHTTP.apron.LastRun = ""
		testRouter.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v, %s", status, http.StatusOK, rr.Body)
		}

		if testWinkHTTP.apron.LastRun != v.Cmd {
			t.Errorf("Wrong cmd - Actual: %s, Expected: %s", testWinkMQTT.apron.LastRun, v.Cmd)
		}
	}
}

func TestUpdateDeviceName(t *testing.T) {
	var tests = []struct {
		URL  string
		Body string
		Cmd  string
	}{
		{"/device/1/update", "TestingName", "aprontest -m 1 --set-name TestingName"},
		{"/device/1/update", "", ""},
		{"/device/A/update", "", ""},
	}

	for _, v := range tests {
		body := strings.NewReader(fmt.Sprintf("{ \"Name\": \"%s\" }", v.Body))
		req, err := http.NewRequest("POST", v.URL, body)
		req.SetBasicAuth(testAuthenticatorUsername, testAuthenticatorPassword)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		testWinkHTTP.apron.LastRun = ""
		testRouter.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v, %s", status, http.StatusOK, rr.Body)
		}

		if testWinkHTTP.apron.LastRun != v.Cmd {
			t.Errorf("Wrong cmd - Actual: %s, Expected: %s", testWinkMQTT.apron.LastRun, v.Cmd)
		}
	}
}

func TestAddGroup(t *testing.T) {
	var tests = []struct {
		URL  string
		Body string
		Cmd  string
	}{
		{"/group/add", "Group1", "aprontest -a -s Group1"},
		{"/group/add", "", ""},
	}

	for _, v := range tests {
		body := strings.NewReader(fmt.Sprintf("{ \"Name\": \"%s\" }", v.Body))
		req, err := http.NewRequest("POST", v.URL, body)
		req.SetBasicAuth(testAuthenticatorUsername, testAuthenticatorPassword)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		testWinkHTTP.apron.LastRun = ""
		testRouter.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v, %s", status, http.StatusOK, rr.Body)
		}

		if testWinkHTTP.apron.LastRun != v.Cmd {
			t.Errorf("Wrong cmd - Actual: %s, Expected: %s", testWinkMQTT.apron.LastRun, v.Cmd)
		}
	}
}

func TestDeleteGroup(t *testing.T) {
	var tests = []struct {
		URL  string
		Body string
		Cmd  string
	}{
		{"/group/delete", "1", "aprontest -d -w 1"},
		{"/group/delete", "A", ""},
	}

	for _, v := range tests {
		body := strings.NewReader(fmt.Sprintf("{ \"ID\": \"%s\" }", v.Body))
		req, err := http.NewRequest("POST", v.URL, body)
		req.SetBasicAuth(testAuthenticatorUsername, testAuthenticatorPassword)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		testWinkHTTP.apron.LastRun = ""
		testRouter.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v, %s", status, http.StatusOK, rr.Body)
		}

		if testWinkHTTP.apron.LastRun != v.Cmd {
			t.Errorf("Wrong cmd - Actual: %s, Expected: %s", testWinkMQTT.apron.LastRun, v.Cmd)
		}
	}
}

func TestAddDeviceToGroup(t *testing.T) {
	var tests = []struct {
		URL  string
		Body string
		Cmd  string
	}{
		{"/group/2/add", "1", "aprontest -a -x 2 -m 1"},
		{"/group/2/add", "", ""},
		{"/group/A/add", "", ""},
	}

	for _, v := range tests {
		body := strings.NewReader(fmt.Sprintf("{ \"DeviceID\": \"%s\" }", v.Body))
		req, err := http.NewRequest("POST", v.URL, body)
		req.SetBasicAuth(testAuthenticatorUsername, testAuthenticatorPassword)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		testWinkHTTP.apron.LastRun = ""
		testRouter.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v, %s", status, http.StatusOK, rr.Body)
		}

		if testWinkHTTP.apron.LastRun != v.Cmd {
			t.Errorf("Wrong cmd - Actual: %s, Expected: %s", testWinkMQTT.apron.LastRun, v.Cmd)
		}
	}
}

func TestDeleteDeviceFromGroup(t *testing.T) {
	var tests = []struct {
		URL  string
		Body string
		Cmd  string
	}{
		{"/group/2/delete", "1", "aprontest -d -x 2 -m 1"},
		{"/group/2/delete", "", ""},
		{"/group/A/delete", "", ""},
	}

	for _, v := range tests {
		body := strings.NewReader(fmt.Sprintf("{ \"DeviceID\": \"%s\" }", v.Body))
		req, err := http.NewRequest("POST", v.URL, body)
		req.SetBasicAuth(testAuthenticatorUsername, testAuthenticatorPassword)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		testWinkHTTP.apron.LastRun = ""
		testRouter.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v, %s", status, http.StatusOK, rr.Body)
		}

		if testWinkHTTP.apron.LastRun != v.Cmd {
			t.Errorf("Wrong cmd - Actual: %s, Expected: %s", testWinkMQTT.apron.LastRun, v.Cmd)
		}
	}
}
