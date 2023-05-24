package storage

import (
	"context"
	"order_service/genproto/order_service"
)

type StorageI interface {
	CloseDB()
	Order() OrderRepoI
}

type OrderRepoI interface {
	Create(ctx context.Context, order *order_service.CreateOrderRequest) (*order_service.OrderPKey, error)
	Get(ctx context.Context, pKey *order_service.OrderPKey) (resp *order_service.Order, err error)
	GetAll(ctx context.Context, req *order_service.GetAllOrdersRequest) (resp *order_service.GetAllOrdersResponse, err error)
	Delete(ctx context.Context, pKey *order_service.OrderPKey) (err error)
	Update(ctx context.Context, pKey *order_service.OrderPKey) (err error)
}