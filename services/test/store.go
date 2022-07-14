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
	// log.Warn().Msg("Todo: NewStore() in services/test/store.go")

	//TODO: is the conf really needed here?
	return &Store{
		log:  log,
		conf: conf,
		db:   db,
	}
}

func (s *Store) Create(test *Test) (*Test, error) {
	// s.log.Warn().Msg("TODO Create() in services/test/store.go")
	// 	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	// 	defer cancel()

	inserted := s.db.Create(test)
	if inserted.Error != nil {

		//s.logger.Warn().Err(err).Msg("failed to insert user")
		return nil, ErrInsertFailed
	}

	return test, nil
}

func (s *Store) GetAll() ([]*Test, error) {
	// 	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	// 	defer cancel()
	// s.log.Warn().Msg("TODO GetAll() in services/test/store.go")
	var tests []*Test

	result := s.db.Find(&tests)
	if result.Error != nil {
		return nil, ErrFindFailed
	}

	return tests, nil
}

func (s *Store) GetByID(id string) (*Test, error) {
	// s.log.Warn().Msg("TODO GetByID in services/test/store.go")
	var test *Test

	result := s.db.First(&test, id)
	if result.Error != nil {
		// s.log.Error().Err(ErrNotFound).Str("id", id).Msg("")
		return nil, ErrNotFound
	}

	return test, nil
}

func (s *Store) Update(id string, test *Test) (*Test, error) {

	// 	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	// 	defer cancel()

	result := s.db.Save(&test)
	if result.Error != nil {
		// s.log.Error().Err(ErrUpdatedFailed).Str("id", id).Msg("")
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
		// s.log.Error().Err(ErrDeleteFailed).Str("id", id).Msg("")
		return ErrDeleteFailed
	} else if result.RowsAffected < 1 {
		// s.log.Error().Err(ErrNotFound).Str("id", id).Msg("")
		return ErrNotFound

	}

	return nil
}
