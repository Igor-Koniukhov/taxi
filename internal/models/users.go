package models

import "time"

type User struct {
	ID            int       `json:"id"`
	Name          string    `json:"name"`
	Phone         string    `json:"phone"`
	StatusFree    bool      `json:"status_free"`
	StartPosition GeoPoint  `json:"start_position"`
	EndPosition   GeoPoint  `json:"end_position"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
