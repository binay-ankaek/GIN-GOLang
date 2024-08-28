package main

import (
	"fmt"
	// "github.com/subosito/gotenv"
	"helloapp/internal/adapter/delivery/http"
	"helloapp/internal/app/contact"
	"helloapp/internal/app/middleware"
	"helloapp/internal/app/user"
	"helloapp/internal/config"
	"helloapp/internal/domain/model"
	"helloapp/internal/domain/repository"
	"helloapp/internal/initializers"
	"log"
	"os"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"

	_ "helloapp/docs" // Import the generated docs

	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

// @title               aok-connect Business API
// @version             1.0
// @description         This is a simple RESTful Service API written in Go using Gin web framework
// @securityDefinitions.apikey BearerToken
// @in Header
// @name Authorization

func init() {

	initializers.ConnectToDB()
	err := initializers.DB.AutoMigrate(&model.User{}, &model.Contact{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

}
func main() {
	// gotenv.Load()
	// Load configurations
	cfg := config.Load()
	//initialize database
	db := initializers.DB

	// Initialize repositories
	userRepository := repository.NewUserRepository(db)
	contactRepository := repository.NewContactRepository(db)

	// Initialize services
	userService := user.NewUserService(userRepository, nil)
	contactService := contact.NewContactService(contactRepository, userRepository)

	// Initialize HTTP server
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:8080"}, // Allow your Vue app's origin
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept", "Authorization"},
	}))

	// Initialize HTTP handlers
	authHandler := http.NewAuthHandler(userService)
	contactHandler := http.NewContactHandler(contactService)
	userHandler := http.NewUserHandler(userService)

	// Routes
	api := router.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			// Add other authentication routes (login, update profile, etc.) here if needed
			auth.POST("/login", authHandler.Login)
			auth.GET("/profile", middleware.AuthMiddleware(os.Getenv("SECRET")), authHandler.GetProfile)
			auth.PUT("/profile", middleware.AuthMiddleware(os.Getenv("SECRET")), authHandler.UpdateProfile)
			auth.DELETE("/profile", middleware.AuthMiddleware(os.Getenv("SECRET")), authHandler.DeleteProfile)
			auth.DELETE("/delete-all-users", userHandler.DeleteAllUsers)
		}

		contact := api.Group("/contact")
		{
			contact.POST("/add", middleware.AuthMiddleware(os.Getenv("SECRET")), contactHandler.AddContact)
			contact.GET("/", middleware.AuthMiddleware(os.Getenv("SECRET")), contactHandler.GetContacts)
			contact.PUT("/:id", middleware.AuthMiddleware(os.Getenv("SECRET")), contactHandler.UpdateContact)
			contact.DELETE("/:id", middleware.AuthMiddleware(os.Getenv("SECRET")), contactHandler.DeleteContact)
			contact.GET("/search", middleware.AuthMiddleware(os.Getenv("SECRET")), contactHandler.SearchContact)
			contact.GET("/profile-phone", contactHandler.GetProfileByPhone)
			// Add other contact management routes (update, delete, search, etc.) here if needed
		}
	}

	// Start HTTP server
	port := cfg.Port // Get port from loaded config
	err := router.Run(":" + port)
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}

	fmt.Println("Server started on port", port)
}
