package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

var globalClient *mongo.Client

func init() {
	client, err := DBinstance()
	if err != nil {
		log.Fatal("Failed to connect to MongoDB client:%v", err)
	}
	globalClient = client
}
func DBinstance() (*mongo.Client, error) {
	MongoDb := os.Getenv("DB_DATABASE") //从环境变量中获取数据库连接地址，这个要去看altas的官方文档

	if MongoDb == "" {
		log.Fatal("Set your 'MONGODB_URI' environment variable. " +
			"See: " +
			"www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(MongoDb)) //连接数据库
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to MongoDB: %w", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) //设置超时时间
	defer cancel()                                                           //关闭连接
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("Failed to ping to MongoDB: %w", err)
	}
	fmt.Println("Connected to MongoDB")
	return client, nil
}
func New() (*mongo.Database, error) {
	client, err := DBinstance()
	if err != nil {
		return nil, err
	}
	return client.Database("Gomongodb"), nil
}

type Service struct {
	db *mongo.Client
}

func (s *Service) OpenCollection(collectionName string) *mongo.Collection {
	collection := s.db.Database("Gomongodb").Collection(collectionName)
	return collection
}
