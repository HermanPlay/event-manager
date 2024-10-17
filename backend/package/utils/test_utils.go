package utils

import (
	"fmt"

	"github.com/HermanPlay/web-app-backend/package/domain/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	test_db_host     = "localhost"
	test_db_port     = 5432
	test_db_user     = "postgres"
	test_db_password = "postgres"
	test_db_name     = "test_db"
)

func ConnectToTestDatabase() *gorm.DB {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%v/%s?sslmode=disable",
		test_db_user,
		test_db_password,
		test_db_host,
		test_db_port,
		test_db_name,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("cannot open db connection")
	}

	db.Migrator().DropTable(&models.User{})
	db.AutoMigrate(&models.User{})
	db.Migrator().DropTable(&models.Event{})
	db.AutoMigrate(&models.Event{})
	db.Migrator().DropTable(&models.EventUser{})
	db.AutoMigrate(&models.EventUser{})

	return db
}
