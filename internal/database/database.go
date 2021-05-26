package database

import (
	"context"
	"database/sql"
	"errors"
	"github.com/BrunoMartins11/onyxSense/internal/model"
	"log"
	"time"
)

type Database struct {
	DB *sql.DB
}

func NewDatabase(db *sql.DB) Database {
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

func (db Database) SaveNewPresence(presence model.Presence, roomID int) error {
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
	rows, err := db.DB.QueryContext(context.TODO(),
		"SELECT * FROM presences WHERE roomID = $1", roomID)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	var presences []model.Presence

	for rows.Next() {
		presences = parsePresences(rows, presences)
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

func (db Database) GetRooms() []model.Room {
	rows, err := db.DB.QueryContext(context.TODO(),
		"SELECT * FROM rooms")

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	var rooms []model.Room

	for rows.Next() {
		var roomID int
		var roomName string
		if err := rows.Scan(&roomID, &roomName); err != nil {
			log.Fatal(err)
		}
		rooms = append(rooms, model.Room{roomID, roomName})
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return rooms
}
func (db Database) GetActivePresencesByRoom(roomName string) []model.Presence {
	rows, err := db.DB.QueryContext(context.TODO(),
		"SELECT p.id, p.MAC, p.lastdetected, p.roomid, p.active FROM rooms r, presences p WHERE p.roomid = r.id AND $1 = r.name AND p.active IS true", roomName)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	var presences []model.Presence

	for rows.Next() {
		presences = parsePresences(rows, presences)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return presences
}

func (db Database) GetAllPresencesByRoom(roomName string) []model.Presence {
	rows, err := db.DB.QueryContext(context.TODO(),
		"SELECT p.id, p.MAC, p.lastdetected, p.roomid, p.active FROM rooms r, presences p WHERE p.roomid = r.id AND $1 = r.name", roomName)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	var presences []model.Presence

	for rows.Next() {
		presences = parsePresences(rows, presences)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return presences
}

func (db Database) UpdatePresenceState(MAC string, state bool) {
	err, _ := db.DB.ExecContext(context.TODO(),
		"UPDATE presences SET active = $2 WHERE mac = $1", &MAC, &state)
	if err != nil {
		log.Println(err)
	}
}

func parsePresences(rows *sql.Rows, presences []model.Presence) []model.Presence {
	var ID int
	var MAC string
	var LastDetected *time.Time
	var RoomID int
	var Active bool
	if err := rows.Scan(&ID, &MAC, &LastDetected, &RoomID, &Active); err != nil {
		log.Fatal(err)
	}
	presences = append(presences, model.Presence{ID, MAC, LastDetected, RoomID, Active})
	return presences
}