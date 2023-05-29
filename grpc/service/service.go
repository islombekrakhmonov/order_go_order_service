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

type productService struct {
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

func (b *productService) Create(ctx context.Context, req *order_service.CreateProductRequest) (resp *order_service.Product, err error) {
	b.log.Info("--CreateProduct--->", logger.Any("req", req))

	fmt.Println("The function is called successfully.")
	pKey, err := b.strg.Product().Create(ctx, req)

	if err != nil{
		b.log.Error("!!!CreateProduct--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return b.strg.Product().GetById(ctx, pKey)
}

func (b *productService) GetById(ctx context.Context, req *order_service.ProductPKey) (resp *order_service.Product, err error) {
	b.log.Info("---GetProduct--->", logger.Any("req", req))

	resp, err = b.strg.Product().GetById(ctx, req)
	if err != nil {
		b.log.Error("!!!GetProduct--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, nil
}

func (b *productService) GetAll(ctx context.Context, req *order_service.GetAllProductsRequest) (resp *order_service.GetAllProductsResponse, err error) {
	b.log.Info("---GetAllProducts--->", logger.Any("req", req))

	resp, err = b.strg.Product().GetAll(ctx, req)
	if err != nil {
		b.log.Error("!!!GetAllProducts--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return resp, nil
}

func (b *productService) Update(ctx context.Context, req *order_service.UpdateOrderRequest) (resp *order_service.Product, err error) {
	b.log.Info("---UpdateProduct--->", logger.Any("req", req))
	return nil, nil
}


func (b *orderService) Create(ctx context.Context, req *order_service.CreateOrderRequest) (resp *order_service.Order, err error) {
	b.log.Info("---CreateOrder--->", logger.Any("req", req))
 

		fmt.Println("The function is called successfully.")

	pKey, err := b.strg.Order().Create(ctx, req)

	if err != nil {
		fmt.Println("The error is coming from here.")
		b.log.Error("!!!CreateOrder--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}


	return b.strg.Order().GetById(ctx, pKey)
}

func (o *orderService) GetById(ctx context.Context, req *order_service.OrderPKey) (resp *order_service.Order, err error) {
	o.log.Info("---GetOrder--->", logger.Any("req", req))

	
	
	resp, err = o.strg.Order().GetById(ctx, req)
	if err != nil {
		o.log.Error("!!!GetOrder--->", logger.Error(err))
		fmt.Println("The error is coming from here.1")
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	user, err := o.services.UserService().GetById(ctx, &user_service.UserPKey{
		Id: resp.UserID,
	})
	if err != nil {
		fmt.Println("The error is coming from here.2")
		o.log.Error("!!!GetOrder--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}


	
	fmt.Println(user)
	resp.UserFirstName = user.FirstName
	resp.UserLastName = user.LastName
	resp.UserPhoneNumber = user.PhoneNumber
	
	return resp, nil
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