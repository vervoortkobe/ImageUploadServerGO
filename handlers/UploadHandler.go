package handlers

import (
	"encoding/base64"
	"fmt"
	"io"
	"server/dbactions"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

func UploadHandler(c *fiber.Ctx) error {
	file, err := c.FormFile("fileUpload")
	if err != nil {
		return c.SendString(fmt.Sprint(err))
	}

	src, err := file.Open()
	if err != nil {
		return c.SendString(fmt.Sprint(err))
	}
	defer src.Close()

	data, err := io.ReadAll(src)
	if err != nil {
		return c.SendString(fmt.Sprint(err))
	}

	dataString := string(base64.StdEncoding.EncodeToString(data))

	if file.Filename != "" {
		if file.Header["Content-Type"][0] == "image/jpeg" {
			if !strings.Contains(file.Filename, ".jpeg") && !strings.Contains(file.Filename, ".jpg") {
				file.Filename = file.Filename + ".jpg"
			}
		} else if file.Header["Content-Type"][0] == "image/png" {
			if !strings.Contains(file.Filename, ".png") {
				file.Filename = file.Filename + ".png"
			}
		} else if file.Header["Content-Type"][0] == "application/octet-stream" {
			if !strings.Contains(file.Filename, ".jpeg") &&
				!strings.Contains(file.Filename, ".jpg") &&
				!strings.Contains(file.Filename, ".png") {
				file.Filename = file.Filename + ".jpg"
			}
		} else {
			file.Filename = file.Filename + ".jpg"
		}
	}

	if file.Filename == "" ||
		file.Header["Content-Type"][0] != "image/jpeg" &&
			file.Header["Content-Type"][0] != "image/png" &&
			file.Header["Content-Type"][0] != "application/octet-stream" ||
		!strings.Contains(file.Filename, ".png") &&
			!strings.Contains(file.Filename, ".jpg") &&
			!strings.Contains(file.Filename, ".jpeg") {

		errMsg := ""

		switch true {
		case file.Filename == "":
			errMsg = fmt.Sprintf("❌ | File %s (%db) was not uploaded, because it has an invalid filename!\n", file.Filename, file.Size)

		case file.Header["Content-Type"][0] != "image/jpeg" &&
			file.Header["Content-Type"][0] != "image/png":
			errMsg = fmt.Sprintf("❌ | File %s (%db) was not uploaded, because it is of the wrong type!\n", file.Filename, file.Size)

		case !strings.Contains(file.Filename, ".png") &&
			!strings.Contains(file.Filename, ".jpg") &&
			!strings.Contains(file.Filename, ".jpeg"):
			errMsg = fmt.Sprintf("❌ | File %s (%db) was not uploaded, because it is of the wrong format!\n", file.Filename, file.Size)
		}

		fmt.Print(errMsg)
		return c.SendString(errMsg)
	}

	if file.Size >= 0 || file.Size <= 5000000 { //5mb max limit
		fmt.Printf("✅ | File %s (%db) is being uploaded...\n", file.Filename, file.Size)
		c.SendString(fmt.Sprintf("✅ | File %s (%db) is being uploaded...\n", file.Filename, file.Size))

		imgId, e := dbactions.InsertOne(file.Filename, strings.Split(dataString, " ")[0], int(time.Now().Unix()))

		if imgId != "err_duplicate_id" && e == nil ||
			imgId != "err_insert_one" && e == nil {

			return c.Redirect(fmt.Sprintf("/%s", imgId))
		} else {
			return nil
		}
	} else {
		fmt.Printf("❌ | File %s (%db) was not uploaded, because it is too large!\n", file.Filename, file.Size)
		return c.SendString(fmt.Sprintf("❌ | File %s (%db) was not uploaded, because it is too large!\n", file.Filename, file.Size))
	}
}
