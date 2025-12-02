package dto

type GetKnightsResponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Rank     string `json:"rank"`
	ImageURL string `json:"image_url"`
}
