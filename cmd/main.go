package main

import (
	"log"

	"github.com/ashiruhabeeb/simple-library/app"
	"github.com/ashiruhabeeb/simple-library/app/db"
	"github.com/ashiruhabeeb/simple-library/pkg/config"
)

func main() {
	cfg, err := config.LoadConfig("./pkg/config/")
	if err != nil {
		log.Fatalf("[ERROR] config.LoadConfig func failure: %v", err)
	} else {
		log.Println("[INIT] App configuration succesfully loaded..⛓️")
	}

	db, err := db.ConnectPSQL(cfg)
	if err != nil {
		log.Fatalf("[ERROR] %v", err)
	} else {
		log.Println("[INIT] PSQL connection successfully established..📡")
	}

	app.SetUpRoutes(db, cfg)
}
