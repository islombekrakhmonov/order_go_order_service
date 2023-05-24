package postgres

import (
	"context"
	"log"
	"order_service/genproto/order_service"
	"os/exec"

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
		INSERT INTO orders (user_id, product_id)
		VALUES ($1, $2)
		`
	
	newUUID, err := exec.Command("uuidgen").Output()
	if err != nil {
        log.Fatal(err)
    }
	newUUID2, err := exec.Command("uuidgen").Output()
	if err != nil {
        log.Fatal(err)
    }

	_, err = o.db.Exec(ctx, query, newUUID, newUUID2)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (o *orderRepo) Get(ctx context.Context, pKey *order_service.OrderPKey) (resp *order_service.Order, err error) {
	var (
		query string
		order order_service.Order
	)

	query = `
		SELECT id, user_id, product_id
		FROM orders
		WHERE id = $1
	`
	err = o.db.QueryRow(ctx, query, pKey.Id).Scan(&order.Id, &order.UserID, &order.ProductID,
	)
	if err != nil {
		return nil, err
	}

	return &order, nil
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
		SELECT id, user_id, product_id
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
		err = rows.Scan(&order.Id, &order.UserID, &order.ProductID)
		if err != nil {
			return nil, err
		}
		resp.Orders = append(resp.Orders, &order)
	}
	return nil, nil
}

func (o *orderRepo) Delete(ctx context.Context, pKey *order_service.OrderPKey) (err error) {

	query := `

		DELETE FROM orders
		WHERE id = $1
	`
	_,err = o.db.Exec(ctx, query, pKey.Id)
	return nil
}

func (o *orderRepo) Update(ctx context.Context, pKey *order_service.OrderPKey) (err error) {
	return nil
}
