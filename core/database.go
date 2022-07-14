package core

import (
	//"github.com/tutorials/go/crud/pkg/models"
	"fmt"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func (c *Core) NewDatabase() *gorm.DB {
	c.Log.Info().Msg("Setup database")

	// convert string to integer
	portInt, err := strconv.Atoi(c.Conf.Database.Port)
	if err != nil {
		fmt.Println("String to int error:", err)
	}
	dbURL := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", c.Conf.Database.Host, portInt, c.Conf.Database.Username, c.Conf.Database.Password, c.Conf.Database.Name)

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil {
		c.Log.Fatal().Err(err).Msg("Error opening db")
	} else {
		c.Log.Info().Msg("DB OK")
	}

	//db.AutoMigrate(&test.Test{})
	return db

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
