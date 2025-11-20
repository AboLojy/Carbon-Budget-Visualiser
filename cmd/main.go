package main

import (
	"database/sql"
	"fmt"
	"log"

	apiserver "github.com/AboLojy/Carbon-Budget-Visualiser/cmd/api"
	configs "github.com/AboLojy/Carbon-Budget-Visualiser/config"
	"github.com/AboLojy/Carbon-Budget-Visualiser/db"
)

func main() {
	db, err := db.NewPostgresDb(configs.Envs)
	if err != nil {
		log.Fatal("Could not connect to the database:", err)
	}
	initStorage(db)
	server := apiserver.NewAPIServer(fmt.Sprintf(":%s", configs.Envs.Port), db)
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}

}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal("Could not ping the database:", err)
	}
	log.Println("DB: Successfully connected!")
}
