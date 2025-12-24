package knight

type KnightDomain struct {
	ID              uint
	Name            string
	Rank            string
	Pv              int
	AtkC            *int
	DefC            int
	DefF            int
	AtkF            *int
	Speed           int
	StatusHit       *float64
	CritLevelF      *float64
	StatusResist    *float64
	CritDamageC     *float64
	ResistDamageC   *float64
	PerfurationDefC *float64
	ReflectDamage   *float64
	Heal            *float64
	CritEffectF     *float64
	CritResistF     *float64
	ResistDamageF   *float64
	PerfurationDefF *float64
	LifeTheft       *float64
	CritBasicF      *float64
	ImageURL        *string
}
