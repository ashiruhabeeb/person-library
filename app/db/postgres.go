package db

import (
	"fmt"
	"log"

	"github.com/ashiruhabeeb/simple-library/app/model"
	"github.com/ashiruhabeeb/simple-library/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectPSQL(cfg *config.AppConfig) (*gorm.DB, error) {
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		cfg.Database.DbHost, cfg.Database.DbUser, cfg.Database.DbPwd, cfg.Database.DbName, cfg.Database.DbPort, cfg.Database.SSLmode)

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("[ERROR] gorm.Open method failure: %v", err)
	}

	// AutoMigrate create needed tables as defined in the model package in PSQL database
	db.AutoMigrate(&model.Books{})
	log.Println("[INIT] Database tables sucessfully migrated..")

	Db, _ := db.DB()
	Db.SetMaxOpenConns(cfg.Database.MaxOpenConnections)
	Db.SetMaxIdleConns(cfg.Database.MaxIdleConnections)

	return db, nil
}
