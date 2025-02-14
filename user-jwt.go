package main

// import (
// 	"context"
// 	"log"
// 	"net/http"
// 	"os"
// 	"os/signal"
// 	"syscall"
// 	"time"

// 	"github.com/dgrijalva/jwt-go"
// 	"github.com/gin-gonic/gin"
// 	"github.com/joho/godotenv"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// var jwtKey = []byte("secret_key")

// type Claims struct {
// 	Username string `json:"username"`
// 	jwt.StandardClaims
// }

// func GenerateToken(username string) (string, error) {
// 	expirationTime := time.Now().Add(24 * time.Hour)
// 	claims := &Claims{
// 		Username: username,
// 		StandardClaims: jwt.StandardClaims{
// 			ExpiresAt: expirationTime.Unix(),
// 		},
// 	}
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	return token.SignedString(jwtKey)
// }

// func JWTMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		tokenStr := c.GetHeader("Authorization")
// 		if tokenStr == "" {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
// 			c.Abort()
// 			return
// 		}

// 		claims := &Claims{}
// 		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
// 			return jwtKey, nil
// 		})

// 		if err != nil || !token.Valid {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
// 			c.Abort()
// 			return
// 		}

// 		c.Set("username", claims.Username)
// 		c.Next()
// 	}
// }

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

// type UserService struct {
// 	repo *UserRepository
// }

// func (s *UserService) CreateUser(user User) (*User, error) {
// 	return s.repo.Insert(user)
// }

// func (s *UserService) GetUsers() ([]User, error) {
// 	return s.repo.GetAll()
// }

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

// 	token, _ := GenerateToken(user.Name)
// 	c.JSON(201, gin.H{"user": createdUser, "token": token})
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
// 	r.GET("/users", JWTMiddleware(), userHandler.GetUsers)
// 	return r
// }

// func main() {
// 	if err := godotenv.Load(); err != nil {
// 		log.Fatal("Error loading .env file")
// 	}

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
