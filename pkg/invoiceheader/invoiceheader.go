package invoiceheader

import "time"

type Model struct {
	Id        uint
	Client    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
