package container

import (
	"qfzack/grpc-demo/internal/server/handler"
	"qfzack/grpc-demo/internal/server/repository"
	"qfzack/grpc-demo/internal/server/service"
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
