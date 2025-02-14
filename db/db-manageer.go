package db

import (
	"context"
	"os"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var databaseManager *DatabaseManager

type DatabaseManager struct {
	client    *mongo.Client
	databases map[string]*mongo.Database
	mutex     sync.RWMutex
}

func GetDatabaseManager() *DatabaseManager {
	return databaseManager
}

func NewManager() *DatabaseManager {
	databaseManager = &DatabaseManager{}
	return databaseManager
}

func (db *DatabaseManager) Connect() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URI"))
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return err
	}

	db.client = client

	return nil
}

func (db *DatabaseManager) GetDataBase(name string) *mongo.Database {
	db.mutex.RLock()
	database := db.databases[name]
	db.mutex.RUnlock()
	if database == nil {
		database = db.client.Database(name)
		db.mutex.Lock()
		db.databases[name] = database
		db.mutex.Unlock()
	}
	return database
}

func (db *DatabaseManager) GetCollection(name string) *mongo.Collection {
	return nil
}
