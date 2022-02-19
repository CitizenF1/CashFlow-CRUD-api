package db

import (
	"context"
	"fmt"
	"log"

	"cashflow/controllers"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//DB Connect
func Connect() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalf("Error connecting to DB: %v \n", err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatalf("Error check DB connect: %v \n", err)
	}

	fmt.Println("connect to DB")

	db := client.Database("CRUD")

	controllers.PaymentsCollection(db)
}
