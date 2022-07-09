package core

// "context"
// "time"

// "go.mongodb.org/mongo-driver/mongo"
// "go.mongodb.org/mongo-driver/mongo/options"
// "go.mongodb.org/mongo-driver/mongo/readpref"

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/lib/pq"
)

func (c *Core) setupDatabase() {
	c.Log.Info().Msg("setting up database")

	// convert string to integer

	number, err := strconv.Atoi(c.conf.Database.Port)

	fmt.Println("number:", number)
	if err != nil {
		fmt.Println("error:", err)
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", c.conf.Database.Host, number, c.conf.Database.Username, c.conf.Database.Password, c.conf.Database.Name)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("error connecting!")
		c.Log.Panic().Err(err).Msg("error connecting db")
		//panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("error pinging!")
		c.Log.Error().Err(err).Msg("error pinging db")
	}

	fmt.Println("Successfully connected!")

	// client, err := mongo.NewClient(options.Client().ApplyURI(c.Config.Database.URI()))
	// if err != nil {
	// 	logger.Fatalf("failed to create database client: %v", err)
	// }

	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// err = client.Connect(ctx)
	// if err != nil {
	// 	logger.Fatalf("failed to connect ot database: %v", err)
	// }

	// err = client.Ping(ctx, readpref.Primary())
	// if err != nil {
	// 	logger.Fatalf("connection to database server failed: %v", err)
	// } else {
	// 	logger.Debugln("successfully pinged the database server")
	// }

	// c.Database = client

	// c.registerShutdownFunc(func() error {
	// 	err = client.Disconnect(ctx)
	// 	if err != nil {
	// 		logger.Errorf("failed to disconnect from database: %v", err)
	// 		return err
	// 	}

	// 	logger.Debugln("database connection closed")
	// 	return nil
	// })
}