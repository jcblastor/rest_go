package invoiceitem

import "time"

type Model struct {
	Id            uint
	InvoiceHeader uint
	ProductId     uint
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

// creamos un nuevo tipo para recibir los productos de la consulta
type Models []*Model

// creamos la intrface Storage
type Storage interface {
	Create(*Model) error
	Update(*Model) error
	GetAll() (Models, error)
	GetById(uint) (*Model, error)
	Delete(uint) error
}
