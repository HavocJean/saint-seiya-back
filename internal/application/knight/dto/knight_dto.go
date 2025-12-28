package dto

type GetKnightsResponse struct {
	ID       uint    `json:"id"`
	Name     string  `json:"name"`
	Rank     string  `json:"rank"`
	ImageURL *string `json:"image_url"`
}

type GetKnightByIdResponse struct {
	ID              uint     `json:"id"`
	Name            string   `json:"name"`
	Rank            string   `json:"rank"`
	Pv              int      `json:"pv"`
	AtkC            *int     `json:"atk_c"`
	DefC            int      `json:"def_c"`
	DefF            int      `json:"def_f"`
	AtkF            *int     `json:"atk_f"`
	Speed           int      `json:"speed"`
	StatusHit       *float64 `json:"status_hit"`
	StatusResist    *float64 `json:"status_resist"`
	CritDamageC     *float64 `json:"crit_damage_c"`
	ResistDamageC   *float64 `json:"resist_damage_c"`
	PerfurationDefC *float64 `json:"perfuration_def_c"`
	ReflectDamage   *float64 `json:"reflect_damage"`
	Heal            *float64 `json:"heal"`
	CritLevelF      *float64 `json:"crit_level_f"`
	CritEffectF     *float64 `json:"crit_effect_f"`
	CritResistF     *float64 `json:"crit_resist_f"`
	ResistDamageF   *float64 `json:"resist_damage_f"`
	PerfurationDefF *float64 `json:"perfuration_def_f"`
	LifeTheft       *float64 `json:"life_theft"`
	CritBasicF      *float64 `json:"crit_basic_f"`
	ImageURL        *string  `json:"image_url"`
}

type CreateKnightRequest struct {
	Name            string   `json:"name" binding:"required"`
	Rank            string   `json:"rank" binding:"required"`
	Pv              int      `json:"pv" binding:"required"`
	AtkC            *int     `json:"atk_c" binding:"required"`
	DefC            int      `json:"def_c" binding:"required"`
	DefF            int      `json:"def_f" binding:"required"`
	AtkF            *int     `json:"atk_f" binding:"required"`
	Speed           int      `json:"speed" binding:"required"`
	StatusHit       *float64 `json:"status_hit" binding:"required"`
	StatusResist    *float64 `json:"status_resist" binding:"required"`
	CritDamageC     *float64 `json:"crit_damage_c" binding:"required"`
	ResistDamageC   *float64 `json:"resist_damage_c" binding:"required"`
	PerfurationDefC *float64 `json:"perfuration_def_c"`
	ReflectDamage   *float64 `json:"reflect_damage" binding:"required"`
	Heal            *float64 `json:"heal" binding:"required"`
	CritLevelF      *float64 `json:"crit_level_f" binding:"required"`
	CritEffectF     *float64 `json:"crit_effect_f" binding:"required"`
	ResistCritF     *float64 `json:"resist_crit_f" binding:"required"`
	ResistDamageF   *float64 `json:"resist_damage_f" binding:"required"`
	PerfurationDefF *float64 `json:"perfuration_def_f" binding:"required"`
	LifeTheft       *float64 `json:"life_theft" binding:"required"`
	CritBasicF      *float64 `json:"crit_basic_f" binding:"required"`
	ImageURL        *string  `json:"image_url"`
}

type CreateKnightSkillRequest struct {
	Name        string                          `json:"name" binding:"required"`
	Type        string                          `json:"type" binding:"required"`
	ImageURL    *string                         `json:"image_url"`
	Description string                          `json:"description" binding:"required"`
	Levels      []CreateKnightSkillLevelRequest `json:"levels" binding:"required"`
}

type CreateKnightSkillLevelRequest struct {
	Level       int    `json:"level" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type CreateKnightSkillResponse struct {
	ID          uint                             `json:"id"`
	KnightID    uint                             `json:"knight_id"`
	Name        string                           `json:"name"`
	Type        string                           `json:"type"`
	ImageURL    *string                          `json:"image_url"`
	Description string                           `json:"description"`
	Levels      []CreateKnightSkillLevelResponse `json:"levels"`
}

type CreateKnightSkillLevelResponse struct {
	ID          uint   `json:"id"`
	SkillID     uint   `json:"skill_id"`
	Level       int    `json:"level"`
	Description string `json:"description"`
}
