package helpers

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

func Primitive(id string) *primitive.ObjectID {
	primitiveID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Printf("Error while string to primitve %s", err.Error())
		return nil
	}
	return &primitiveID
}
