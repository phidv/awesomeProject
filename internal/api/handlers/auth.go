package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"oms/internal/api/dto/request"
	"oms/internal/api/dto/response"
	"oms/internal/services"
	"oms/pkg/utils"
)

type AuthHandler struct {
	AuthService services.AuthService
}

func NewAuthHandler(authService services.AuthService) *AuthHandler {
	return &AuthHandler{AuthService: authService}
}

func (h *AuthHandler) Register(ctx *gin.Context) {
	var req request.RegisterRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.WriteErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}

	user, err := h.AuthService.Register(ctx, req)
	if err != nil {
		utils.WriteErrorResponse(ctx, err, http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusCreated, response.CommonResponse{Data: user})
}

func (h *AuthHandler) Login(ctx *gin.Context) {
	var req request.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		utils.WriteErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}

	token, err := h.AuthService.Authenticate(ctx, req)
	if err != nil {
		utils.WriteErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}

	ctx.JSON(http.StatusOK, response.CommonResponse{
		Data: response.LoginResponse{Email: req.Email, Token: token},
	})
}
