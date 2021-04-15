package main

import (
	"github.com/BrunoMartins11/onyxSense/internal/comms"
	"github.com/BrunoMartins11/onyxSense/internal/database"
	"github.com/BrunoMartins11/onyxSense/internal/model"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func main(){
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	database.InitDatabase()
	model.MigrateRoom()
	model.MigratePresence()
	model.MigrateSensor()
	comms.CreateRabbitChannel()
	comms.SubscribeQueue()

	http.HandleFunc("/createRoom", CreateRoomHandler)
	http.HandleFunc("/createSensor", CreateSensorHandler)

	log.Fatal(http.ListenAndServe(":3000", nil))
}
