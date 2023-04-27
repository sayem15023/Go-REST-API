package main

import (
	"fmt"
	"sayem/ota/repository"
	"sayem/ota/service"
	"sayem/ota/routes"

)

func main() {
	fmt.Println("hello")

	bookingRepository := repository.NewBookingRepository()
	bookingService := service.NewBookingService(bookingRepository)

	r := routes.SetupBookingRouter(bookingService)
	r.Run(":8080")

}
