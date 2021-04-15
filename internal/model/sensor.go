package model

import (
	"errors"
	"fmt"
	"github.com/BrunoMartins11/onyxSense/internal/database"
	"log"
)

type Sensor struct {
	Name string
	RoomID uint
}

func MigrateSensor() {
	err := database.DB.AutoMigrate(&Sensor{})
	if err != nil {
		log.Fatal(err, "Failed Presence migration")
	}
}

func CreateSensor(name string, roomName string) error{
	if GetSensorByName(name).Name != "" {
		return errors.New("sensor already exists")
	}
	var room Room
	if room = GetRoomByName(roomName); room.Name == ""  {
		return errors.New("room already exists")
	}
	sensor := Sensor{Name: name, RoomID: room.ID}
	database.DB.Create(&sensor)

	_ = database.DB.Model(&room).Association("Sensors").Append(&sensor)
	fmt.Println(room)
	return nil
}

func GetSensorByName(name string) Sensor{
	var sensor Sensor
	database.DB.First(&sensor, "Name = ? ", name)
	return sensor
}
