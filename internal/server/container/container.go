package container

import (
	"qfzack/go-web-starter/internal/server/handler"
	"qfzack/go-web-starter/internal/server/repository"
	"qfzack/go-web-starter/internal/server/service"
)

type Container struct {
	// repositories
	UserRepo *repository.UserRepository

	//services
	UserService *service.UserService

	//handlers
	UserHandler *handler.UserHandler
	RPCServer   *handler.RPCServer
}

func NewContainer() *Container {
	userRepo := repository.NewUserRepository()
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)
	rpcServer := handler.NewRPCServer()

	return &Container{
		UserRepo:    &userRepo,
		UserService: &userService,
		UserHandler: userHandler,
		RPCServer:   rpcServer,
	}
}
