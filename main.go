package main

import (
	"context"
	"gcp-go/db"
	"gcp-go/todo"
	"gcp-go/util"

	"github.com/gofiber/fiber/v2"
)

func main() {
    ctx := context.Background()
	app := fiber.New()

    db.Connect(ctx)
    util.CreateValidator()

    todoGroup := app.Group("/todos")
    todoGroup.Get("/", todo.GetAll)
    todoGroup.Get("/:id", todo.GetOne)
    todoGroup.Post("/:id?", todo.Post)
    todoGroup.Delete("/:id", todo.Delete)

    app.Listen("localhost:8080")

    defer db.Client.Close()
}
