package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

// Warehouse model
type Warehouse struct {
	ID		uint	`gorm:"primaryKey" json:"id"`
	Name	string	`json:"name"`
	Adress	string	`json:"adress"`
	Budget	int64	`json:"budget"`
}

func main() {
	// Подключение к БД
	dsn := "host=localhost user=user password=password dbname=warehouse_db sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	
	db.AutoMigrate(&Warehouse{})

	r := gin.Default()

	r.POST("/warehouses", func(c *gin.Context) {
		var input Warehouse
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		db.Create(&input)
		c.JSON(201, input)
	})

	// Эндпоинт: Список всех складов
	r.GET("/warehouses", func(c *gin.Context) {
		var warehouses []Warehouse
		db.Find(&warehouses)
		c.JSON(200, warehouses)
	})

	// Запуск сервера
	r.Run(":8080")
}