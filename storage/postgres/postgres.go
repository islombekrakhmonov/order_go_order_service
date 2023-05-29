package postgres

import (
	"context"
	"fmt"
	"order_service/config"
	"order_service/storage"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Store struct {
	db       *pgxpool.Pool
	order    storage.OrderRepoI
	product storage.ProductRepoI
}

func NewConnectPostgresql(cfg *config.Config) (storage.StorageI, error) {
	config, err := pgxpool.ParseConfig(fmt.Sprintf(
		"host=%s user=%s dbname=%s password=%s port=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresUser,
		cfg.PostgresDatabase,
		cfg.PostgresPassword,
		cfg.PostgresPort,
	))
	if err != nil {
		return nil, err
	}

	pgpool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}

	return &Store{
		db:       pgpool,
		order:  NewOrderRepo(pgpool),	
	}, nil
}

func (s *Store) CloseDB() {
	s.db.Close()
}

func (s *Store) Order() storage.OrderRepoI {
	if s.order == nil {
		s.order = NewOrderRepo(s.db)
	}

	return s.order
}

func (s *Store) Product() storage.ProductRepoI {
	if s.product == nil {
		s.product = NewProductRepo(s.db)
	}

	return s.product
}
