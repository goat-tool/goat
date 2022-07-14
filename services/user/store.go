package user

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

func (s *Store) Create(user *User) (*User, error) {
	// s.log.Warn().Msg("TODO Create() in services/test/store.go")
	// 	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	// 	defer cancel()

	inserted := s.db.Create(user)
	if inserted.Error != nil {

		//s.logger.Warn().Err(err).Msg("failed to insert user")
		return nil, ErrInsertFailed
	}

	return user, nil
}

func (s *Store) GetAll() ([]*User, error) {
	// 	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	// 	defer cancel()
	// s.log.Warn().Msg("TODO GetAll() in services/test/store.go")
	var users []*User

	result := s.db.Find(&users)
	if result.Error != nil {
		return nil, ErrFindFailed
	}

	return users, nil
}

func (s *Store) GetByID(id string) (*User, error) {
	// s.log.Warn().Msg("TODO GetByID in services/test/store.go")
	var user *User

	result := s.db.First(&user, id)
	if result.Error != nil {
		//s.log.Error().Err(ErrNotFound).Str("id", id).Msg("")
		return nil, ErrNotFound
	}

	return user, nil
}

// func (s *Store) GetByUsername(username string) (*User, error) {
// 	return s.getByKeyValue("username", username)
// }

// func (s *Store) GetByEmail(email string) (*User, error) {
// 	return s.getByKeyValue("email", email)
// }

func (s *Store) Update(id string, user *User) (*User, error) {

	// 	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	// 	defer cancel()

	result := s.db.Save(&user)
	if result.Error != nil {
		//s.log.Error().Err(ErrUpdatedFailed).Str("id", id).Msg("")
		return nil, ErrUpdatedFailed
	}

	return user, nil
}

func (s *Store) Delete(id string) error {
	// ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	// defer cancel()
	var user *User

	result := s.db.Delete(&user, id)
	if result.Error != nil {
		//s.log.Error().Err(ErrDeleteFailed).Str("id", id).Msg("")
		return ErrDeleteFailed
	} else if result.RowsAffected < 1 {
		//s.log.Error().Err(ErrNotFound).Str("id", id).Msg("")
		return ErrNotFound

	}

	return nil
}

// func (s *Store) getByKeyValue(key string, value interface{}) (*User, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	var user User

// 	err := s.collection.FindOne(ctx, bson.M{key: value}).Decode(&user)
// 	if err != nil {
// 		switch err {
// 		case mongo.ErrNoDocuments:
// 			return nil, ErrNotFound
// 		default:
// 			s.logger.Warnf("error while finding user: %v", err)
// 			return nil, ErrFindFailed
// 		}
// 	}

// 	return &user, nil
// }
