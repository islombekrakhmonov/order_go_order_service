package service

import (
	"context"
	"fmt"
	"order_service/config"
	"order_service/genproto/order_service"
	"order_service/genproto/user_service"
	"order_service/grpc/client"
	"order_service/pkg/logger"
	"order_service/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type orderService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	order_service.UnimplementedOrderServiceServer
}

func NewOrderService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, svcs client.ServiceManagerI) *orderService {
	return &orderService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: svcs,
	}
}

func (b *orderService) Create(ctx context.Context, req *order_service.CreateOrderRequest) (resp *order_service.Order, err error) {
	b.log.Info("---CreateOrder--->", logger.Any("req", req))
 

		b.services.UserService().Create(ctx, &user_service.CreateUserRequest{
			FirstName:"test",
			LastName: "test",
			PhoneNumber: "test",
		})

		fmt.Println("The function is called successfully.")

	pKey, err := b.strg.Order().Create(ctx, req,)

	if err != nil {
		b.log.Error("!!!CreateOrder--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}


	return b.strg.Order().Get(ctx, pKey)
}

func (o *orderService) GetById(ctx context.Context, req *order_service.OrderPKey) (resp *order_service.Order, err error) {
	o.log.Info("---GetOrder--->", logger.Any("req", req))

	resp, err = o.strg.Order().Get(ctx, req)
	if err != nil {
		o.log.Error("!!!GetOrder--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return o.strg.Order().Get(ctx, req)
}

func (o *orderService) GetAll(ctx context.Context, req *order_service.GetAllOrdersRequest) (resp *order_service.GetAllOrdersResponse, err error) {
	o.log.Info("---GetAllOrders--->", logger.Any("req", req))

	resp, err = o.strg.Order().GetAll(ctx, req)
	if err != nil {
		o.log.Error("!!!GetAllOrders--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, err
}

func (o *orderService) Delete(ctx context.Context, req *order_service.OrderPKey) (resp *emptypb.Empty,err error) {
	o.log.Info("---DeleteOrder--->", logger.Any("req", req))

	resp = &emptypb.Empty{}
	
	err = o.strg.Order().Delete(ctx, req)
	if err != nil {
		o.log.Error("!!!DeleteOrder--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, err
}