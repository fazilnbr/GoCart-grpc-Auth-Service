package db

import (
	"fmt"

	"github.com/fazilnbr/banking-grpc-auth-service/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(cfg config.Config) (*gorm.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", "localhost", "postgres", "test", "5432", "1111")
	db, dbErr := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	// db.AutoMigrate(&domain.Users{})

	return db, dbErr
}
