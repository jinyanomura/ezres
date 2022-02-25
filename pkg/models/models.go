package models

import "time"

type TemplateData struct {
	StringMap map[string]string
	IntMap map[string]int
	FloatMap map[string]float32
	Data map[string]interface{}
	CSRFToken string
	Flash string
	Warning string
	Error string
}

type Table struct {
	ID int
	Capacity int
	Restrictions []Restriction
}

type Restriction struct {
	ID int
	StartTime time.Time
	EndTime time.Time
	TableID int
	RestrictionID int
	ReservationID int
	CreatedAt time.Time
	UpdatedAt time.Time
}