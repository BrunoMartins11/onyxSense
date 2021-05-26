package store

import (
	"encoding/json"
	"github.com/BrunoMartins11/onyxSense/internal/model"
	"log"
	"time"
)

type MSG struct {
	DeviceID   string
	MacAddress string
	Active     bool //in milliseconds
	Timestamp  time.Time
}

type Store interface {
	SaveNewRoom(room model.Room) error
	GetRoomByName(roomName string) model.Room
	SaveNewSensor(sensor model.Sensor, roomID int) error
	SaveNewPresence(presence model.Presence, roomID int) error
	GetRoomPresences(roomID int) []model.Presence
	GetRoomBySensorName(sensorName string) model.Room
	GetRooms() []model.Room
	GetActivePresencesByRoom(roomName string) []model.Presence
	GetAllPresencesByRoom(roomName string) []model.Presence
	UpdatePresenceState(MAC string, state bool)
}

type Queue interface {
	SubscribeToQueue(queueName string, channel chan []byte)
}

type Manager struct {
	Store Store
	Queue Queue
}

func NewManager(store Store, queue Queue) Manager {
	return Manager{store, queue}
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

func (manager Manager) GetCurrentRoomPresences(room model.Room) []model.Presence {
	return manager.Store.GetRoomPresences(room.ID)
}

func (manager Manager) GetRoomByName(room string) model.Room {
	return manager.Store.GetRoomByName(room)
}
func (manager Manager) RegisterNewPresence(presence model.Presence, roomID int) error {
	err := manager.Store.SaveNewPresence(presence, roomID)
	if err != nil {
		return err
	}
	return nil
}

func (manager Manager) InitializeSubscriber(queue string) {
	channel := make(chan []byte)
	go func() {
		for {
			msg := <-channel
			manager.SaveMSGReceived(msg)
		}
	}()
	manager.Queue.SubscribeToQueue(queue, channel)


}

func (manager Manager) SaveMSGReceived(payload []byte) {
	var msg MSG
	err := json.Unmarshal(payload, &msg)
	if err != nil {
		log.Fatal(err)
	}
	presence := model.Presence{
		MAC: msg.MacAddress,
		LastDetected: &msg.Timestamp,
		Active: msg.Active,
	}
	room := manager.Store.GetRoomBySensorName(msg.DeviceID)
	if msg.Active {
		err = manager.RegisterNewPresence(presence, room.ID)
	} else {
		manager.Store.UpdatePresenceState(msg.MacAddress, msg.Active)
	}
	if err != nil {
		log.Fatal(err)
	}
}

func (manager Manager) GetAllRooms() []model.Room{
	return manager.Store.GetRooms()
}

func (manager Manager) GetRoomActivePresences(roomName string) []model.Presence {
	return manager.Store.GetActivePresencesByRoom(roomName)
}

func (manager Manager) GetAllPresencesByRoomAndDelta(roomName string, delta string) []model.Presence{
	presences := manager.Store.GetAllPresencesByRoom(roomName)
	var ret []model.Presence
	for _, p := range presences {
		if delta == "M" && p.LastDetected.After(BeginningOfMonth(time.Now())){
			ret = append(ret, p)
		}
		if delta == "D" && p.LastDetected.After(BeginningOfDay(time.Now())){
			ret = append(ret, p)
		}
		if delta == "W" {
			year, week := p.LastDetected.ISOWeek()
			y, w := time.Now().ISOWeek()
			if year == y && w == week {
				ret = append(ret, p)
			}
		}
	}
	return ConvertTime(ret)
}

func BeginningOfMonth(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
}

func BeginningOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

func ConvertTime(list []model.Presence) []model.Presence{
	for _, p := range list {
		t := time.Date(p.LastDetected.Year(), p.LastDetected.Month(), p.LastDetected.Day(), 0, 0, 0, 0, p.LastDetected.Location())
		p.LastDetected = &t
	}
	return list
}
