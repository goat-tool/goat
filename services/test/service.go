package test

import (
	"goat/conf"
	"goat/log"
	"time"

	"gorm.io/gorm"
)

type Service struct {
	Log   *log.Logger
	Store *Store
}

func NewService(log *log.Logger, conf *conf.Config, db *gorm.DB) *Service {
	// log.Warn().Msg("TODO: NewService() in services/test/service.go")
	return &Service{
		Log:   log,
		Store: NewStore(log, conf, db),
	}
}

func (s *Service) Create(input TestInput) (*Test, error) {
	var t Test
	// s.Log.Warn().Msg("Todo: Create() in services/test/service.go")

	//TODO validation
	t.TestName = input.TestName
	t.TestValue = input.TestValue

	timeNow := time.Now().Unix()
	t.CreatedAt = timeNow
	t.UpdatedAt = timeNow

	createdTest, err := s.Store.Create(&t)
	if err != nil {
		return nil, err
	}

	return createdTest, nil
}

func (s *Service) GetAll() ([]*Test, error) {
	//s.Log.Warn().Msg("TODO: GetAll() in services/test/service.go")

	return s.Store.GetAll()

}

func (s *Service) Update(id string, input *TestInput) (*Test, error) {
	u, err := s.Store.GetByID(id)
	if err != nil {
		return nil, err
	}

	//TODO validation
	if input.TestName != "" {
		u.TestName = input.TestName
	}

	if input.TestValue != "" {
		u.TestValue = input.TestValue
	}

	u.UpdatedAt = time.Now().Unix()

	return s.Store.Update(id, u)
}

func (s *Service) Delete(id string) error {

	return s.Store.Delete(id)
}

// func (s *Service) hashPassword(password string) (string, error) {
// 	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
// 	if err != nil {
// 		//TODO better error handling
// 		return "", err
// 	}

// 	return string(bytes), nil
// }

// func (s *Service) verifyPassword(hash string, password string) bool {
// 	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
// 	if err != nil {
// 		//TODO add logging
// 		return false
// 	}

// 	return true
// }
