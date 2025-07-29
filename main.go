package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
	"github.com/xo-yosi/Talent-SMPS/internal/app/database"
	"github.com/xo-yosi/Talent-SMPS/internal/app/handler"
	"github.com/xo-yosi/Talent-SMPS/internal/app/infra"
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
	s3Client := infra.NewClient()
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

	StudentRepo := postgres.NewStudentPostgres(db)
	fmt.Println("Student repository initialized successfully!")
	StudentService := services.NewStudentService(StudentRepo, s3Client)
	fmt.Println("Student service initialized successfully!")
	StudentHandler := handler.NewStudentHandler(StudentService, StudentRepo)
	fmt.Println("Student handler initialized successfully!")


	MealRepo := postgres.NewMealPostgres(db)
	fmt.Println("Meal repository initialized successfully!")
	MealHandler := handler.NewMealHandler(MealRepo, StudentRepo)
	fmt.Println("Meal handler initialized successfully!")

	c := cron.New()

	_, err = c.AddFunc("0 0 * * *", func() {
		log.Println("⏰ Resetting all student meal flags...")
		if err := StudentRepo.ResetAllMeals(); err != nil {
			log.Println("❌ Failed to reset meals:", err)
		} else {
			log.Println("✅ All student meals reset to false")
		}
	})
	if err != nil {
		log.Fatalf("❌ Failed to schedule meal reset: %v", err)
	}
	c.Start()

	fmt.Println("✅ Meal reset scheduler started. Running...")

	r := gin.Default()

	routes.SetupUserRoutes(r, UserHandler)
	fmt.Println("User routes set up successfully!")
	routes.SetupStudentRoutes(r, StudentHandler)
	fmt.Println("Student routes set up successfully!")
	routes.SetupMealRoutes(r, MealHandler)
	fmt.Println("Meal routes set up successfully!")

	if err := r.Run(":" + cfg.AppPort); err != nil {
		fmt.Println("Error starting the server:", err)
		return
	}
	fmt.Println("Server started successfully on", cfg.AppPort)

	fmt.Println("DTMS Backend is running!")

}
