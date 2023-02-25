package usecase

import (
	"GoStorageService/internal/entity"
	"database/sql"
)

type MarketStorage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) Storage {
	return &MarketStorage{db: db}
}

func (s *MarketStorage) Add(data []entity.Data, name string, category string) error {
	return nil
}

func (s *MarketStorage) clear(name string, category string) error {
	return nil
}
