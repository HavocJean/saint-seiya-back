package user

import "golang.org/x/crypto/bcrypt"

type User struct {
	ID       uint
	Name     string
	Email    string
	Password string
}

func NewUser(id uint, name, email, passwordHash string) *User {
	return &User{
		ID:       id,
		Name:     name,
		Email:    email,
		Password: passwordHash,
	}
}

func HashedPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func (u *User) VerifyPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
