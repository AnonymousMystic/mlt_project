package models

import "time"

type CalendarEvent struct {
	Eventid      string    `json:"eventid" gorm:"primaryKey"`
	Uuid         string    `json:"uuid" gorm:"primaryKey"`
	Title        string    `json:"title" gorm:"not null"`
	EventLength  *int      `json:"event_length"`
	EventDate    time.Time `json:"event_date" gorm:"not null"`
	IsApt        bool      `json:"isapt" gorm:"not null"`
	CustomFields *string   `json:"custom_fields"`
}
