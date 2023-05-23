package dbactions

import (
	"context"
	"fmt"
	"server/exports"

	"go.mongodb.org/mongo-driver/bson"
)

func LogAll() error {
	coll := exports.MongoClient().Database("Server").Collection("images")

	cursor, err := coll.Find(context.TODO(), bson.D{})
	if err != nil {
		return err
	}

	var results []exports.Image
	if err = cursor.All(context.TODO(), &results); err != nil {
		return err
	}

	for i, result := range results {
		cursor.Decode(&result)
		fmt.Println(i, result.Id)
	}
	return nil
}
