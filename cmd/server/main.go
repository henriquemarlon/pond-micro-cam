package main

import (
	"log"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	_ "github.com/henriquemarlon/pond-micro-cam/backend/api"
	"github.com/henriquemarlon/pond-micro-cam/backend/configs"
	"github.com/henriquemarlon/pond-micro-cam/backend/internal/infra/repository"
	"github.com/henriquemarlon/pond-micro-cam/backend/internal/infra/web/handler"
	"github.com/henriquemarlon/pond-micro-cam/backend/internal/infra/web/middleware"
	"github.com/henriquemarlon/pond-micro-cam/backend/internal/usecase"
	"github.com/swaggo/gin-swagger"
)

//	@title			Ponderada API
//	@version		1.0
//	@description	This is a.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	Manager API Support
//	@contact.url	https://github.com/henriquemarlon/pond-micro-cam/backend
//	@contact.email	gomedicine@inteli.edu.br

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host	localhost
//	@BasePath	/api/v1
// 	@query.collection.format multi

func main() {

	/////////////////////// Configs /////////////////////////

	db := configs.SetupPostgres()
	defer db.Close()

	///////////////////////// Gin ///////////////////////////

	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true, // TODO: change to false and make it for production
		AllowMethods:     []string{"PUT", "PATCH, POST, GET, OPTIONS, DELETE"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	api := router.Group("/api/v1")
	api.Use(middleware.AuthMiddleware())

	///////////////////// Swagger //////////////////////

	api.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	///////////////////// Healthcheck //////////////////////

	//TODO: "http://localhost:8080/api/healthz" is the best pattern for healthcheck?

	router.GET("/api/v1/healthz", handler.HealthCheckHandler)

	//////////////////////// User ///////////////////////////

	userRepository := repository.NewUserRepositoryPostgres(db)
	userUseCase := usecase.NewUserUseCase(userRepository)
	userHandlers := handler.NewUserHandlers(userUseCase)

	{
		userGroup := api.Group("/users")
		{
			userGroup.POST("", userHandlers.CreateUser)
			userGroup.GET("", userHandlers.FindAllUsersHandler)
			userGroup.GET("/:id", userHandlers.FindUserByIdHandler)
			userGroup.POST("/login", userHandlers.LoginUserHandler)
		}
	}

	{
		login := api.Group("/login")
		{
			login.POST("", userHandlers.LoginUserHandler)
		}
	}

	err := router.Run(":8080")
	if err != nil {
		log.Fatal("Error running server:", err)
	}
}
