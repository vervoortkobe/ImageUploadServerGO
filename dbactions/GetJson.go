package dbactions

import (
	"context"
	"encoding/json"
	"server/exports"

	"go.mongodb.org/mongo-driver/bson"
)

func GetJson() (string, error) {
	coll := exports.MongoClient().Database("Server").Collection("images")

	cursor, err := coll.Find(context.TODO(), bson.D{})
	if err != nil {
		return "err_not_found", err
	}
	defer cursor.Close(context.TODO())

	var results []bson.M

	for cursor.Next(context.Background()) {
		var result bson.M
		if err := cursor.Decode(&result); err != nil {
			return "err_decode_proc", err
		}
		results = append(results, result)
	}

	if err := cursor.Err(); err != nil {
		return "err_cursor_err", err
	}

	bufferString, err := json.Marshal(results)
	if err != nil {
		return "err_json_marshal", err
	}

	return string(bufferString), nil
}
