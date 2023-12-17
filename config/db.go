package config

import (
	"os"
	"strconv"
	"time"
)

type DB struct {
	Host        string
	Port        string
	SslMode     string
	Name        string
	User        string
	Password    string
	MaxOpenConn int
	MaxIdleConn int
	MaxConnTime time.Duration
}

var db = &DB{}

func ConfigDB() *DB {
	return db
}

func ConfigsDB() {
	db.Host = os.Getenv("DB_HOST")
	db.Port = os.Getenv("DB_PORT")
	db.SslMode = os.Getenv("DB_SSLMODE")
	db.Name = os.Getenv("DB_NAME")
	db.User = os.Getenv("DB_USER")
	db.Password = os.Getenv("DB_PASSWORD")
	db.MaxOpenConn, _ = strconv.Atoi(os.Getenv("DB_MAXOPENCONN"))
	db.MaxIdleConn, _ = strconv.Atoi(os.Getenv("DB_MAXIDLECONN"))
	db.MaxConnTime, _ = time.ParseDuration(os.Getenv("DB_MAXCONNTIME"))
}
