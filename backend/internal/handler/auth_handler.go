package handler

//go:generate mockgen -source=auth_handler.go -destination=mock/mock_auth_handler.go -package=mock

import (
	"backend/internal/model"
	"backend/internal/repository"
	"backend/internal/utils"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandlerInterface interface {
	Login(c echo.Context) error
	Register(c echo.Context) error
}

type AuthHandler struct {
	userRepo repository.UserRepository
}

func NewAuthHandler(userRepo repository.UserRepository) AuthHandlerInterface {
	return &AuthHandler{userRepo: userRepo}
}

// Login godoc
// @Summary ユーザーログイン
// @Description ユーザー名とパスワードでログインし、JWTトークンを返します
// @Tags auth
// @Accept json
// @Produce json
// @Param request body model.LoginRequest true "ログイン情報"
// @Success 200 {object} model.LoginResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /login [post]
func (h *AuthHandler) Login(c echo.Context) error {
	req := new(model.LoginRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	// ユーザーを検索
	user, err := h.userRepo.FindByUsername(req.Username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Database error"})
	}

	if user == nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid credentials"})
	}

	// パスワード検証
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid credentials"})
	}

	// JWTトークン生成
	token, err := utils.GenerateJWT(user.ID, user.Username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to generate token"})
	}

	response := model.LoginResponse{
		Token: token,
		User:  *user,
	}

	return c.JSON(http.StatusOK, response)
}

// Register godoc
// @Summary ユーザー登録
// @Description 新しいユーザーを登録します
// @Tags auth
// @Accept json
// @Produce json
// @Param request body model.RegisterRequest true "登録情報"
// @Success 201 {object} model.LoginResponse
// @Failure 400 {object} map[string]string
// @Failure 409 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /register [post]
func (h *AuthHandler) Register(c echo.Context) error {
	req := new(model.RegisterRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	// ユーザー名の重複チェック
	existingUser, err := h.userRepo.FindByUsername(req.Username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Database error"})
	}
	if existingUser != nil {
		return c.JSON(http.StatusConflict, map[string]string{"error": "Username already exists"})
	}

	// メールアドレスの重複チェック
	existingUser, err = h.userRepo.FindByEmail(req.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Database error"})
	}
	if existingUser != nil {
		return c.JSON(http.StatusConflict, map[string]string{"error": "Email already exists"})
	}

	// パスワードをハッシュ化
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to hash password"})
	}

	// ユーザー作成
	user, err := h.userRepo.Create(req.Username, req.Email, string(hashedPassword))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create user"})
	}

	// JWTトークン生成
	token, err := utils.GenerateJWT(user.ID, user.Username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to generate token"})
	}

	response := model.LoginResponse{
		Token: token,
		User:  *user,
	}

	return c.JSON(http.StatusCreated, response)
}
