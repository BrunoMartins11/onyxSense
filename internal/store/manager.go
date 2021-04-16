package store

import (
	"github.com/BrunoMartins11/onyxSense/internal/model"
)

type Store interface {
	SaveNewRoom(room model.Room) error
	GetRoomByName(roomName string) model.Room
	SaveNewSensor(sensor model.Sensor, roomID int64) error
	SaveNewPresence(presence model.Presence, roomID int64) error
	GetRoomPresences(roomID int64) []model.Presence
}

type Manager struct {
	Store Store
}

func NewManager(store Store) Manager{
	return Manager{store}
}

func (manager Manager) RegisterNewRoom(room model.Room) error {
	err := manager.Store.SaveNewRoom(room)
	if err != nil {
		return err
	}
	return nil
}
func (manager Manager) RegisterNewSensor(sensor model.Sensor, roomName string) error {
	roomID := manager.Store.GetRoomByName(roomName).ID
	err := manager.Store.SaveNewSensor(sensor, roomID)
	if err != nil {
		return err
	}
	return nil
}
func (manager Manager) GetCurrentRoomPresences(room model.Room) []model.Presence{
	return manager.Store.GetRoomPresences(room.ID)
}

func (manager Manager) GetRoomByName(room string) model.Room{
	return manager.Store.GetRoomByName(room)
}
func (manager Manager) RegisterNewPresence(presence model.Presence, roomName string) error {
	roomID := manager.Store.GetRoomByName(roomName).ID
	err := manager.Store.SaveNewPresence(presence, roomID)
	if err != nil {
		return err
	}
	return nil
}

