package postgres

import (
	"context"
	"order_service/genproto/order_service"
	"os/exec"
	"strings"

	"github.com/jackc/pgx/v4/pgxpool"
)


type productRepo struct {
	db *pgxpool.Pool
}

func NewProductRepo(db *pgxpool.Pool) *productRepo {
	return &productRepo{
		db: db,
	}
}

func (o *productRepo) Create(ctx context.Context, product *order_service.CreateProductRequest) (pKey *order_service.ProductPKey, err error){
	newUUID, err := exec.Command("uuidgen").Output()
	if err != nil{
		return nil, err
	}
	uuid := strings.TrimSpace(string(newUUID))

	query:= `
	INSERT INTO products (
		id,
		name,
		price
	) values (
		$1,
		$2,
		$3
	)` 

	_, err = o.db.Exec(ctx, query,
		uuid,
		product.Name,
		product.Price,
	)
	if err != nil{
		return nil, err
	}

	return &order_service.ProductPKey{Id: uuid}, nil
}

func (o *productRepo) GetById(ctx context.Context, pKey *order_service.ProductPKey) (resp *order_service.Product, err error) {
	var query string

	resp = &order_service.Product{}

	query =` 
	SELECT 
		id, 
		name,
		price
	FROM products 
	WHERE id = $1`

	err = o.db.QueryRow(ctx, query, pKey.Id,).Scan(
		&resp.Id,
		&resp.Name,
		&resp.Price,
	)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (o *productRepo) GetAll(ctx context.Context, req *order_service.GetAllProductsRequest) (resp *order_service.GetAllProductsResponse, err error) {
	resp = &order_service.GetAllProductsResponse{}
	var (
		query string
		filter = " WHERE TRUE "
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
	SELECT 
		id, 
		name,
		price
	FROM products`

	query += filter + offset + limit
	
	rows, err := o.db.Query(ctx, query)
	if err != nil{
		return nil, err
	}
	defer rows.Close()
	for rows.Next(){
		var product order_service.Product
		err := rows.Scan(
			&product.Id,
			&product.Name,
			&product.Price,
		)
		if err != nil{
			return nil, err
		}
		resp.Products = append(resp.Products, &product)
	}
	return resp, nil
}

func (o *productRepo) Delete(ctx context.Context, pKey *order_service.ProductPKey) (err error) {
	query := `
	DELETE FROM products
	WHERE id = $1`

	_, err = o.db.Exec(ctx, query, pKey.Id)
	if err != nil{
		return err
	}
	return nil
}

func (o *productRepo) Update(ctx context.Context, pKey *order_service.UpdateProduct) (err error) {
	return nil
}
