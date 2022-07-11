package test

// import "go.mongodb.org/mongo-driver/bson/primitive"

type Test struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	TestValue string `json:"testvalue"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

// type TestInput struct {
// 	Testname  string `json:"testname" validate:"required,min=2,max=50"`
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

// type Test struct {
// 	ID        int		`json:"id" gorm:"primaryKey"`
// 	Testname  string    `json:"testname"`
// 	Email     string    `json:"email"`
// 	FirstName string    `json:"first_name"`
// 	LastName  string    `json:"last_name"`
// 	Hash      string    `json:"hash"`
// 	CreatedAt int64     `json:"created_at"`
// 	UpdatedAt int64     `json:"updated_at"`
// }

// type TestInput struct {
// 	Testname  string `json:"testname" validate:"required,min=2,max=50"`
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
