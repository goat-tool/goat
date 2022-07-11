package core

import (
	//"github.com/tutorials/go/crud/pkg/models"
	"fmt"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func (c *Core) newDatabase() {
	c.Log.Info().Msg("Setup database")
	// gorm *******************************************************************************
	//dbURL := "postgres://pg:pass@localhost:5432/crud"

	// convert string to integer
	portInt, err := strconv.Atoi(c.Conf.Database.Port)
	if err != nil {
		fmt.Println("String to int error:", err)
	}
	dbURL := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", c.Conf.Database.Host, portInt, c.Conf.Database.Username, c.Conf.Database.Password, c.Conf.Database.Name)

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		c.Log.Fatal().Err(err).Msg("Error opening db")
	} else {
		c.Log.Info().Msg("DB OK")
	}
	//db.AutoMigrate(&models.Book{})
	c.Database = db

	// old stuff *******************************************************************************
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
