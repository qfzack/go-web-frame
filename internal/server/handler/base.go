package handler

import (
	"errors"
	"net/http"
	"qfzack/go-web-starter/internal/server/model"

	"github.com/gin-gonic/gin"
)

type BaseHandler struct{}

func (h *BaseHandler) handleError(c *gin.Context, statusCode int, message string, err error) {
	c.JSON(statusCode, model.ErrorResponse{
		Code:    statusCode,
		Messgae: message,
		Details: err.Error(),
	})
}

func (h *BaseHandler) handleServiceError(c *gin.Context, err error) {
	var businessErr *model.BusinessError
	if errors.As(err, &businessErr) {
		statusCode := h.getStatusCode(businessErr)
		c.JSON(statusCode, model.ErrorResponse{
			Code:    statusCode,
			Messgae: businessErr.Message,
		})
		return
	}

	c.JSON(http.StatusInternalServerError, model.ErrorResponse{
		Code:    http.StatusInternalServerError,
		Messgae: "Internal server error",
	})
}

func (h *BaseHandler) getStatusCode(err *model.BusinessError) int {
	switch err.Code {
	case "USER_NOT_FOUND":
		return http.StatusNotFound
	case "USER_EXISTS":
		return http.StatusConflict
	default:
		return http.StatusBadRequest
	}
}
