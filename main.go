package main

import (
	"efishery-project/config"
	"efishery-project/router"
	"fmt"
	"log"
)

func main() {
	//Initialize DB
	configdb := config.ConfigurationDB()

	db, err := config.NewConnection(configdb)

	if err != nil {
		log.Panic("Database connection failed")
	}

	fmt.Println("Database connection established", db)

	//Serve
	e := router.InitRouter(db)

	e.Logger.Fatal(e.Start("127.0.0.1:5000"))
}
