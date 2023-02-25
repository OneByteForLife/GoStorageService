package usecase

import "GoStorageService/internal/entity"

type (
	// Для storage
	Storage interface {
		Add(data []entity.Data, name string, category string) error
	}

	Service interface {
		Add(payload []byte, name string, category string) error
	}
)
