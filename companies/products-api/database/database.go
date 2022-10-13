package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
  "go.mongodb.org/mongo-driver/mongo/readpref"
	log "github.com/sirupsen/logrus"
)

var instance *mongo.Client

func close(client *mongo.Client, ctx context.Context, cancel context.CancelFunc){
	defer cancel()
	
	defer func(){
		if err := client.Disconnect(ctx); err != nil{
			panic(err)
		}
	}()
}

func connect(uri string)(*mongo.Client, context.Context, context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	return client, ctx, cancel, err
}

func ping(client *mongo.Client, ctx context.Context) error{ 
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
			return err
	}
	log.Info("Mongo connected successfully")
	return nil
}

func StartDatabase(uri string) *mongo.Client {
	client, ctx, cancel, err := connect(uri)
	if err != nil	{
			panic(err)
	}

	defer close(client, ctx, cancel)

	instance = client

	err = ping(client, ctx)
	if err != nil	{
		panic(err)
	}

	return instance
}

func GetInstance() *mongo.Client {
	if instance == nil {
		panic("Database not initialized")
	}
	return instance
}