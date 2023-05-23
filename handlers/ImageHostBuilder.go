package handlers

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"log"
	"server/exports"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func ImageHostBuilder(c *fiber.Ctx) error {
	if len(c.Params("value")) != 15 {
		return nil
	}

	img, err := ImageHost(string(c.Params("value")))
	if img != exports.EmptyImage && err == nil {
		imageData, err := base64.StdEncoding.DecodeString(img.Data)
		if err != nil {
			return err
		}

		buf := new(bytes.Buffer)

		imgImage, err := exports.DecBufToImg(imageData)
		if err != nil {
			log.Fatal(err)
		}

		/*width := imgImage.Bounds().Dx()
		height := imgImage.Bounds().Dy()
		fmt.Printf("Image dimensions: %d x %d\n", width, height)*/

		if strings.Contains(img.Name, ".png") {
			err = exports.EncPngImgToBuf(imgImage, buf)
			if err != nil {
				log.Fatal(err)
			}

			c.Set("Content-Type", "image/png")
		} else if strings.Contains(img.Name, ".jpg") {
			err = exports.EncJpgImgToBuf(imgImage, buf)
			if err != nil {
				log.Fatal(err)
			}

			c.Set("Content-Type", "image/jpeg")
		}
		c.Set("Content-Length", fmt.Sprint(buf.Len()))

		return c.Send(buf.Bytes())
	}
	return nil
}
