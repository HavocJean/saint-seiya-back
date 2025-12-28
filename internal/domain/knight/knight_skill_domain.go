package knight

type KnightSkillDomain struct {
	ID          uint
	KnightID    uint
	Name        string
	Type        string
	ImageURL    *string
	Description string
	Levels      []KnightSkillLevelDomain
}

type KnightSkillLevelDomain struct {
	ID          uint
	SkillID     uint
	Level       int
	Description string
}
