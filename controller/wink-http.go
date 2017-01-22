package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/abbot/go-http-auth"
	"github.com/gorilla/mux"
	"github.com/mannkind/wink-local/handlers"
)

type winkHTTP struct {
	settings struct {
		Port     int16
		Username string
		Password string
	}
	apron       handlers.Apron
	statuslight handlers.RGB
}

func (t *winkHTTP) start() error {
	t.checkDefaults()

	log.Printf("Setting up to HTTP: %d", t.settings.Port)
	a := auth.NewBasicAuthenticator("winkhub", func(user, realm string) string {
		if user == t.settings.Username {
			return t.settings.Password
		}

		return ""
	})

	r := mux.NewRouter()
	r.HandleFunc("/device/add", a.Wrap(t.addDevice))
	r.HandleFunc("/device/delete", a.Wrap(t.deleteDevice))
	r.HandleFunc("/device/list", a.Wrap(t.listDevice))
	r.HandleFunc("/device/{id}/update", a.Wrap(t.updateDeviceName))
	r.HandleFunc("/group/add", a.Wrap(t.addGroup))
	r.HandleFunc("/group/delete", a.Wrap(t.deleteGroup))
	r.HandleFunc("/group/list", a.Wrap(t.listGroup))
	r.HandleFunc("/group/{id}/add", a.Wrap(t.addDeviceToGroup))
	r.HandleFunc("/group/{id}/delete", a.Wrap(t.deleteDeviceFromGroup))
	r.HandleFunc("/status_light/state/update", a.Wrap(t.updateStatuslightState))
	r.HandleFunc("/status_light/rgb/update", a.Wrap(t.updateStatuslightRGB))
	r.HandleFunc("/{x:.*}", a.Wrap(func(res http.ResponseWriter, req *auth.AuthenticatedRequest) {
		http.FileServer(http.Dir("/opt/wink-local/dist")).ServeHTTP(res, &req.Request)
	}))

	http.Handle("/", r)

	go func() {
		if err := http.ListenAndServe(fmt.Sprintf(":%d", t.settings.Port), nil); err != nil {
			log.Printf("Failed to setup HTTP: %s", err)
		}
	}()

	log.Print("Setup HTTP")

	return nil
}

func (t *winkHTTP) checkDefaults() {
	if t.settings.Port == 0 {
		t.settings.Port = 8080
	}

	if len(t.settings.Username) == 0 {
		t.settings.Username = "root"
	}

	if len(t.settings.Password) == 0 {
		t.settings.Password = "$1$/5Ehiv3F$qWqGG6V9aE5rNbfq6WV4.0"
	}
}

func (t *winkHTTP) addDevice(w http.ResponseWriter, r *auth.AuthenticatedRequest) {
	var jsonData struct {
		Radio string
	}

	err := json.NewDecoder(r.Body).Decode(&jsonData)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	t.apron.AddDevice(jsonData.Radio)
}

func (t *winkHTTP) listDevice(w http.ResponseWriter, r *auth.AuthenticatedRequest) {
	devices := t.apron.ListDevices()

	json.NewEncoder(w).Encode(devices)
}

func (t *winkHTTP) deleteDevice(w http.ResponseWriter, r *auth.AuthenticatedRequest) {
	var jsonData struct {
		ID string
	}

	err := json.NewDecoder(r.Body).Decode(&jsonData)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	t.apron.DeleteDevice(jsonData.ID)
}

func (t *winkHTTP) updateDeviceName(w http.ResponseWriter, r *auth.AuthenticatedRequest) {
	vars := mux.Vars(&r.Request)
	id := vars["id"]

	var jsonData struct {
		Name string
	}

	err := json.NewDecoder(r.Body).Decode(&jsonData)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	t.apron.UpdateDeviceName(id, jsonData.Name)
}

func (t *winkHTTP) addGroup(w http.ResponseWriter, r *auth.AuthenticatedRequest) {
	var jsonData struct {
		Name string
	}

	err := json.NewDecoder(r.Body).Decode(&jsonData)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	t.apron.AddGroup(jsonData.Name)
}

func (t *winkHTTP) listGroup(w http.ResponseWriter, r *auth.AuthenticatedRequest) {
	groups := t.apron.ListGroups()

	json.NewEncoder(w).Encode(groups)
}

func (t *winkHTTP) addDeviceToGroup(w http.ResponseWriter, r *auth.AuthenticatedRequest) {
	vars := mux.Vars(&r.Request)
	id := vars["id"]

	var jsonData struct {
		DeviceID string
	}

	err := json.NewDecoder(r.Body).Decode(&jsonData)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	t.apron.AddDeviceToGroup(jsonData.DeviceID, id)
}

func (t *winkHTTP) deleteDeviceFromGroup(w http.ResponseWriter, r *auth.AuthenticatedRequest) {
	vars := mux.Vars(&r.Request)
	id := vars["id"]

	var jsonData struct {
		DeviceID string
	}

	err := json.NewDecoder(r.Body).Decode(&jsonData)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	t.apron.DeleteDeviceFromGroup(jsonData.DeviceID, id)
}

func (t *winkHTTP) deleteGroup(w http.ResponseWriter, r *auth.AuthenticatedRequest) {
	var jsonData struct {
		ID string
	}

	err := json.NewDecoder(r.Body).Decode(&jsonData)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	t.apron.DeleteGroup(jsonData.ID)
}

func (t *winkHTTP) updateStatuslightState(w http.ResponseWriter, r *auth.AuthenticatedRequest) {
	var jsonData struct {
		State string
	}

	err := json.NewDecoder(r.Body).Decode(&jsonData)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	if jsonData.State == "ON" {
		t.statuslight.Update("255,255,255")
	} else if jsonData.State == "OFF" {
		t.statuslight.Update("0,0,0")
	}
}

func (t *winkHTTP) updateStatuslightRGB(w http.ResponseWriter, r *auth.AuthenticatedRequest) {
	var jsonData struct {
		Color string
	}

	err := json.NewDecoder(r.Body).Decode(&jsonData)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	parts := strings.Split(jsonData.Color, " ")
	if len(parts) == 3 {
		t.statuslight.Flash(parts[0], parts[1], parts[2])
	} else {
		t.statuslight.Update(parts[0])
	}
}
