package dto

type CreateTeamRequest struct {
	Name     string `json:"name" binding:"required"`
	UserID   uint   `json:"-"`
	IsPublic bool   `json:"is_public" binding:"required"`
}

type CreateTeamKnight struct {
	TeamID   uint `json:"team_id"`
	KnightID uint `json:"knight_id" binding:"required"`
}

type CreateTeamVote struct {
}

type CreateTeamFavorite struct {
}

type TeamWithKnightResponse struct {
	ID      uint                 `json:"id"`
	Name    string               `json:"name"`
	Knights []TeamKnightResponse `json:"knights"`
}

type TeamKnightResponse struct {
	KnightID uint   `json:"knight_id"`
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}
