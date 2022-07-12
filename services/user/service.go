package user

// "goat/log"

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

func (s *Service) GetAll() ([]*User, error) {
	//	fmt.Println("TODO: services/user/service.go GetAll")
	return nil, nil //s.Store.GetAll()
}
