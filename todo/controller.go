package todo

import (
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GetAll(c *fiber.Ctx) error {
	todos, err := findAll()
	if err != nil {
		return c.SendStatus(500)
	}

	return c.JSON(todos)
}

func GetOne(c *fiber.Ctx) error {
	id := c.Params("id")

	todo, err := findOne(id)
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return c.SendStatus(404)
		}

		return c.SendStatus(500)
	}

	return c.JSON(todo)
}

func Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	err := delete(id)
	if err != nil {
		return c.SendStatus(500)
	}

	return c.SendStatus(204)
}

func Post(c *fiber.Ctx) error {
	id := c.Params("id")

	todo, err := validate(c)
	if err != nil {
		return c.SendStatus(400)
	}

	if (id == "" && todo.Id != "") {
		return c.Status(400).SendString("Ids must be equal if present")
	}

	if (todo.Id != "" && id != todo.Id) {
		return c.Status(400).SendString("Ids must be equal if present")
	}

	todo.Id = id
	todo, err = upsert(todo)
	if err != nil {
		return c.SendStatus(500)
	}

	return c.JSON(todo)
}
