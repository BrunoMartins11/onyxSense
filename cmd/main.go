package main

import (
	"github.com/BrunoMartins11/onyxSense/internal/database"
	"github.com/BrunoMartins11/onyxSense/internal/model"
	"log"
	"net/http"
)

func main(){
	database.InitDatabase()
	model.MigrateRoom()
	model.MigratePresence()
	model.MigrateSensor()


	http.HandleFunc("/createRoom", CreateRoomHandler)
	http.HandleFunc("/createSensor", CreateSensorHandler)

	log.Fatal(http.ListenAndServe(":3000", nil))
}
