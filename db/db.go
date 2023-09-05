package db

import (
	"fmt"
	"go-ecom/models"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	conf := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, username, password, dbName, port)
	db, err := gorm.Open(postgres.Open(conf), &gorm.Config{})

	if err != nil {
		fmt.Println("db err: ", err)
		os.Exit(-1)
	}

	// ДЛЯ ТЕСТА удаляет все таблицы и включает режим отладки
	env := os.Getenv("ENV")
	if env == "dev" {
		db = db.Debug()
		ClearEverything(db)
	}

	db.AutoMigrate(&models.Users{}, &models.Orders{}, &models.Items{}, &models.OrdersItems{})

	fmt.Println("База данных успешно подключена")
	DB = db

	// ДЛЯ ТЕСТА генерирует тестовые данные
	if env == "dev" {
		Seed()
	}
}

func ClearEverything(db *gorm.DB) {
	err0 := db.Exec(`DROP TABLE IF EXISTS users, items, orders, orders_items CASCADE`).Error

	fmt.Printf("Deleting the records: \n%v", err0)
}
