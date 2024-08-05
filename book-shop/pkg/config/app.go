package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"net/url"
)

var (
	db *gorm.DB
)

var dsn url.URL = url.URL{
	User:     url.UserPassword("postgres", "postgres"),
	Scheme:   "postgres",
	Host:     "localhost:5432",
	Path:     "bookstore",
	RawQuery: (&url.Values{"sslmode": []string{"disable"}}).Encode(),
}

func Connect() {
	d, err := gorm.Open("postgres", dsn.String())
	if err != nil {
		panic(err)
	}

	db = d
}

func GetDb() *gorm.DB {
	return db
}
