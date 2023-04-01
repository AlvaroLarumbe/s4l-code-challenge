package model

import "time"

type BookingPetition struct {
	PetitionId  string
	CheckIn     time.Time
	Nights      int
	SellingRate int
	Margin      int
}
