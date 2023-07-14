package daos

import (
	"fmt"

	"demo.hotel/models"
)

type RoomDao struct {
	OccupiedRooms   map[string]models.Room
	UnoccupiedRooms map[string]models.Room
	TotalRooms      int
}

func (r *RoomDao) Book(beds int) models.Room {
	var roomid string = ""
	for room_id, room := range r.UnoccupiedRooms {
		if room.Beds == beds {
			roomid = room_id
			break
		}
	}
	room := r.UnoccupiedRooms[roomid]
	r.OccupiedRooms[roomid] = room
	delete(r.UnoccupiedRooms, roomid)
	room.Occupied = true
	return room
}

func (r *RoomDao) Leave(roomid string) {
	room := r.OccupiedRooms[roomid]
	r.UnoccupiedRooms[roomid] = room
	room.Occupied = false
	delete(r.OccupiedRooms, roomid)
}

func (r *RoomDao) GetBookedRoom(roomid string) models.Room {
	return r.OccupiedRooms[roomid]
}

func (r *RoomDao) IsRoomPresent() bool {
	return len(r.UnoccupiedRooms) > 0
}

func (r *RoomDao) InsertRoom(beds int, price int) {
	room := models.Room{
		Beds:     beds,
		Price:    price,
		Occupied: false,
		RoomId:   "",
	}
	r.TotalRooms += 1
	roomid := fmt.Sprintf("room_%d", r.TotalRooms)
	room.RoomId = roomid
	r.UnoccupiedRooms[roomid] = room
}

func (r *RoomDao) PrintOccupiedRooms() {
	for _, room := range r.OccupiedRooms {
		fmt.Println("OCCUPIED", room.RoomId, "COST IS", room.Price, "BEDS", room.Beds)
	}
}

func (r *RoomDao) PrintUnOccupiedRooms() {
	for _, room := range r.UnoccupiedRooms {
		fmt.Println("UNOCCUPIED", room.RoomId, "COST IS", room.Price, "BEDS", room.Beds)
	}
}
