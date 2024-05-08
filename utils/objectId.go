package utils

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ConvertObjectIdToString(id interface{}) string {	// cast objectID to primitive.ObjectID
	primitive, _ := id.(primitive.ObjectID);
	return primitive.Hex()
}