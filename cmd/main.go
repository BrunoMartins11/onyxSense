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
	db, err := sql.Open("postgres",
		"onyx:onyx@tcp(127.0.0.1:5432)/onyx_dev")
	if err != nil {
		log.Fatal(err)
	}
	data := database.NewDatabase(db)
	manager = store.NewManager(data)

	http.HandleFunc("/createRoom", CreateRoomHandler)
	http.HandleFunc("/createSensor", CreateSensorHandler)

	log.Fatal(http.ListenAndServe(":3000", nil))
}
