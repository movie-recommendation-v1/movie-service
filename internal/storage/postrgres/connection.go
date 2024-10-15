package postrgres

import (
	"database/sql"
	"fmt"
	"movie-service/internal/config"
	"movie-service/internal/logger"

)

func Connect() (*sql.DB, error) {
	conf := config.Load()
	logs, err := logger.NewLogger()
	if err != nil {
		return nil, err
	}
	dns := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		conf.DBHOST, conf.DBPORT, conf.DBUSER, conf.DBPASSWORD, conf.DBNAME)
	db, err := sql.Open("postgres", dns)
	if err != nil {
		logs.Error("Error connecting to database: " + err.Error())
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		logs.Error("Error connecting to database: " + err.Error())
		return nil, err
	}
	return db, nil
}
