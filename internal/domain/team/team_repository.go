package team

type Repository interface {
	Create(team *TeamDomain) (*TeamDomain, error)
	GetByID(id uint) (*TeamDomain, error)
	GetByUserID(userID uint) ([]TeamDomain, error)
	CountByUserID(userID uint) (int64, error)
	Delete(id uint, userID uint) error
	AddKnight(teamKnight *TeamKnightDomain) (*TeamKnightDomain, error)
	CountKnightsByTeamID(teamID uint) (int64, error)
	KnightExistsInTeam(teamID uint, knightID uint) (bool, error)
	DeleteTeamKnight(teamID uint, knightID uint) error
}
