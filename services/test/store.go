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

		// 		s.logger.Warnf("failed to insert user: %v", err)
		return nil, ErrInsertFailed
	}

	// 	user.ID = inserted.InsertedID.(primitive.ObjectID)

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

	// 	cursor, err := s.collection.Find(ctx, bson.M{})
	// 	if err != nil {
	// 		s.logger.Warnf("failed to find users: %v", err)
	// 		return nil, ErrFindFailed
	// 	}

	// 	err = cursor.All(ctx, &users)
	// 	if err != nil {
	// 		s.logger.Warnf("failed to load users: %v", err)
	// 		return nil, ErrFindFailed
	// 	}

	// 	if len(users) == 0 {
	// 		return nil, ErrNotFound
	// 	}

	// return &tests, nil
}

func (s *Store) GetByID(id string) (*Test, error) {
	// s.log.Warn().Msg("TODO GetByID in services/test/store.go")
	var test *Test

	result := s.db.First(&test, id)
	if result.Error != nil {
		s.log.Error().Err(ErrNotFound).Str("id", id).Msg("")
		return nil, ErrNotFound
	}

	return test, nil
}

// func (s *Store) GetByUsername(username string) (*User, error) {
// 	return s.getByKeyValue("username", username)
// }

// func (s *Store) GetByEmail(email string) (*User, error) {
// 	return s.getByKeyValue("email", email)
// }

func (s *Store) Update(id string, test *Test) (*Test, error) {

	// 	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	// 	defer cancel()

	result := s.db.Save(&test)
	if result.Error != nil {
		s.log.Error().Err(ErrUpdatedFailed).Str("id", id).Msg("")
		return nil, ErrUpdatedFailed
	}
	// 	_, err := s.collection.UpdateOne(ctx, bson.M{"_id": user.ID}, bson.M{"$set": user})
	// 	if err != nil {
	// 		switch err {
	// 		case mongo.ErrNoDocuments:
	// 			return nil, ErrNotFound
	// 		default:
	// 			s.logger.Warnf("error while updating user: %v", err)
	// 			return nil, ErrUpdatedFailed
	// 		}
	// 	}

	return test, nil
}

// func (s *Store) Delete(id primitive.ObjectID) error {
// 	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	_, err := s.collection.DeleteOne(ctx, bson.M{"_id": id})
// 	if err != nil {
// 		switch err {
// 		case mongo.ErrNoDocuments:
// 			return ErrNotFound
// 		default:
// 			s.logger.Warnf("error while deleting user: %v", err)
// 			return ErrDeleteFailed
// 		}
// 	}

// 	return nil
// }

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
