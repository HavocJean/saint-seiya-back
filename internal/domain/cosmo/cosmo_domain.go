package cosmo

type CosmoColor string

const (
	CosmoRed       CosmoColor = "red"
	CosmoBlue      CosmoColor = "blue"
	CosmoYellow    CosmoColor = "yellow"
	CosmoLegendary CosmoColor = "legendary"
)

type CosmoDomain struct {
	ID                uint
	Name              string
	Rank              string
	Color             CosmoColor
	SetBonusValue     float64
	SetBonusName      string
	SetBonusIsPercent bool
	BaseAttributes    []CosmoAttributeDomain
	ImageURL          *string
}

type CosmoAttributeDomain struct {
	ID        uint
	CosmoID   uint
	Name      string
	Value1    float64
	Value10   float64
	IsPercent bool
}
