package Config

import (
	"github.com/souvik666/go_todo/Utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func MakeDBConnection() {
	var err error
	var postgresqlUrl string = Utils.GetEnv("PGURL")
	DB, err = gorm.Open(postgres.Open(postgresqlUrl), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	pg, err := DB.DB()
	if err != nil {
		panic("Failed to get SQL database: " + err.Error())
	}
	err = pg.Ping()
	if err != nil {
		panic("Failed to ping database: " + err.Error())
	}

}
