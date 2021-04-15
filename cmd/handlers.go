package main

import (
	"github.com/BrunoMartins11/onyxSense/internal/model"
	"net/http"
)

func CreateRoomHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	roomName := req.URL.Query().Get("RoomName")

	err := model.CreateRoom(roomName)
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func CreateSensorHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	roomName := req.URL.Query().Get("NameId")
	sensorName := req.URL.Query().Get("SensorName")
	if roomName == "" || sensorName == "" {
		w.WriteHeader(http.StatusBadRequest)
	}

	err := model.CreateSensor(sensorName, roomName)
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		_, _ = w.Write([]byte(`{"error": "` + err.Error() + `"}`))
		return
	}
	w.WriteHeader(http.StatusOK)
}
