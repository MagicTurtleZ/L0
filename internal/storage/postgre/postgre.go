package storage

import (
	"context"
	"fmt"
	"woonbeaj/L0/internal/jsonStruct"

	"github.com/jackc/pgx/v5"
)

type Storage struct {
	db *pgx.Conn
}


func New(storageURL string) (*Storage, error) {
	const op = "storage.potgre.New"

	db, err := pgx.Connect(context.Background(), storageURL)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{db: db}, nil
}

func (s *Storage) Close(ctx context.Context) {
	s.db.Close(ctx)
}

func (s *Storage) Save(orderUid string, orderInfo []byte) error {
	const op = "storage.potgre.Save"

	_, err := s.db.Exec(context.Background(), "INSERT INTO orders (order_uid, order_info) VALUES ($1, $2)", orderUid, orderInfo)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Storage) Get(orderUid string) ([]byte, error) {
	const op = "storage.potgre.Get"

	var resJSON []byte
	err := s.db.QueryRow(context.Background(), "SELECT order_info FROM orders WHERE order_uid = $1", orderUid).Scan(&resJSON)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	
	return resJSON, nil
}

func (s *Storage) GetAll() (*jsonStruct.AllRows, error) {
	const op = "storage.potgre.GetAll"

	rows, err := s.db.Query(context.Background(), "SELECT order_uid, order_info FROM orders")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()
	var (
		key			string
		value  		[]byte
		res	  		jsonStruct.AllRows
	)
	for rows.Next() {
		err = rows.Scan(&key, &value)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		res.OrdersUIDs = append(res.OrdersUIDs, key)
		res.OrderINFOs = append(res.OrderINFOs, value)
	}
	
	return &res, nil
}