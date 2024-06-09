package product

import (
	"errors"
	"fmt"
	"time"
)

var (
	ErrIdNotFound = errors.New("Product ID not found")
)

type Model struct {
	ID           uint
	Name         string
	Observations string
	Price        float64
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (m *Model) String() string {
	return fmt.Sprintf("%02d | %-20s | %-20s | %5.0f | %10s | %10s", m.ID, m.Name, m.Observations, m.Price, m.CreatedAt.Format("2006-01-02"), m.UpdatedAt.Format("2006-01-02"))
}

type Models []*Model

type Storage interface {
	Migrate() error
	Create(*Model) error
	Update(*Model) error
	GetAll() (Models, error)
	GetByID(uint) (*Model, error)
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

func (s *Service) GetByID(id uint) (*Model, error) {
	return s.storage.GetByID(id)
}

func (s *Service) Update(m *Model) error {
	if m.ID == 0 {
		return ErrIdNotFound
	}

	m.UpdatedAt = time.Now()
	return s.storage.Update(m)
}
