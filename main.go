package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"users-rest/controller"
	"users-rest/db"
	"users-rest/repository"
	"users-rest/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// import (
// 	"context"
// 	"log"
// 	"net/http"
// 	"os"
// 	"os/signal"
// 	"syscall"
// 	"time"

// 	"github.com/gin-gonic/gin"
// 	"github.com/joho/godotenv"
// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

func SetupRouter(controller *controller.UserController) *gin.Engine {
	r := gin.Default()
	r.POST("/users", controller.CreateUser)
	r.GET("/users", controller.GetUsers)
	return r
}

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to MongoDB
	dbManager := db.NewManager()
	err := dbManager.Connect()
	if err != nil {
		log.Fatal("err: ", err)
	}
	log.Println("db connected...")
	repo := repository.NewUserRepository(dbManager.GetCollection("test", "user"))
	service := service.NewUserService(repo)
	controller := controller.NewUserController(service)

	r := SetupRouter(controller)
	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	// Graceful shutdown
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatal("Server Error: ", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown Failed: ", err)
	}

	log.Println("Server gracefully stopped")
}
