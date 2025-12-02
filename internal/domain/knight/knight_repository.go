package knight

type Repository interface {
	GetKnights(page, limit int, rank, name string) ([]Knight, error)
	GetKnightById(id uint) (*Knight, error)
	Create(knight *Knight) (*Knight, error)
}
