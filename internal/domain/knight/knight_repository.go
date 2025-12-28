package knight

type Repository interface {
	GetKnights(page, limit int, rank, name string) ([]KnightDomain, error)
	GetKnightById(id uint) (*KnightDomain, error)
	Create(knight *KnightDomain) (*KnightDomain, error)
	CreateSkill(skill *KnightSkillDomain) (*KnightSkillDomain, error)
}
