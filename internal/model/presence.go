package model

import (
	"github.com/BrunoMartins11/onyxSense/internal/database"
	"gorm.io/gorm"
	"log"
	"time"
)

type Presence struct {
	gorm.Model
	Id int `gorm:"uniqueIndex"`
	MAC string
	LastDetected *time.Time
	RoomID uint
}

func MigratePresence() {
	err := database.DB.AutoMigrate(&Presence{})
	if err != nil {
		log.Fatal(err, "Failed Presence migration")
	}
}

func CreatePresence(MAC string, lastDet *time.Time, room Room){
	pres := Presence{MAC: MAC, LastDetected: lastDet}
	database.DB.Create(&pres)
	err := database.DB.Model(&room).Association("Presences").Append(pres)
	if err != nil {
		log.Fatal(err, "Error associating Presence to Room")
	}
}