package service

import (
	"encoding/json"
	"net/http"
	"sayem/ota/model"
	"sayem/ota/repository"

	"strconv"

	"github.com/gin-gonic/gin"
)

type BookingService struct {
	bookingRepository repository.BookingRepository
}

func NewBookingService(repo repository.BookingRepository) *BookingService {
	return &BookingService{repo}
}
func (b *BookingService) Search(c *gin.Context) {
	var searchDto model.SearchDto
	decoder := json.NewDecoder(c.Request.Body)
	err := decoder.Decode(&searchDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	departureDate := searchDto.DepartureDate
	searchResultDtos, err :=
		b.bookingRepository.SearchResult(departureDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(200, searchResultDtos)
}

func (b BookingService) FlightDetails(c *gin.Context) {
	str := c.Param("id")
	id, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	flight, err := b.bookingRepository.FlightDetails(int(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(200, flight)
}
