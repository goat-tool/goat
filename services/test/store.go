package test

import (
	"goat/conf"

	// "context"
	// "time"
	"goat/log"
)

type Store struct {
	log  *log.Logger
	conf *conf.Config
	//collection *mongo.Collection
}

func NewStore(log *log.Logger, conf *conf.Config) *Store {
	// 	collection := client.Database(config.Database.Name).Collection("users")
	log.Warn().Msg("Todo: NewStore() in services/test/store.go")

	//TODO: is the conf really needed here?
	return &Store{
		log:  log,
		conf: conf,
		//collection: collection,
	}
}

// func (s *Store) Create(user *User) (*User, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	inserted, err := s.collection.InsertOne(ctx, user)
// 	if err != nil {
// 		s.logger.Warnf("failed to insert user: %v", err)
// 		return nil, ErrInsertFailed
// 	}

// 	user.ID = inserted.InsertedID.(primitive.ObjectID)

// 	return user, nil
// }

func (s *Store) GetAll() ([]*Test, error) {
	// 	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	// 	defer cancel()

	var tests []*Test

	s.log.Warn().Msg("Todo: GetAll() in services/test/store.go")

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

	return tests, nil
}

// func (s *Store) GetByID(id string) (*User, error) {
// 	idObj, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		s.logger.Debugf("failed to parse user id: %v", err)
// 		return nil, ErrIdParseFailed
// 	}

// 	return s.getByKeyValue("_id", idObj)
// }

// func (s *Store) GetByUsername(username string) (*User, error) {
// 	return s.getByKeyValue("username", username)
// }

// func (s *Store) GetByEmail(email string) (*User, error) {
// 	return s.getByKeyValue("email", email)
// }

// func (s *Store) Update(user *User) (*User, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

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

// 	return user, nil
// }

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
