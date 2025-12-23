package invoiceitem

import "time"

type Model struct {
	Id            uint
	InvoiceHeader uint
	ProductId     uint
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type Storage interface {
	Migrate() error
}

type Service struct {
	storage Storage
}

func NewService(s Storage) *Service {
	return &Service{s}
}

func (s *Service) Migrate() error {
	return s.storage.Migrate()
}
