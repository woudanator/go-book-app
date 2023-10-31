package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	Db *gorm.DB
)

func Connect() {
	var err error
	var dsn string = "user=reyn password='R@Private2002' port=5432 dbname=bookstore sslmode=disabled TimeZone=Africa/Johannesburg"
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}

func GetDb() *gorm.DB {
	return Db
}
