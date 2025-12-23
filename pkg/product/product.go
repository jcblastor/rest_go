package product

import (
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
