package handlers

import (
	"backend/internal/application/auth"
	"backend/internal/application/usecase"
	"backend/internal/domain/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	UserUsecase usecase.UserUsecase
}

func NewAuthHandler(uc usecase.UserUsecase) *AuthHandler {
	return &AuthHandler{UserUsecase: uc}
}

type signupRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required"`
}

// @Summary      User signup
// @Description  Sign up and add new user info
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request body signupRequest true "Signup credentials"
// @Success      200 {object} signupResponse
// @Failure      401 {object} errorResponse
// @Router       /auth/signup [post]
func (h *AuthHandler) Signup(c *gin.Context) {
	var req signupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newUser := &user.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  req.Password,
	}

	createdUser, err := h.UserUsecase.Signup(newUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdUser.Password = ""
	c.JSON(http.StatusOK, createdUser)
}

type loginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// @Summary      User login
// @Description  Authenticate user and get JWT token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request body loginRequest true "Login credentials"
// @Success      200 {object} loginResponse
// @Failure      401 {object} errorResponse
// @Router       /auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var req loginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := h.UserUsecase.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid email or password"})
		return
	}

	token, err := auth.GenerateJWT(int(u.ID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token":      token,
		"id":         u.ID,
		"first_name": u.FirstName,
		"last_name":  u.LastName,
		"email":      u.Email,
	})
}

type signupResponse struct {
	ID        uint      `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
}

type loginResponse struct {
	Token     string `json:"token"`
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type errorResponse struct {
	Error string `json:"error"`
}