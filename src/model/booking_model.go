package model

type SearchDto struct {
	From          string `json:"from"`
	To            string `json:"to"`
	DepartureDate string `json:"departure_date"`
}
type SearchResultDto struct {
	ID            int    `json:"id"`
	DepartureTime string `json:"departure_time"`
	ArrivalTime   string `json:"arrival_time"`
	Airline       string `json:"airline"`
	Price         int64  `json:"price"`
	DepartureDate string `json:"departure_date"`
}
type FlightDetailsDto struct {
	ID            int    `json:"id"`
	Airline       string `json:"Airline"`
	From          string `json:"Departure"`
	To            string `json:"Arrival"`
	DepartureDate string `json:"Departure_date"`
	ReturnDate    string `json:"Return_date"`
}
