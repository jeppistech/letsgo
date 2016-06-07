package weather

import "time"

type WeatherPoint struct {
	ID             int32     `json:"id"`
	Date           time.Time `json:"date"`
	Location       Location  `json:"location"`
	WeatherComment string    `json:"weatherComment"`
	Temperature    int32     `json:"temperature"`
	Humidity       int32     `json:"humidity"`
}
