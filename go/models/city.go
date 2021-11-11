package models

type City struct {
	Name       string
	Lat        float64
	Lng        float64
	Population int
	Country    *Country
}
