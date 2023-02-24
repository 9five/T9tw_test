package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	aesKey string
	DB     *gorm.DB
)

func init() {
	aesKey = "astaxie12798akljzmknm.ahkjkljl;k"
	DB = newDBConfig()
}

func newDBConfig() *gorm.DB {
	dsn := "host=127.0.0.1 user=postgres password=postgres dbname=t9Test port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{TablePrefix: "t9Test."},
	})
	if err != nil {
		panic(err.Error())
	}
	return db
}
