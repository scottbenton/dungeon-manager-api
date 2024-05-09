package utils

import (
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ConvertStringToObjectId(id string) primitive.ObjectID {	// cast objectID to primitive.ObjectID
	objectId, err := primitive.ObjectIDFromHex(id);
	if(err != nil) {
		log.Println("Failed to convert string to objectID", err);
	}
	log.Println("Converted objectID: ", objectId);
	return objectId
}

func ConvertObjectIdToString(id interface{}) string {	// cast objectID to primitive.ObjectID
	primitive, _ := id.(primitive.ObjectID);
	return primitive.Hex()
}