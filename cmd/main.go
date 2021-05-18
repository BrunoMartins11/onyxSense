package main

import (
	"database/sql"
	"github.com/BrunoMartins11/onyxSense/internal/comms"
	"github.com/BrunoMartins11/onyxSense/internal/database"
	"github.com/BrunoMartins11/onyxSense/internal/store"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)
var manager store.Manager

func main(){
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	connStr := "user=onyx dbname=onyx_dev password=onyx host=localhost sslmode=disable"
	db, err := sql.Open(os.Getenv("DB_DRIVER"),
		connStr)
	if err != nil {
		log.Fatal(err)
	}
	data := database.NewDatabase(db)
	queue := comms.CreateMQClient()
	manager = store.NewManager(data, queue)

	go manager.InitializeSubscriber("QueueService1")

	http.HandleFunc("/createRoom", CreateRoomHandler)
	http.HandleFunc("/createSensor", CreateSensorHandler)
	http.HandleFunc("/getRooms", GetRooms)
	http.HandleFunc("/getRoomActivePresences", GetRoomActivePresences)
	http.HandleFunc("/getRoomPresencesByDelta", GetRoomPresencesByDelta)


	log.Fatal(http.ListenAndServe(os.Getenv("PORT"), nil))
}
