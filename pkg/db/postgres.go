package db

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/wisnunu254/api-auth-golang/config"
)

type DB struct{ *sqlx.DB }

var defaultDB = &DB{}

func (db *DB) connect(cfg *config.DB) (err error) {
	uriDB := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host,
		cfg.Port,
		cfg.Name,
		cfg.User,
		cfg.Password,
		cfg.SslMode,
	)
	db.DB, err = sqlx.Connect("postgres", uriDB)
	if err != nil {
		return err
	}

	db.SetConnMaxIdleTime(time.Duration(cfg.MaxIdleConn))
	db.SetMaxIdleConns(int(cfg.MaxIdleConn))
	db.SetConnMaxLifetime(cfg.MaxConnTime)

	if err := db.Ping(); err != nil {
		return err
	}

	return nil

}

func StartDB() *DB {
	return defaultDB
}

func Connects() error {
	return defaultDB.connect(config.ConfigDB())
}
