package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Entry struct {
	ID       primitive.ObjectID `bson:"id"`
	Quantity *int               `bson:"quantity"`
	Dish     *string            `bson:"dish"`
	Calories *int               `bson:"calories"`
	Protein  *int               `bson:"protein"`
}
