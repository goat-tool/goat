package user

// import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        int
	Username  string
	Email     string
	FirstName string
	LastName  string
	Hash      string
	CreatedAt int64
	UpdatedAt int64
}
