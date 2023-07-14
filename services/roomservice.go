package services

import (
	"demo.hotel/daos"
	"demo.hotel/models"
)

type RoomService struct {
	roomDaoObj daos.RoomDao
}

func (r *RoomService) BookRoom(beds int) models.Room {
	var room models.Room
	if r.roomDaoObj.IsRoomPresent() {
		room = r.roomDaoObj.Book(beds)
	}
	return room
}

func (r *RoomService) LeaveRoom(roomid string) {
	r.roomDaoObj.Leave(roomid)
}

func (r *RoomService) InsertRoom(beds int, price int) {
	r.roomDaoObj.InsertRoom(beds, price)
}

func (r *RoomService) PrintDetails() {
	r.roomDaoObj.PrintOccupiedRooms()
	r.roomDaoObj.PrintUnOccupiedRooms()
}

func NewRoomService() RoomService {
	return RoomService{
		roomDaoObj: daos.RoomDao{
			OccupiedRooms:   make(map[string]models.Room),
			UnoccupiedRooms: make(map[string]models.Room),
			TotalRooms:      int(0),
		},
	}
}
