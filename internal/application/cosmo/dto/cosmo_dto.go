package dto

type GetCosmosResponse struct {
	ID       uint    `json:"id"`
	Name     string  `json:"name"`
	Rank     string  `json:"rank"`
	Color    string  `json:"color"`
	ImageURL *string `json:"image_url"`
}

type CosmoAttributeResponse struct {
	ID        uint    `json:"id"`
	CosmoID   uint    `json:"cosmo_id"`
	Name      string  `json:"name"`
	Value1    float64 `json:"value1"`
	Value10   float64 `json:"value_10"`
	IsPercent bool    `json:"is_percent"`
}

type GetCosmoByIdResponse struct {
	ID                uint                     `json:"id"`
	Name              string                   `json:"name"`
	Rank              string                   `json:"rank"`
	Color             string                   `json:"color"`
	SetBonusValue     float64                  `json:"set_bonus"`
	SetBonusName      string                   `json:"set_bonus_name"`
	SetBonusIsPercent bool                     `json:"set_bonus_is_percent"`
	BaseAttributes    []CosmoAttributeResponse `json:"base_attributes"`
	ImageURL          *string                  `json:"image_url"`
}
