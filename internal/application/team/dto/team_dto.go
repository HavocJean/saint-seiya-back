package dto

type CreateTeamRequest struct {
	Name     string `json:"name" binding:"required"`
	IsPublic bool   `json:"is_public" binding:"required"`
}

type CreateTeamResponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	UserID   uint   `json:"user_id"`
	IsPublic bool   `json:"is_public"`
}

type AddKnightToTeamRequest struct {
	KnightID uint `json:"knight_id" binding:"required"`
}

type AddKnightToTeamResponse struct {
	ID       uint `json:"id"`
	TeamID   uint `json:"team_id"`
	KnightID uint `json:"knight_id"`
}
