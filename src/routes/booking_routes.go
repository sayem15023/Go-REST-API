package routes

import (
	"sayem/ota/service"

	"github.com/gin-gonic/gin"
)

func SetupBookingRouter(bookingService *service.BookingService) *gin.Engine {
	router := gin.Default()
	bookingRoutes := router.Group("/booking")
	{
		bookingRoutes.POST("/search", bookingService.Search)
		bookingRoutes.GET("/bookingdetails/:id", bookingService.FlightDetails)

	}

	return router
}
