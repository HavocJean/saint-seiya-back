package controllers

import (
	"net/http"
	"saint-seiya-back/internal/application/auth"
	authdto "saint-seiya-back/internal/application/auth/dto"
	"saint-seiya-back/internal/responses"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AuthController struct {
	loginUsecase       *auth.LoginUseCase
	registerUsecase    *auth.RegisterUseCase
	getUserByIdUseCase *auth.GetUserByIdUseCase
}

func NewAuthController(
	loginU *auth.LoginUseCase,
	registerU *auth.RegisterUseCase,
	getUserByIdU *auth.GetUserByIdUseCase,
) *AuthController {
	return &AuthController{
		loginUsecase:       loginU,
		registerUsecase:    registerU,
		getUserByIdUseCase: getUserByIdU,
	}
}

func (a *AuthController) RegisterUser(c *gin.Context) {
	var req authdto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		if _, ok := err.(validator.ValidationErrors); ok {
			responses.ValidationError(c, http.StatusBadRequest, err)
			return
		}

		responses.Error(c, http.StatusBadRequest, "invalid JSON sent", err.Error())
		return
	}

	output, err := a.registerUsecase.Execute(c.Request.Context(), auth.RegisterInput{
		Name: req.Name, Nickname: req.Nickname, Email: req.Email, Password: req.Password,
	})

	if err != nil {
		responses.Error(c, http.StatusBadRequest, "Registration failed", err.Error())
		return
	}

	responses.Success(c, http.StatusCreated, "Registration successful", authdto.AuthResponse{
		Token: output.Token,
	})
}

func (a *AuthController) LoginUser(c *gin.Context) {
	var req authdto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		responses.Error(c, http.StatusBadRequest, "Invalid JSON sent", err.Error())
		return
	}

	output, err := a.loginUsecase.Execute(c.Request.Context(), auth.LoginInput{
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		responses.Error(c, http.StatusUnauthorized, "Invalid credentials", err.Error())
		return
	}

	responses.Success(c, http.StatusOK, "Login successful", authdto.AuthResponse{
		Token: output.Token,
	})
}

func (a *AuthController) GetProfileUser(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		responses.Error(c, http.StatusUnauthorized, "User ID not found in token", "Unauthorized")
		return
	}

	result, err := a.getUserByIdUseCase.Execute(userID.(uint))
	if err != nil {
		responses.Error(c, http.StatusNotFound, "User not found", err.Error())
		return
	}

	responses.Success(c, http.StatusOK, "Profile found", result)

}
