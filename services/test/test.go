package test

// import "go.mongodb.org/mongo-driver/bson/primitive"

type Test struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	TestName  string `json:"testname"`
	TestValue string `json:"testvalue"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

type TestInput struct {
	TestName  string `json:"testname"`
	TestValue string `json:"testvalue"`
}
