package grpc

import (
	"order_service/config"
	"order_service/grpc/client"
	"order_service/genproto/order_service"
	"order_service/pkg/logger"
	"order_service/storage"
	"order_service/grpc/service"


	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)


func SetUpServer(cfg config.Config, log logger.LoggerI, strg storage.StorageI, svcs client.ServiceManagerI) (grpcServer *grpc.Server) {
	grpcServer = grpc.NewServer()

	order_service.RegisterOrderServiceServer(grpcServer, service.NewOrderService(cfg, log, strg, svcs))

	reflection.Register(grpcServer)
	return
}