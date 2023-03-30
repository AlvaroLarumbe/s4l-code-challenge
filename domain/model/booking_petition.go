package model

import "time"

type BookingPetition struct {
	petitionId  string
	checkIn     time.Time
	nights      int
	sellingRate int
	margin      int
}
