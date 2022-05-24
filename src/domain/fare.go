package domain

type Fare struct {
	RideID uint64 `json:"id_ride"`
	Amount float64 `json:"fare_estimate"`
}