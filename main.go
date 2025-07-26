package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/xo-yosi/Talent-SMPS/internal/app/database"
	"github.com/xo-yosi/Talent-SMPS/internal/config"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Println("Error loading configuration:", err)
		return
	}
	fmt.Println("Configuration loaded successfully!")

	db, err := database.Connect(cfg.DatabaseURL)
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	}
	fmt.Println("Database connected successfully!")

	err = database.Migrate(db)
	if err != nil {
		fmt.Println("Error migrating the database:", err)
		return
	}
	fmt.Println("Database migrated successfully!")

	// err = database.Seeders(db)
	// if err != nil {
	// 	fmt.Println("Error seeding data:", err)
	// 	return
	// }
	// fmt.Println("Data seeded successfully!")
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
	fmt.Printf("Hashed Password: %s\n", hashedPassword)

	r := gin.Default()

	if err := r.Run(":" + cfg.AppPort); err != nil {
		fmt.Println("Error starting the server:", err)
		return
	}
	fmt.Println("Server started successfully on", cfg.AppPort)

	fmt.Println("DTMS Backend is running!")

}
