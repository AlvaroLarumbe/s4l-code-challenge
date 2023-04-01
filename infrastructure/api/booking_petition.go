package api

import (
	"s4l-code-challenge/domain/model"
	"s4l-code-challenge/helper"
	"time"
)

// BookingPetitionRequest has all fields as optionals in order to distinguish if the field is present or not
type BookingPetitionRequest struct {
	PetitionId  helper.Optional[string] `json:"request_id"`
	CheckIn     helper.Optional[string] `json:"check_in"`
	Nights      helper.Optional[int]    `json:"nights"`
	SellingRate helper.Optional[int]    `json:"selling_rate"`
	Margin      helper.Optional[int]    `json:"margin"`
}

func (r BookingPetitionRequest) toModel() (*model.BookingPetition, error) {
	// Calling MustGet method because our request has been validated in this point
	checkInDate, err := time.Parse(time.DateOnly, r.CheckIn.MustGet())

	if err != nil {
		return nil, err
	}

	return &model.BookingPetition{
		PetitionId:  r.PetitionId.MustGet(),
		CheckIn:     checkInDate,
		Nights:      r.Nights.MustGet(),
		SellingRate: r.SellingRate.MustGet(),
		Margin:      r.Margin.MustGet(),
	}, nil
}
