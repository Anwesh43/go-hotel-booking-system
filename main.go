package main

import "demo.hotel/controllers"

func main() {
	bc := controllers.NewBookingController()
	bc.StartIO()
}
