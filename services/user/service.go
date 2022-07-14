package user

import (
	"fmt"
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

func (s *Service) Create(input UserInput) (*User, error) {
	var u User
	// s.Log.Warn().Msg("Todo: Create() in services/test/service.go")

	//TODO validation

	if input.UserName != "" {
		u.UserName = input.UserName
	}

	if input.Email != "" {
		u.Email = input.Email
	}

	if input.FirstName != "" {
		u.FirstName = input.FirstName
	}

	if input.LastName != "" {
		u.LastName = input.LastName
	}

	if input.Hash != "" {

		u.Hash = input.Hash

		fmt.Println("hash: ", input.Hash)

	}

	timeNow := time.Now().Unix()
	u.CreatedAt = timeNow
	u.UpdatedAt = timeNow

	createdUser, err := s.Store.Create(&u)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

func (s *Service) GetAll() ([]*User, error) {
	//s.Log.Warn().Msg("TODO: GetAll() in services/test/service.go")

	return s.Store.GetAll()

}

func (s *Service) Update(id string, input *UserInput) (*User, error) {
	u, err := s.Store.GetByID(id)
	if err != nil {
		return nil, err
	}

	if input.UserName != "" {
		u.UserName = input.UserName
	}

	if input.Email != "" {
		u.Email = input.Email
	}

	if input.FirstName != "" {
		u.FirstName = input.FirstName
	}

	if input.LastName != "" {
		u.LastName = input.LastName
	}

	if input.Hash != "" {
		return nil, ErrPasswordChangeNotAllowed
	}

	u.UpdatedAt = time.Now().Unix()

	return s.Store.Update(id, u)
}

func (s *Service) Delete(id string) error {

	return s.Store.Delete(id)
}

//TODO: no hashing needed the client should already hash before sending! delete that
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
