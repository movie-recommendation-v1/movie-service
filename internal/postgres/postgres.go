package postgres

import (
	"database/sql"
	"fmt"
	"movie-service/internal/config"
	"movie-service/internal/logger"
)

func ConnectPostgres() (*sql.DB, error) {
	logs, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}

	conf := config.Load()
	dns := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		conf.DBHOST, conf.DBPORT, conf.DBUSER, conf.DBPASSWORD, conf.DBNAME)
	db, err := sql.Open("postgres", dns)
	if err != nil {
		logs.Error("Error while connecting postgres")
	}
	err = db.Ping()
	if err != nil {
		logs.Error("Error while pinging postgres")
	}
	logs.Info("Successfully connected to postgres")
	return db, nil
}
