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

type CreateCosmoRequest struct {
	Name              string                        `json:"name" binding:"required"`
	Rank              string                        `json:"rank" binding:"required"`
	Color             string                        `json:"color" binding:"required"`
	SetBonusValue     float64                       `json:"set_bonus" binding:"required"`
	SetBonusName      string                        `json:"set_bonus_name" binding:"required"`
	SetBonusIsPercent bool                          `json:"set_bonus_is_percent" binding:"required"`
	ImageURL          *string                       `json:"image_url"`
	BaseAttributes    []CreateCosmoAttributeRequest `json:"base_attributes" binding:"required"`
}

type CreateCosmoAttributeRequest struct {
	Name      string  `json:"name" binding:"required"`
	Value1    float64 `json:"value1" binding:"required"`
	Value10   float64 `json:"value_10" binding:"required"`
	IsPercent bool    `json:"is_percent" binding:"required"`
}

type CreateCosmoResponse struct {
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
