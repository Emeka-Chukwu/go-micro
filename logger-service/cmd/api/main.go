package main

import (
	"context"
	"fmt"
	"log"
	"log-service/data"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	webport  = "80"
	rpcPort  = "5001"
	mongoURL = "mongodb://mongo:27017"
	grpcPort = "5001"
)

var client *mongo.Client

type Config struct {
	Models data.Models
}

func main() {
	// connect o mongo

	mongoClient, err := connectMongo()
	if err != nil {
		log.Panic(err)
	}
/// create a context to disconnect

ctx, cancel := context.WithTimeout(context.Background(),15*time.Second)
defer cancel()

// close connection

defer func(){
	if err = client.Disconnect(); err != nil{
		panic(err)
	}
}()

app := Config{
	Models: data.New(client),
}
}


func (app *Config) serve(){
	srv := &http.Server{
		Addr: fmt.Sprintf(":%s", webport),
	}
}

func connectMongo() (*mongo.Client, error) {
	// create connection options
	clientOptions := options.Client().ApplyURI(mongoURL)
	clientOptions.SetAuth(options.Credential{
		Username: "admin",
		Password: "password",
	})

	// connect
	c, err : mongo.Connect(context.TODO(), clientOptions)
	if err!= nil{
		return nil,err
	}
	return c,nil
}
