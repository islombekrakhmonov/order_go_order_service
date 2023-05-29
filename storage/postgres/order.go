package postgres

import (
	"context"
	"order_service/genproto/order_service"
	"order_service/pkg/helper"
	"os/exec"
	"strings"

	"github.com/jackc/pgx/v4/pgxpool"
)

type orderRepo struct {
	db *pgxpool.Pool
}

func NewOrderRepo(db *pgxpool.Pool) *orderRepo {
	return &orderRepo{
		db: db,
	}
}

func (o *orderRepo) Create(ctx context.Context, order *order_service.CreateOrderRequest) (*order_service.OrderPKey, error){
	query := `
		INSERT INTO orders (id, user_id, product_id, totalsum)
		VALUES ($1, $2, $3, $4)
		`
		// Generate a UUID from the byte array.
		newUUID, err := exec.Command("uuidgen").Output()
		if err != nil {
			return nil, err
		}
		uuid := strings.TrimSpace(string(newUUID))



	_, err = o.db.Exec(ctx, query, 
		uuid, 
		order.UserID, 
		order.ProductID, 
		order.TotalSum)
	if err != nil {
		return nil, err
	}
	return &order_service.OrderPKey{Id:uuid}, nil
}

func (o *orderRepo) GetById(ctx context.Context, pKey *order_service.OrderPKey) (resp *order_service.Order, err error) {
	var (
		query string
	)
	resp = &order_service.Order{}


	query = `
		SELECT id, user_id, product_id, totalsum
		FROM orders
		WHERE id = $1
	`

	err = o.db.QueryRow(ctx, query, pKey.Id).Scan(
		&resp.Id, 
		&resp.UserID, 
		&resp.ProductID,
		&resp.TotalSum,
	)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (o *orderRepo) GetAll(ctx context.Context, req *order_service.GetAllOrdersRequest) (resp *order_service.GetAllOrdersResponse, err error) {
	resp = &order_service.GetAllOrdersResponse{}
	var (
		query  string
		filter = " WHERE TRUE "
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT id, user_id, product_id, totalsum
		FROM orders
	`

	query += filter + offset + limit

	rows, err := o.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var order order_service.Order
		err = rows.Scan(&order.Id, &order.UserID, &order.ProductID, &order.TotalSum)
		if err != nil {
			return nil, err
		}
		resp.Orders = append(resp.Orders, &order)
	}
	resp.Count = int64(len(resp.Orders))
	return resp, nil
}

func (o *orderRepo) Delete(ctx context.Context, pKey *order_service.OrderPKey) (err error) {

	query := `

		DELETE FROM orders
		WHERE id = $1
	`
	_,err = o.db.Exec(ctx, query, pKey.Id)
	if err != nil {
		return err
	}

	return nil
}

func (o *orderRepo) Update(ctx context.Context, req *order_service.UpdateOrderRequest) (err error) {
	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
		staffs
		SET
			id = :id, 
			user_id = :user_id,
			product_id = :product_id,
			total_sum = :total_sum,
		WHERE id = :id
	`
	params = map[string]interface{}{
		"id":    req.Id,
		"name":  req.UserID,
		"price": req.ProductID,
		"total_sum": req.TotalSum,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	_, err = o.db.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	return 
}
