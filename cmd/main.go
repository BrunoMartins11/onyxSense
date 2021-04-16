package main

import (
	"database/sql"
	"github.com/BrunoMartins11/onyxSense/internal/database"
	"github.com/BrunoMartins11/onyxSense/internal/store"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)
var manager store.Manager

func main(){
	connStr := "user=onyx dbname=onyx_dev password=onyx host=localhost sslmode=disable"
	db, err := sql.Open("postgres",
		connStr)
	if err != nil {
		log.Fatal(err)
	}
	data := database.NewDatabase(db)
	manager = store.NewManager(data)

	http.HandleFunc("/createRoom", CreateRoomHandler)
	http.HandleFunc("/createSensor", CreateSensorHandler)

	log.Fatal(http.ListenAndServe(":3000", nil))
}
