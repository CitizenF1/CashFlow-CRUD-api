package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Payments struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name     string             `bson:"name" json:"name"`
	Price    int                `bson:"price" json:"price"`
	Date     time.Time          `bson:"date" json:"date"`
	Type     string             `bson:"type" json:"type"`
	Comment  string             `bson:"comment" json:"comment"`
	Category string             `bson:"category" json:"category"`
}

var collection *mongo.Collection

func PaymentsCollection(c *mongo.Database) {
	collection = c.Collection("paymets")
}

func parseID(c *gin.Context) (primitive.ObjectID, error) {
	id := c.Param("id")
	parsedID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "The paymets ID is not valid",
		})
		return parsedID, err
	}
	return parsedID, nil
}

func ListPayments(c *gin.Context) {
	var payments []*Payments
	cursor, err := collection.Find(context.TODO(), bson.M{}, options.Find())
	if err != nil {
		fmt.Printf("Error getting payments: %s \n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	for cursor.Next(context.TODO()) {
		var pay Payments
		err := cursor.Decode(&pay)
		if err != nil {
			fmt.Printf("Failed to decode: %s \n", err)
		}

		payments = append(payments, &pay)
	}

	if err := cursor.Err(); err != nil {
		fmt.Printf("Cursor error: %s \n", err)
	}

	cursor.Close(context.TODO())
	c.JSON(http.StatusOK, gin.H{
		"data": payments,
	})
}

func CreatePayment(c *gin.Context) {
	var pay Payments
	c.BindJSON(&pay)
	name, price, cat, typ, comm := pay.Name, pay.Price, pay.Category, pay.Type, pay.Comment //, pay.Type
	newPay := Payments{
		Name:     name,
		Price:    price,
		Date:     time.Time{},
		Category: cat,
		Comment:  comm,
		Type:     typ,
	}

	createdPayment, err := collection.InsertOne(context.TODO(), newPay)

	if err != nil {
		fmt.Printf("Failed to create payment: %v \n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	data := map[string]interface{}{
		"_id":      createdPayment.InsertedID,
		"name":     name,
		"price":    price,
		"category": cat,
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": http.StatusCreated,
		"data":   data,
	})
}

func UpdatePayment(c *gin.Context) {
	id, err := parseID(c)
	if err != nil {
		return
	}

	var pay Payments
	c.BindJSON(&pay)
	updatePayment := bson.M{}
	if pay.Name != "" {
		updatePayment["name"] = pay.Name
	}
	if pay.Price != 0 {
		updatePayment["price"] = pay.Price
	}
	if pay.Name != "" {
		updatePayment["category"] = pay.Category
	}

	filter := bson.M{"_id": id}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	update := bson.D{{Key: "$set", Value: updatePayment}}

	var updatedDocument bson.M
	e := collection.FindOneAndUpdate(context.TODO(), filter, update, opts).Decode(&updatedDocument)
	if e != nil {
		// ErrNoDocuments means that the filter did not match any documents in the collection
		if e == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{
				"status":  http.StatusNotFound,
				"message": "Could not find payment",
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   updatedDocument,
	})
}

func DeletePayments(c *gin.Context) {
	id, err := parseID(c)
	if err != nil {
		return
	}

	filter := bson.M{"_id": id}
	opts := options.FindOneAndDelete()
	var deletedDocument bson.M
	e := collection.FindOneAndDelete(context.TODO(), filter, opts).Decode(&deletedDocument)
	if e != nil {
		// ErrNoDocuments means that the filter did not match any documents in the collection
		if e == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{
				"status":  http.StatusNotFound,
				"message": "Could not find user",
			})
			return
		}
	}

	c.JSON(http.StatusNoContent, nil)
}
