package main

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
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

// // User model
// type User struct {
// 	ID    string `bson:"_id,omitempty" json:"id"`
// 	Name  string `bson:"name" json:"name"`
// 	Email string `bson:"email" json:"email"`
// }

// // Database connection variables
// var client *mongo.Client

// type UserRepository struct {
// 	collection *mongo.Collection
// }

// // Initialize MongoDB connection
// func ConnectDB() (*mongo.Client, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URI"))
// 	client, err := mongo.Connect(ctx, clientOptions)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return client, nil
// }

// Insert a user
func (r *UserRepository) Insert(user User) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.collection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Get all users
func (r *UserRepository) GetAll() ([]User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	var users []User
	if err = cursor.All(ctx, &users); err != nil {
		return nil, err
	}

	return users, nil
}

// // UserService to handle business logic
// type UserService struct {
// 	repo *UserRepository
// }

// func (s *UserService) CreateUser(user User) (*User, error) {
// 	return s.repo.Insert(user)
// }

// func (s *UserService) GetUsers() ([]User, error) {
// 	return s.repo.GetAll()
// }

// // HTTP Handlers
// type UserHandler struct {
// 	service *UserService
// }

// func (h *UserHandler) CreateUser(c *gin.Context) {
// 	var user User
// 	if err := c.ShouldBindJSON(&user); err != nil {
// 		c.JSON(400, gin.H{"error": err.Error()})
// 		return
// 	}

// 	createdUser, err := h.service.CreateUser(user)
// 	if err != nil {
// 		c.JSON(500, gin.H{"error": "Failed to create user"})
// 		return
// 	}

// 	c.JSON(201, createdUser)
// }

// func (h *UserHandler) GetUsers(c *gin.Context) {
// 	users, err := h.service.GetUsers()
// 	if err != nil {
// 		c.JSON(500, gin.H{"error": "Failed to fetch users"})
// 		return
// 	}

// 	c.JSON(200, users)
// }

// func SetupRouter(userHandler *UserHandler) *gin.Engine {
// 	r := gin.Default()
// 	r.POST("/users", userHandler.CreateUser)
// 	r.GET("/users", userHandler.GetUsers)
// 	return r
// }

// func main() {
// 	// Load environment variables
// 	if err := godotenv.Load(); err != nil {
// 		log.Fatal("Error loading .env file")
// 	}

// 	// Connect to MongoDB
// 	var err error
// 	client, err = ConnectDB()
// 	if err != nil {
// 		log.Fatal("Could not connect to MongoDB: ", err)
// 	}
// 	db := client.Database("testdb")
// 	repo := &UserRepository{collection: db.Collection("users")}
// 	service := &UserService{repo: repo}
// 	handler := &UserHandler{service: service}

// 	r := SetupRouter(handler)
// 	server := &http.Server{
// 		Addr:    ":8080",
// 		Handler: r,
// 	}

// 	// Graceful shutdown
// 	go func() {
// 		if err := server.ListenAndServe(); err != nil {
// 			log.Fatal("Server Error: ", err)
// 		}
// 	}()

// 	quit := make(chan os.Signal, 1)
// 	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
// 	<-quit

// 	log.Println("Shutting down server...")
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()
// 	if err := server.Shutdown(ctx); err != nil {
// 		log.Fatal("Server Shutdown Failed: ", err)
// 	}

// 	log.Println("Server gracefully stopped")
// }
