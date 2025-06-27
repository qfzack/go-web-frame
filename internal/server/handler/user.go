package handler

import (
	"net/http"
	"qfzack/grpc-demo/internal/server/model"
	"qfzack/grpc-demo/internal/server/service"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	BaseHandler

	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) GetUser(c *gin.Context) {
	id := c.Param("id")
	user, err := h.userService.GetUser(id)
	if err != nil {
		h.handleServiceError(c, err)
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Code: http.StatusOK,
		Data: user,
	})
}
