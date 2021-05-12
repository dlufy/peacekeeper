package database

import (
	"context"
	"fmt"
	"time"

	"github.com/dlufy/peacekeeper/admin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var Client *mongo.Client

func Connect() (err error) {
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb+srv://gatekeeper:%s@cluster0.yjxfz.mongodb.net/%s?retryWrites=true&w=majority", admin.GetConfig().Database.Password, admin.GetConfig().Database.DBName))
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// Connect to MongoDB
	Client, err = mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		fmt.Println(err)
		return err
	}

	// Check the connection
	err = Client.Ping(ctx, readpref.Primary())

	if err != nil {
		fmt.Println(err)
	}
	return err
}

func GetClient() *mongo.Client {
	return Client
}

func CloseConnection() {
	Client.Disconnect(nil)
}
