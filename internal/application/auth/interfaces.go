package auth

type JWTService interface {
	GenerateToken(userID uint, email string) (string, error)
}
