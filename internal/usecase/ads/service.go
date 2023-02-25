package usecase

import (
	"GoStorageService/internal/entity"
	"errors"

	gojson "github.com/goccy/go-json"
	"github.com/sirupsen/logrus"
)

type MarketService struct {
	storage Storage
}

func NewService(s Storage) Service {
	return &MarketService{s}
}

func (s *MarketService) Add(payload []byte, name string, category string) error {
	var data []entity.Data

	if err := gojson.Unmarshal(payload, &data); err != nil {
		logrus.Errorf("error unmarshal data: %v", err)
		return errors.New("error parce json obj check")
	}

	err := s.storage.Add(data, name, category)
	if err != nil {
		return err
	}

	return nil
}
