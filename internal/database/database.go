package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/BrunoMartins11/onyxSense/internal/model"
	"log"
)

type Database struct {
	DB *sql.DB
}

func NewDatabase(db *sql.DB) Database{
	return Database{db}
}

func (db Database) SaveNewRoom(room model.Room) error {
	_, err := db.DB.ExecContext(context.TODO(),
		"INSERT INTO rooms (name) VALUES ($1)",
		room.Name,
	)
	if err != nil {
		return errors.New("check if room name already exists")
	}
	return nil
}

func (db Database) GetRoomByName(roomName string) model.Room {
	var roomID int
	var name string
	err := db.DB.QueryRowContext(context.TODO(),
		"SELECT * FROM rooms WHERE name = $1", roomName).Scan(&roomID, &name)
	if err != nil {
		return model.Room{}
	}
	return model.Room{roomID, name}
}

func (db Database) SaveNewSensor(sensor model.Sensor, roomID int) error {
	_, err := db.DB.ExecContext(context.TODO(),
		"INSERT INTO sensors (roomid, name) VALUES ($1, $2)",
		roomID,
		sensor.Name,
	)
	if err != nil {
		return errors.New("check if sensor name already exists")
	}
	return err
}

func (db Database) SaveNewPresence(presence model.Presence , roomID int) error {
	fmt.Println(presence)
	fmt.Println(roomID)
	_, err := db.DB.ExecContext(context.TODO(),
		"INSERT INTO presences (mac, lastdetected, active, roomid) VALUES ($1, $2, $3, $4)",
		presence.MAC,
		presence.LastDetected,
		presence.Active,
		roomID,
	)

	return err
}

func (db Database) GetRoomPresences(roomID int) []model.Presence {
	//TODO fix like get room by name
	rows, err := db.DB.QueryContext(context.TODO(),
		"SELECT * FROM presences WHERE roomID = $1", roomID)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	var presences []model.Presence

	for rows.Next() {
		var pres model.Presence
		if err := rows.Scan(&pres); err != nil {
			log.Fatal(err)
		}
		presences = append(presences, pres)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return presences
}

func (db Database) GetRoomBySensorName(sensorName string) model.Room {
	var roomID int
	var name string
	err := db.DB.QueryRowContext(context.TODO(),
		"SELECT r.id, r.name FROM rooms r, sensors s WHERE s.name = $1 AND s.roomid = r.id", sensorName).Scan(&roomID, &name)
	if err != nil {
		log.Println(err)
		return model.Room{}
	}
	return model.Room{roomID, name}
}

