package core

import (
	"fmt"
	"strconv"

	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

func (c *Core) setupDatabase() {
	c.Log.Info().Msg("Setup database")

	// convert string to integer
	portInt, err := strconv.Atoi(c.Conf.Database.Port)
	if err != nil {
		fmt.Println("String to int error:", err)
	}

	dbUrl := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", c.Conf.Database.Host, portInt, c.Conf.Database.Username, c.Conf.Database.Password, c.Conf.Database.Name)
	conn, err := pgxpool.Connect(context.Background(), dbUrl)
	if err != nil {
		c.Log.Panic().Err(err).Msg("Setup database connection error")
	}

	defer conn.Close()

	err = conn.Ping(context.Background())
	if err != nil {
		c.Log.Error().Err(err).Msg("Setup database ping error")
	} else {
		c.Log.Info().Msg("Setup database ping successful")
	}

	c.Database = conn

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
