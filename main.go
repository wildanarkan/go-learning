package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"wildanarkan.com/go-learning/controllers"
	"wildanarkan.com/go-learning/models"
)

func main() {
	// Ubah DSN sesuai kredensial Laragon
	dsn := "root:@tcp(127.0.0.1:3306)/go_crud?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Koneksi database gagal:", err)
		return
	}

	// Test koneksi
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Gagal mendapatkan database:", err)
		return
	}

	// Test ping database
	err = sqlDB.Ping()
	if err != nil {
		log.Fatal("Gagal ping database:", err)
		return
	}

	log.Println("Koneksi database berhasil!")

	// Migrasi database
	err = db.AutoMigrate(&models.Book{})
	if err != nil {
		log.Fatal("Gagal migrasi database:", err)
		return
	}

	// Inisialisasi controller
	bookController := &controllers.BookController{DB: db}

	// Setup router
	r := gin.Default()

	// Routes
	v1 := r.Group("/api/v1")
	{
		books := v1.Group("/books")
		{
			books.GET("/", bookController.GetBooks)
			books.GET("/:id", bookController.GetBook)
			books.POST("/", bookController.CreateBook)
			books.PUT("/:id", bookController.UpdateBook)
			books.DELETE("/:id", bookController.DeleteBook)
		}
	}

	// Run server
	log.Println("Server berjalan di http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Gagal menjalankan server:", err)
	}
}
