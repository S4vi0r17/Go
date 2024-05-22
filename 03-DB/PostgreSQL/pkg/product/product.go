package product

import "time"

type Model struct {
	ID           uint
	Name         string
	Observations string
	Price        float64
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Models []*Model

type Storage interface {
	Create(*Model) error
	Update(*Model) error
	GetAll() (Models, error)
	GetByID(uint) (*Model, error)
	Delete(uint) error
}
