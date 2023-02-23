package usecase

import (
	"database/sql"
)

type MarketStorage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) Storage {
	return &MarketStorage{db: db}
}

func (s *MarketStorage) Add(name string, category string) error {
	return nil
}
