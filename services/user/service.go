package user

// "time"

// "goat/log"

// "go.mongodb.org/mongo-driver/bson/primitive"
// "golang.org/x/crypto/bcrypt"

type Service struct {
	// logger log.Logger
	// Store  *Store
}

func NewService() *Service {
	return &Service{
		// logger: logger.WithPrefix("service.user"),
		// Store:  store,
	}
}

// func NewService(logger log.Logger, store *Store) *Service {
// 	return &Service{
// 		logger: logger.WithPrefix("service.user"),
// 		Store:  store,
// 	}
// }

// func (s *Service) Create(input UserInput) (*User, error) {
// 	var u User

// 	hash, err := s.hashPassword(u.Hash)
// 	if err != nil {
// 		return nil, err
// 	}

// 	//TODO validation
// 	u.Username = input.Username
// 	u.Email = input.Email
// 	u.FirstName = input.FirstName
// 	u.LastName = input.LastName
// 	u.Hash = hash

// 	timeNow := time.Now().Unix()
// 	u.CreatedAt = timeNow
// 	u.UpdatedAt = timeNow

// 	createdUser, err := s.Store.Create(&u)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return createdUser, nil
// }

func (s *Service) GetAll() ([]*User, error) {
	//	fmt.Println("TODO: services/user/service.go GetAll")
	return nil, nil //s.Store.GetAll()
}

// func (s *Service) Update(id string, input UserInput) (*User, error) {
// 	u, err := s.Store.GetByID(id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	//TODO validation
// 	if input.Password != "" {
// 		return nil, ErrPasswordChangeNotAllowed
// 	}

// 	if input.Username != "" {
// 		u.Username = input.Username
// 	}

// 	if input.Email != "" {
// 		u.Email = input.Email
// 	}

// 	if input.FirstName != "" {
// 		u.FirstName = input.FirstName
// 	}

// 	if input.LastName != "" {
// 		u.LastName = input.LastName
// 	}

// 	u.UpdatedAt = time.Now().Unix()

// 	return s.Store.Update(u)
// }

// func (s *Service) Delete(id string) error {
// 	idObj, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		s.logger.Debugf("failed to parse user id: %v", err)
// 		return ErrIdParseFailed
// 	}

// 	return s.Store.Delete(idObj)
// }

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
