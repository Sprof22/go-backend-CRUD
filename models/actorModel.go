package models

import "gorm.io/gorm"

type Actor struct {
	gorm.Model
	ActorName       string
	ActorRating     int
	ImagePath       string
	AlternativeName string
	ActorID         int
}
