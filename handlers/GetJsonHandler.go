package handlers

import (
	"encoding/json"
	"fmt"
	"server/dbactions"

	"github.com/gofiber/fiber/v2"
)

func GetJsonHandler(c *fiber.Ctx) error {
	jsonString, err := dbactions.GetJson()

	if err != nil {
		fmt.Print(jsonString, err)
		return c.SendString(fmt.Sprintf("❌ | Error while fetching all records: %s, %e", jsonString, err))
	} else {
		buffer := []byte("{\"json\": " + jsonString + "}")
		var bufferSingleMap map[string]interface{}
		json.Unmarshal(buffer, &bufferSingleMap)
		jsonData := bufferSingleMap
		return c.JSON(jsonData)
	}
}
