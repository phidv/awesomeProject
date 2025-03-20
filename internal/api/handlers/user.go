package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"oms/internal/api/dto/request"
	"oms/internal/api/dto/response"
	"oms/internal/services"
	"oms/pkg/utils"
)

type UserHandler struct {
	UserService services.UserService
}

func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{UserService: userService}
}

func (u *UserHandler) Get(ctx *gin.Context) {
	var req request.GetUserRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		utils.WriteErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}

	user, err := u.UserService.Get(ctx, req)
	if err != nil {
		utils.WriteErrorResponse(ctx, err, http.StatusBadRequest)
		return
	}

	ctx.JSON(http.StatusOK, response.CommonResponse{
		Data: user,
	})
}
