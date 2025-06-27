package router

import (
	"qfzack/go-web-starter/internal/server/handler"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(rg *gin.RouterGroup, userHandler *handler.UserHandler) {
	userGroup := rg.Group("/users")
	{
		// open routes
		userGroup.POST("/:id", userHandler.GetUser)

		// // routes reuquire auth
		// authGroup := userGroup.Group("")
		// authGroup.Use(middleware.AuthRequired())
		// {
		// 	authGroup.GET("/profile", userHandler.GetUserProfile)
		// 	authGroup.PUT("/profile", userHandler.UpdateUserProfile)
		// 	authGroup.DELETE("/account", userHandler.DeleteUser)
		// 	authGroup.GET("/orders", userHandler.GetUserOrders)
		// }

		// // routes for admin
		// adminGroup := userGroup.Group("/admin")
		// adminGroup.Use(middleware.AuthRequired(), middleware.AdminRequired())
		// {
		// 	adminGroup.GET("", userHandler.ListUsers)
		// 	adminGroup.GET("/:id", userHandler.GetUserByID)
		// 	adminGroup.PUT("/:id/status", userHandler.UpdateUserStatus)
		// }
	}
}
