package main

import (
    "context"
    "database/sql"
    "fmt"
    "log"
    "time"

    "github.com/go-redis/redis/v8"
    "github.com/go-sql-driver/mysql"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
    // MySQL Connection
    cfg := mysql.Config{
        User:   "user",
        Passwd: "password",
        Net:    "tcp",
        Addr:   "mysql:3306",
        DBName: "dbname",
    }
    db, err := sql.Open("mysql", cfg.FormatDSN())
    if err != nil {
        log.Fatalf("Could not connect to MySQL: %v", err)
    }
    defer db.Close()

    fmt.Println("Successfully connected to MySQL")

    // MongoDB Connection
    mongoCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    mongoClient, err := mongo.Connect(mongoCtx, options.Client().ApplyURI("mongodb://mongouser:mongopass@mongodb:27017"))
    if err != nil {
        log.Fatalf("Could not connect to MongoDB: %v", err)
    }
    defer mongoClient.Disconnect(mongoCtx)

    fmt.Println("Successfully connected to MongoDB")

    // Redis Connection
    redisClient := redis.NewClient(&redis.Options{
        Addr: "redis:6379",
    })

    _, err = redisClient.Ping(context.Background()).Result()
    if err != nil {
        log.Fatalf("Could not connect to Redis: %v", err)
    }

    fmt.Println("Successfully connected to Redis")
}
