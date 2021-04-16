package model

import (
	"time"
)

type Presence struct {
	Id int64
	MAC string
	LastDetected *time.Time
	RoomID int64
	Active bool
}
