package pkg

import "time"

type Model struct {
	Id          uint
	Name        string
	Description string
	Weight      float64
	Max_age     int32
	Create_at   time.Time
	Update_at   time.Time
}

//interface CRUD

type Storage interface {
	Migrate() error
}
type Service struct {
	storage Storage
}

func NewServices(s Storage) *Service {
	return &Service{s}
}

func (s *Service) Migrate() error {
	return s.storage.Migrate()
}
