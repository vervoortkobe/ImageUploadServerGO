package dbactions

import (
	"context"
	"encoding/json"
	"fmt"
	"server/exports"

	"go.mongodb.org/mongo-driver/bson"
)

func GetAll() ([]exports.Image, error) {
	coll := exports.MongoClient().Database("Server").Collection("images")

	cursor, err := coll.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}

	var results []exports.Image
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}

	for _, result := range results {
		cursor.Decode(&result)
		output, err := json.MarshalIndent(result, "", "    ")
		if err != nil {
			return nil, err
		}
		_ = output
		fmt.Print(result)
	}
	return results, nil
}
