package domain

import "time"

type Location struct {
	Lat       float64 `json:"lat"`
	Lon       float64 `json:"lon"`
	Timestamp int64   `json:"timestamp"`
}

func (location *Location) IsDatyTime() bool {
	time := time.Unix(int64(location.Timestamp), 0).UTC()
	hour := time.Hour()

	return hour >= 5 && hour < 24
}
