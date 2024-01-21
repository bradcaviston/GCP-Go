package todo

import (
	"gcp-go/util"

	"github.com/gofiber/fiber/v2"
)

func validate(c *fiber.Ctx) (Todo, error) {
	var todo Todo
	err := c.BodyParser(&todo)

	err = util.Validate.Struct(todo)

	return todo, err
}
