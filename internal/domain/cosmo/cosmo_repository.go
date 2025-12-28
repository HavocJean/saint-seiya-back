package cosmo

type Repository interface {
	GetCosmos(color string) ([]CosmoDomain, error)
	GetCosmoByID(id uint) (*CosmoDomain, error)
	Create(cosmo *CosmoDomain) (*CosmoDomain, error)
}
