package model

import (
	"time"
)

type Presence struct {
	ID int
	MAC string
	LastDetected *time.Time
	RoomID int64
	Active bool
}
