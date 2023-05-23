package dbactions

import (
	"context"
	"fmt"
	"math/rand"
	"server/exports"
	"strings"
	"time"
)

func InsertOne(name string, data string, timestamp int) (string, error) {
	coll := exports.MongoClient().Database("Server").Collection("images")

	image := exports.Image{
		Id:        genId(15),
		Name:      name,
		Data:      data,
		Timestamp: timestamp,
	}

	fmt.Printf("❗ | New upload to DB: %s (%s) on %dts\n", image.Name, image.Id, image.Timestamp)

	img, e := FindOne(image.Id)
	if img != exports.EmptyImage && e == nil {
		fmt.Printf("❌ | Failed upload, id: %s (duplicate id)! Retrying until success...\n", image.Id)
		InsertOne(name, data, timestamp)
		var err error
		return "err_duplicate_id", err

	} else {

		result, err := coll.InsertOne(context.TODO(), image)
		if err != nil {
			return "err_insert_one", err
		}

		objid := strings.Split(string(fmt.Sprintf("%v", result)), `&{ObjectID("`)[1]
		objid = objid[:len(objid)-3]
		fmt.Printf("✅ | Successfull upload, id: %s (objid: %s)\n", image.Id, objid)
		return image.Id, nil
	}
}

var charset = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func genId(n int) string {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]rune, n)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
