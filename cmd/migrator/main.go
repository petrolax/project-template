package main

import (
	"log"

	"github.com/carprice-tech/migorm"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gitlab.com/propetrov/project_example/config"
	_ "gitlab.com/propetrov/project_example/migrations"
)

func main() {
	conf, err := config.GetConfig()
	if err != nil {
		log.Fatalln("Can't parse config")
	}

	db, err := gorm.Open("postgres", conf.DB.GetDSN())
	if err != nil {
		log.Fatalln(err)
	}

	migorm.Run(migorm.NewMigrater(db))
}
