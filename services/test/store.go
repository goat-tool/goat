package test

// "context"
// "time"

// "goat/config"
// "goat/log"

// "go.mongodb.org/mongo-driver/bson"
// "go.mongodb.org/mongo-driver/bson/primitive"
// "go.mongodb.org/mongo-driver/mongo"

// type Store struct {
// 	logger     log.Logger
// 	config     *config.Config
// 	collection *mongo.Collection
// }

// func NewStore(logger log.Logger, config *config.Config, client *mongo.Client) *Store {
// 	collection := client.Database(config.Database.Name).Collection("users")

// 	return &Store{
// 		logger:     logger.WithPrefix("store.user"),
// 		collection: collection,
// 	}
// }

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

// func (s *Store) GetAll() ([]*User, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
// 	defer cancel()

// 	var users []*User

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

// 	return users, nil
// }

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
