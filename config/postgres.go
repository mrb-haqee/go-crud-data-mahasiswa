package config

import (
	"fmt"

	"github.com/mrb-haqee/go-crud-data-mahasiswa/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)



var dbCred model.Credential = model.Credential{
	Host:         "localhost",
	Username:     "postgres",
	Password:     "mrb28",
	DatabaseName: "crud_mahasiswa",
	Port:         5432,
	Schema:       "public",
}

func Connect() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta", dbCred.Host, dbCred.Username, dbCred.Password, dbCred.DatabaseName, dbCred.Port)
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return conn, nil

}


