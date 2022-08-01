package test

import (
	"goat/conf"
	"time"

	"git.bitcubix.io/go/log"
	"gorm.io/gorm"
)

type Service struct {
	Log   *log.Logger
	Store *Store
}

func NewService(log *log.Logger, conf *conf.Config, db *gorm.DB) *Service {
	return &Service{
		Log:   log,
		Store: NewStore(log, conf, db),
	}
}

func (s *Service) Create(input TestInput) (*Test, error) {
	var t Test

	//TODO validation
	if input.TestName != "" {
		t.TestName = input.TestName
	}

	if input.TestValue != "" {
		t.TestValue = input.TestValue
	}

	timeNow := time.Now().Unix()
	t.CreatedAt = timeNow
	t.UpdatedAt = timeNow

	createdTest, err := s.Store.Create(&t)
	if err != nil {
		return nil, err
	}

	return createdTest, nil
}

func (s *Service) GetByID(id string) (*Test, error) {

	return s.Store.GetByID(id)

}

func (s *Service) GetAll() ([]*Test, error) {

	return s.Store.GetAll()

}

func (s *Service) Update(id string, input *TestInput) (*Test, error) {
	t, err := s.Store.GetByID(id)
	if err != nil {
		return nil, err
	}

	//TODO validation
	if input.TestName != "" {
		t.TestName = input.TestName
	}

	if input.TestValue != "" {
		t.TestValue = input.TestValue
	}

	t.UpdatedAt = time.Now().Unix()

	return s.Store.Update(id, t)
}

func (s *Service) Delete(id string) error {

	return s.Store.Delete(id)
}
