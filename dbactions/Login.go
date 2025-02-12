package dbactions

import (
	"context"
	"fmt"
	"server/exports"
	"strings"
	"time"
)

func Login(username string, password string) (string, error) {
	coll := exports.MongoClient().Database("Server").Collection("logins")

	login := exports.UserCreds{
		Username: username,
		Password: password,
	}

	fmt.Printf("❗ | New user login request to DB: %s on %dts\n", login.Username, int(time.Now().Unix()))

	b, err := FindLogin(login.Username)
	if !b && err == nil {
		CheckPass(login.Username, login.Password)
		fmt.Printf("found login username: %s (duplicate username)! Returning error!\n", username)
		return "err_duplicate_username", err

	} else {
		if !b && err != nil {
			fmt.Printf("❌ | Failed login request, username: %s (mongodb findone error)! Returning error!\n", username)
			return "err_mongodb_findone", err
		}

		result, err := coll.InsertOne(context.TODO(), login)
		if err != nil {
			return "err_insert_one", err
		}

		objid := strings.Split(string(fmt.Sprintf("%v", result)), `&{ObjectID("`)[1]
		objid = objid[:len(objid)-3]
		fmt.Printf("✅ | Successfull login upload, username: %s (objid: %s)\n", login.Username, objid)
		return login.Username, nil
	}
}
