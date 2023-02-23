package usecase

type (
	// Для storage
	Storage interface {
		Add(name string, category string) error
	}

	Service interface {
		Add(payload []byte, name string, category string) error
	}
)
