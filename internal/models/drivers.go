package models

import "time"

type Driver struct {
	ID            int       `json:"id"`
	Name          string    `json:"name"`
	Car           string    `json:"car"`
	StatusFree    bool      `json:"status_free"`
	StartPosition GeoPoint  `json:"start_position"`
	EndPosition   GeoPoint  `json:"end_position"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	DeletedAt     time.Time `json:"deleted_at"`
}

type GeoPoint struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Route struct {
	ID     int `json:"id"`
	Points []GeoPoint
}
