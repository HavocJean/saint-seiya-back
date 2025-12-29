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

type TeamWithKnightsDomain struct {
	ID      uint
	Name    string
	Knights []TeamKnightInfoDomain
}

type TeamKnightInfoDomain struct {
	KnightID uint
	Name     string
	ImageURL *string
}
