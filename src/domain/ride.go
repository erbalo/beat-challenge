package domain

type Ride struct {
	ID uint64 `json:"id_ride"`
	Locations []Location `json:"locations"`
}