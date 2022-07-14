package test

import (
	"goat/conf"
	"goat/log"

	"gorm.io/gorm"
)

type Store struct {
	log  *log.Logger
	conf *conf.Config
	db   *gorm.DB
}

func NewStore(log *log.Logger, conf *conf.Config, db *gorm.DB) *Store {

	//TODO: is the conf really needed here?
	return &Store{
		log:  log,
		conf: conf,
		db:   db,
	}
}

func (s *Store) Create(test *Test) (*Test, error) {
	// 	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	// 	defer cancel()

	inserted := s.db.Create(test)
	if inserted.Error != nil {
		return nil, ErrInsertFailed
	}

	return test, nil
}

func (s *Store) GetAll() ([]*Test, error) {
	// 	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	// 	defer cancel()
	var tests []*Test

	result := s.db.Find(&tests)
	if result.Error != nil {
		return nil, ErrFindFailed
	}

	return tests, nil
}

func (s *Store) GetByID(id string) (*Test, error) {
	var test *Test

	result := s.db.First(&test, id)
	if result.Error != nil {
		return nil, ErrNotFound
	}

	return test, nil
}

func (s *Store) Update(id string, test *Test) (*Test, error) {

	// 	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	// 	defer cancel()

	result := s.db.Save(&test)
	if result.Error != nil {
		return nil, ErrUpdatedFailed
	}

	return test, nil
}

func (s *Store) Delete(id string) error {
	// ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	// defer cancel()
	var test *Test

	result := s.db.Delete(&test, id)
	if result.Error != nil {
		return ErrDeleteFailed
	} else if result.RowsAffected < 1 {
		return ErrNotFound

	}

	return nil
}
