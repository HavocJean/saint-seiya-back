package team

type TeamDomain struct {
	ID          uint
	Name        string
	UserID      uint
	IsPublic    bool
	TeamKnights []TeamKnightDomain
}

type TeamKnightDomain struct {
	ID       uint
	TeamID   uint
	KnightID uint
}
