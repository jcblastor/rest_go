package product

import (
	"fmt"
	"strings"
	"time"
)

type Model struct {
	Id           uint
	Name         string
	Observations string
	Price        int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (m *Model) String() string {
	return fmt.Sprintf("%02d | %-20s | %-20s | %5d | %10s |%10s",
		m.Id, m.Name, m.Observations, m.Price,
		m.CreatedAt.Format("2025-12-23"), m.UpdatedAt.Format("2025-12-23"),
	)
}

// creamos un nuevo tipo para recibir los productos de la consulta
type Models []*Model

func (m Models) String() string {
	builder := strings.Builder{}
	fmt.Fprintf(&builder, "%02s | %-20s | %-20s | %5s | %10s |%10s",
		"id", "name", "observations", "price", "created_at", "updated_at\n")

	for _, model := range m {
		builder.WriteString(model.String() + "\n")
	}

	return builder.String()
}

// creamos la intrface Storage
type Storage interface {
	Migrate() error
	Create(*Model) error
	// Update(*Model) error
	GetAll() (Models, error)
	// GetById(uint) (*Model, error)
	// Delete(uint) error
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

func (s *Service) Create(m *Model) error {
	m.CreatedAt = time.Now()
	return s.storage.Create(m)
}

func (s *Service) GetAll() (Models, error) {
	return s.storage.GetAll()
}
