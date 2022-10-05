package main

import (
	"log"
	"reflect"
	"rest-api/config"
	_db "rest-api/pkg/db"
	"rest-api/pkg/models"
	_userRepo "rest-api/pkg/repository/postgres"
	"rest-api/pkg/server"
)

func main() {
	//load config
	cfg, errConfig := config.Init()

	if errConfig != nil {
		log.Fatal("Config loading error: ", errConfig.Error())
	}

	printConfig(cfg)

	db, errDb := _db.InitDB(cfg)

	if errDb != nil {
		log.Fatal("Database initialisation error: ", errDb.Error())
	}

	if cfg.DbMigrate {
		err := db.AutoMigrate(&models.User{})

		if err != nil {
			log.Fatal("Database migration error: ", err.Error())
		}
	}

	userRepo := _userRepo.NewUserRepository(db)
	log.Printf("userRepo: %v", userRepo)

	s := server.NewServer(cfg, userRepo)

	go func() {
		err := server.StartServer(cfg, s)
		if err != nil {
			log.Fatal("Server start error: ", err.Error())
		}
	}()

	quit := make(chan bool)
	<-quit

}

func printConfig(cfg *config.Config) {
	e := reflect.ValueOf(cfg).Elem()

	for i := 0; i < e.NumField(); i++ {
		varName := e.Type().Field(i).Name
		varType := e.Type().Field(i).Type
		varValue := e.Field(i).Interface()
		log.Printf("%v (%v) = %v\n", varName, varType, varValue)
	}
}
