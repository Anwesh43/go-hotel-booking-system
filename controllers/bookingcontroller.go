package controllers

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"demo.hotel/services"
)

type BookingController struct {
	roomService services.RoomService
	is          services.InputService
}

func (bc *BookingController) StartIO() {
	ch := make(chan []string)
	go bc.is.StartProcessing(ch)
	inputs := <-ch
	completeCh := make(chan bool)
	go func() {
		for _, inp := range inputs {
			strParts := strings.Split(inp, " ")
			if strParts[0] == "Insert" {
				beds, err1 := strconv.ParseInt(strParts[1], 10, 64)
				price, err2 := strconv.ParseInt(strParts[2], 10, 64)
				if err1 == nil && err2 == nil {
					bc.roomService.InsertRoom(int(beds), int(price))
					fmt.Println("Inserted a room")
				}
			} else if strParts[0] == "Book" {
				beds, err1 := strconv.ParseInt(strParts[1], 10, 32)
				if err1 == nil {
					room := bc.roomService.BookRoom(int(beds))
					fmt.Println("Room is haveing ", room.Beds, "And costs", room.Price)
				}

			} else if strParts[0] == "Leave" {
				bc.roomService.LeaveRoom(strParts[1])
				fmt.Println("Left a room")
			} else {
				fmt.Println("Please enter Insert, Leave, Book query")
			}
			time.Sleep(time.Second)
		}
		completeCh <- true
	}()
	<-completeCh
	fmt.Println("Completed transactions")
	bc.roomService.PrintDetails()
}

func NewBookingController() BookingController {
	return BookingController{
		roomService: services.NewRoomService(),
		is:          services.NewInputService(),
	}
}
