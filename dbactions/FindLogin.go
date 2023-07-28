package dbactions

import (
	"context"
	"fmt"
	"server/exports"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindLogin(username string) (bool, error) {
	coll := exports.MongoClient().Database("Server").Collection("logins")

	filter := bson.D{{Key: "username", Value: username}}

	var result bson.M

	err := coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Printf("❌ | No login record found with id: %s\n", username)
			return false, err
		}
		return false, err
	}
	fmt.Printf("🔍 | Found login record with username: %s (%s)\n", username, result["id"])
	return true, nil
}
