package main

import (
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hegervalesin/helplinego/internal/config"
	"github.com/hegervalesin/helplinego/internal/handlers"
	"github.com/hegervalesin/helplinego/internal/middleware"
	"github.com/hegervalesin/helplinego/internal/repository"
	"github.com/hegervalesin/helplinego/internal/services"
	"github.com/joho/godotenv"
)

func main() {
	// Carregar variáveis de ambiente
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, using environment variables")
	}

	// Configurar o banco de dados
	db := config.SetupDatabase()

	// Inicializar repositórios
	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)
	// Adicione outros repositórios conforme necessário

	// Inicializar serviços
	authService := services.NewAuthService(userRepo)
	// Adicione outros serviços conforme necessário

	// Inicializar handlers
	authHandler := handlers.NewAuthHandler(authService)
	// Adicione outros handlers conforme necessário

	sectorRepo := repository.NewSectorRepository(db)
	sectorService := services.NewSectorService(sectorRepo)
	sectorHandler := handlers.NewSectorHandler(sectorService)

	facilityRepo := repository.NewFacilityRepository(db)
	facilityService := services.NewFacilityService(facilityRepo)
	facilityHandler := handlers.NewFacilityHandler(facilityService)

	sevicesRepo := repository.NewServiceRepository(db)
	sevicesService := services.NewServiceService(sevicesRepo, sectorRepo)
	sevicesHandler := handlers.NewServiceHandler(sevicesService)

	// Configurar o router
	r := gin.Default()

	// Configurar CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://130.11.0.59:3000"}, // Adicione seus domínios permitidos
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "Person"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Rotas públicas
	r.POST("/api/auth/login", authHandler.Login)
	r.POST("/api/auth/register", authHandler.Register)
	r.GET("/api/services", sevicesHandler.List)
	r.GET("/api/facility/complete", facilityHandler.List)
	r.GET("/api/sectors", sectorHandler.List)

	// Rotas protegidas
	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware())
	{
		// Rotas de usuário
		users := api.Group("/users")
		{
			// Adicione rotas para usuários
			users.POST("/", userHandler.Create)
			users.GET("/", userHandler.List)
			users.GET("/:id", userHandler.GetByID)
			users.PUT("/:id", userHandler.Update)
			users.DELETE("/:id", userHandler.Delete)
		}

		// Rotas de setores
		sectors := api.Group("/sectors")
		{
			sectors.POST("/", sectorHandler.Create)
			//sectors.GET("/", sectorHandler.List)
			sectors.GET("/:id", sectorHandler.GetByID)
			sectors.PUT("/:id", sectorHandler.Update)
			sectors.DELETE("/:id", sectorHandler.Delete)
		}

		// Rotas de serviços
		services := api.Group("/services")
		{
			// Adicione rotas para serviços
			services.POST("/", sectorHandler.Create)
			services.GET("/", sectorHandler.List)
			services.GET("/:id", sectorHandler.GetByID)
			services.PUT("/:id", sectorHandler.Update)
			services.DELETE("/:id", sectorHandler.Delete)
		}

		// Rotas de facilities
		facilities := api.Group("/facility")
		{
			// Adicione rotas para facilities
			facilities.POST("/", sectorHandler.Create)
			facilities.GET("/", sectorHandler.List)
			facilities.GET("/complete/", sectorHandler.List)
			facilities.GET("/:id", sectorHandler.GetByID)
			facilities.PUT("/:id", sectorHandler.Update)
			facilities.DELETE("/:id", sectorHandler.Delete)
		}

		// Rotas de andares
		floors := api.Group("/floors")
		{
			// Adicione rotas para andares
			floors.POST("/", sectorHandler.Create)
			floors.GET("/", sectorHandler.List)
			floors.GET("/:id", sectorHandler.GetByID)
			floors.PUT("/:id", sectorHandler.Update)
			floors.DELETE("/:id", sectorHandler.Delete)
		}

		// Rotas de salas
		rooms := api.Group("/rooms")
		{
			// Adicione rotas para salas
			rooms.POST("/", sectorHandler.Create)
			rooms.GET("/", sectorHandler.List)
			rooms.GET("/:id", sectorHandler.GetByID)
			rooms.PUT("/:id", sectorHandler.Update)
			rooms.DELETE("/:id", sectorHandler.Delete)
		}
	}

	// Rotas de admin
	admin := r.Group("/api/admin")
	admin.Use(middleware.AuthMiddleware(), middleware.RoleMiddleware("admin"))
	{
		// Rotas administrativas
		// Adicione rotas administrativas conforme necessário

	}

	// Iniciar o servidor
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server running on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
