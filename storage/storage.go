package storage

import (
	"context"
	"order_service/genproto/order_service"
)

type StorageI interface {
	CloseDB()
	Order() OrderRepoI
	Product() ProductRepoI
}

type OrderRepoI interface {
	Create(ctx context.Context, order *order_service.CreateOrderRequest) (*order_service.OrderPKey, error)
	GetById(ctx context.Context, pKey *order_service.OrderPKey) (resp *order_service.Order, err error)
	GetAll(ctx context.Context, req *order_service.GetAllOrdersRequest) (resp *order_service.GetAllOrdersResponse, err error)
	Delete(ctx context.Context, pKey *order_service.OrderPKey) (err error)
	Update(ctx context.Context, req *order_service.UpdateOrderRequest) (err error)
}

type ProductRepoI interface {
	Create(ctx context.Context, product *order_service.CreateProductRequest) (pKey *order_service.ProductPKey,err  error)
	GetById(ctx context.Context, pKey *order_service.ProductPKey) (resp *order_service.Product, err error)
	GetAll(ctx context.Context, req *order_service.GetAllProductsRequest) (resp *order_service.GetAllProductsResponse, err error)
	Delete(ctx context.Context, pKey *order_service.ProductPKey) (err error)
	Update(ctx context.Context, req *order_service.UpdateProduct) (err error)
}