package api

type BookingPetitionRequest struct {
	PetitionId  string `json:"request_id"`
	CheckIn     string `json:"check_in"`
	Nights      int    `json:"nights"`
	SellingRate int    `json:"selling_rate"`
	Margin      int    `json:"margin"`
}
