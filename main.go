package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/xo-yosi/Talent-SMPS/internal/app/database"
	"github.com/xo-yosi/Talent-SMPS/internal/app/handler"
	"github.com/xo-yosi/Talent-SMPS/internal/app/postgres"
	"github.com/xo-yosi/Talent-SMPS/internal/app/services"

	"github.com/xo-yosi/Talent-SMPS/internal/app/routes"
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

	UserRepo := postgres.NewUserPostgres(db)
	fmt.Println("User repository initialized successfully!")

	UserService := services.NewUserService(UserRepo)
	fmt.Println("User service initialized successfully!")

	UserHandler := handler.NewUserHandler(UserService, UserRepo)
	fmt.Println("User handler initialized successfully!")

	r := gin.Default()

	routes.SetupUserRoutes(r, UserHandler)
	fmt.Println("User routes set up successfully!")

	if err := r.Run(":" + cfg.AppPort); err != nil {
		fmt.Println("Error starting the server:", err)
		return
	}
	fmt.Println("Server started successfully on", cfg.AppPort)

	fmt.Println("DTMS Backend is running!")

}
