package repository

import (
	"encoding/json"
	"os"
	"sayem/ota/model"
)

type BookingRepository interface {
	SearchResult(string) ([]model.SearchResultDto, error)
	FlightDetails(int) (*model.FlightDetailsDto, error)
}
type bookingRepository struct{}

func NewBookingRepository() *bookingRepository {
	return &bookingRepository{}
}
func (b *bookingRepository) SearchResult(date string) ([]model.SearchResultDto, error) {
	file, err := os.Open("../flight_details.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var searchResultDtos []model.SearchResultDto
	var result []model.SearchResultDto
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&searchResultDtos)
	if err != nil {
		return nil, err
	}
	for _, search := range searchResultDtos {
		if search.DepartureDate == date {
			result = append(result, search)
		}
	}

	return result, nil
}
func (b bookingRepository) FlightDetails(id int) (*model.FlightDetailsDto, error) {
	file, err := os.Open("../flight_details.json")
	if err != nil {
		// log.Fatal("error opening file")
		return &model.FlightDetailsDto{}, err
	}

	defer file.Close()

	var flightDetails []model.FlightDetailsDto
	var flight model.FlightDetailsDto

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&flightDetails)
	if err != nil {
		return nil, err
	}

	for _, search := range flightDetails {
		if search.ID == id {
			flight = search
			return &flight, nil
		}
	}

	return &model.FlightDetailsDto{}, nil
}
