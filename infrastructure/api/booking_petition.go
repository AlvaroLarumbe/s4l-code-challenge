package api

import (
	"errors"
	"net/http"
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

func validateRequest(
	s Server,
	response http.ResponseWriter,
	requestList []BookingPetitionRequest,
) error {
	if len(requestList) == 0 {
		s.writeBadRequestError(response)

		return errors.New("no requests given")
	}

	for _, request := range requestList {
		if !request.PetitionId.IsDefined() {
			// TODO: try to improve this raw JSON name with reflection or some validation library
			s.writeMissingFieldError(response, "request_id")

			return errors.New("missing field")
		}

		if !request.CheckIn.IsDefined() {
			s.writeMissingFieldError(response, "check_in")

			return errors.New("missing field")
		}

		if !request.Nights.IsDefined() {
			s.writeMissingFieldError(response, "nights")

			return errors.New("missing field")
		}

		if !request.SellingRate.IsDefined() {
			s.writeMissingFieldError(response, "selling_rate")

			return errors.New("missing field")
		}

		if !request.Margin.IsDefined() {
			s.writeMissingFieldError(response, "margin")

			return errors.New("missing field")
		}
	}

	return nil
}
