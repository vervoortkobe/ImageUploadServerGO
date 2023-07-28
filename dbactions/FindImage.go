package dbactions

import (
	"context"
	"fmt"
	"server/exports"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindImage(id string) (exports.Image, error) {
	coll := exports.MongoClient().Database("Server").Collection("images")

	filter := bson.D{{Key: "id", Value: id}}

	var result bson.M

	err := coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Printf("❌ | No image record found with id: %s\n", id)
			return exports.EmptyImage, err
		}
		return exports.EmptyImage, err
	}

	var timestamp32 int32 = result["timestamp"].(int32)

	image := exports.Image{
		Id:        result["id"].(string),
		Name:      result["name"].(string),
		Data:      result["data"].(string),
		Timestamp: int(timestamp32),
	}
	fmt.Printf("🔍 | Found image record with id: %s\n%v\n", id, result["name"])
	return image, nil
}
