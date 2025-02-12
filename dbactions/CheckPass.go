package dbactions

import (
	"context"
	"fmt"
	"server/exports"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CheckPass(username string, password string) (bool, error) {
	coll := exports.MongoClient().Database("Server").Collection("logins")

	filter := bson.D{{Key: "username", Value: username}}

	var result bson.M

	err := coll.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Printf("❌ | No login record found with username: %s\n", username)
			return true, nil
		}
		return false, err
	}
	if result["username"] == username && result["password"] == password {
		fmt.Printf("✅ | User %s (%s) successfully logged in!\n", username, result["id"])
		//create sess
	} else {
		fmt.Printf("❌ | User %s (%s) entered the wrong password!\n", username, result["id"])
	}
	return false, nil
}
