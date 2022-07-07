package user

// import "go.mongodb.org/mongo-driver/bson/primitive"

// type User struct {
// 	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
// 	Username  string             `json:"username" bson:"username"`
// 	Email     string             `json:"email" bson:"email"`
// 	FirstName string             `json:"first_name" bson:"first_name"`
// 	LastName  string             `json:"last_name" bson:"last_name"`
// 	Hash      string             `json:"-" bson:"hash"`
// 	CreatedAt int64              `json:"created_at" bson:"created_at"`
// 	UpdatedAt int64              `json:"updated_at" bson:"updated_at"`
// }

// type UserInput struct {
// 	Username  string `json:"username" validate:"required,min=2,max=50"`
// 	Email     string `json:"email" validate:"required,email"`
// 	FirstName string `json:"first_name" validate:"alphanumunicode"`
// 	LastName  string `json:"last_name" validate:"alphanumunicode"`
// 	Password  string `json:"password" validate:"required,password"`
// }

// type PasswordInput struct {
// 	ActivePassword string `json:"active_password" validate:"required,password"`
// 	NewPassword    string `json:"new_password" validate:"required,password"`
// 	RepeatPassword string `json:"repeat_password" validate:"required,password"`
// }
