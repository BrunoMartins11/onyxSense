package main

import (
	"encoding/json"
	"github.com/BrunoMartins11/onyxSense/internal/model"
	"net/http"
)

func CreateRoomHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	roomName := req.URL.Query().Get("RoomName")

	err := manager.RegisterNewRoom(model.Room{Name: roomName})
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func GetRooms(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if req.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	data := manager.GetAllRooms()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func CreateSensorHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	roomName := req.URL.Query().Get("RoomName")
	sensorName := req.URL.Query().Get("SensorName")
	if roomName == "" || sensorName == "" {
		w.WriteHeader(http.StatusBadRequest)
	}

	err := manager.RegisterNewSensor(model.Sensor{Name: sensorName}, roomName)
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func GetRoomActivePresences(w http.ResponseWriter, req *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if req.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	roomName := req.URL.Query().Get("RoomName")
	if roomName == ""{
		w.WriteHeader(http.StatusBadRequest)
	}

	data := manager.GetRoomActivePresences(roomName)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func GetRoomPresencesByDelta(w http.ResponseWriter, req *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if req.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	roomName := req.URL.Query().Get("RoomName")
	delta := req.URL.Query().Get("delta")
	if roomName == "" || delta == ""{
		w.WriteHeader(http.StatusBadRequest)
	}

	data := manager.GetAllPresencesByRoomAndDelta(roomName, delta)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

