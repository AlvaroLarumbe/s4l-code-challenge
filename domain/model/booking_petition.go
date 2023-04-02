package model

import "time"

type BookingPetition struct {
	PetitionId  string
	CheckIn     time.Time
	Nights      int
	SellingRate int
	Margin      int
}

func (p BookingPetition) GetCheckOut() time.Time {
	return p.CheckIn.AddDate(0, 0, p.Nights-1)
}

func (p BookingPetition) GetProfit() float64 {
	return float64(p.SellingRate) * float64(p.Margin) / 100
}

// ByCheckInDate implements sort.Interface in order to allow sorting by CheckIn field
type ByCheckInDate []BookingPetition

func (b ByCheckInDate) Len() int {
	return len(b)
}

func (b ByCheckInDate) Less(i, j int) bool {
	return b[i].CheckIn.Before(b[j].CheckIn)
}

func (b ByCheckInDate) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}
