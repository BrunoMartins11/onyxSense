package model

import (
	"errors"
	"github.com/BrunoMartins11/onyxSense/internal/database"
	"gorm.io/gorm"
	"log"
)

type Room struct {
	gorm.Model
	Name string
	Presences []Presence
	Sensors []Sensor
}

func MigrateRoom() {
	err := database.DB.AutoMigrate(&Room{})
	if err != nil {
		log.Fatal(err, "Failed Presence migration")
	}
}

func CreateRoom(name string) error{
	room := Room{Name: name}
	if GetRoomByName(name).Name != "" {
		return errors.New("room already exists")
	}
	database.DB.Create(&room)
	return nil
}
func GetRoomByName(name string) Room {
	var room Room
	database.DB.First(&room, "Name = ? ", name)
	return room
}
