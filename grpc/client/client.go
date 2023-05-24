package client

import (
	"order_service/config"
	"order_service/genproto/user_service"

	"google.golang.org/grpc"
)

type ServiceManagerI interface {
	UserService() user_service.UserServiceClient
}

type grpcClients struct {
	userService user_service.UserServiceClient
}

func NewGrpcClients(cfg config.Config) (ServiceManagerI, error) {
	connUserService, err := grpc.Dial(cfg.UserServiceHost + cfg.UserServicePort, grpc.WithInsecure(), 
	)
    if err != nil {
		return nil, err 
	} 

	return &grpcClients{
		userService: user_service.NewUserServiceClient(connUserService),
	}, nil
}

func (g *grpcClients) UserService() user_service.UserServiceClient {
	return g.userService
}


