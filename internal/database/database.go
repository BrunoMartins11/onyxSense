package database

import (
	"context"
	"database/sql"
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
	return err
}

func (db Database) GetRoomByName(roomName string) model.Room {
	var room model.Room
	_ = db.DB.QueryRowContext(context.TODO(),
		"SELECT * FROM rooms WHERE name = $1", roomName).Scan(&room)
	return room
}

func (db Database) SaveNewSensor(sensor model.Sensor, roomID int64) error {
	_, err := db.DB.ExecContext(context.TODO(),
		"INSERT INTO sensors (name, roomID) VALUES ($1, $2)",
		sensor.Name,
		roomID,
	)
	return err
}

func (db Database) SaveNewPresence(presence model.Presence , roomID int64) error {
	_, err := db.DB.ExecContext(context.TODO(),
		"INSERT INTO presences (MAC, lastDetected, roomID) VALUES ($1, $2, $3)",
		presence.MAC,
		presence.LastDetected,
		roomID,
	)
	return err
}

func (db Database) GetRoomPresences(roomID int64) []model.Presence {
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

