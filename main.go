package main

import (
	"fmt"
	"os"

	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

type Todo struct {
	ID        int    `json:"id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

func main() {
	fmt.Println("Hello World")

	app := fiber.New()
	 err := godotenv.Load(".env")
	 if err!= nil {
		log.Fatal("Error while Loading .env file ")
	 }
	 PORT :=  os.Getenv("PORT")

	todos := []Todo{}

	app.Get("/api/todos", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(todos)
	})

	//create a Todo

	app.Post("/api/todos", func(c *fiber.Ctx) error {
		todo := &Todo{} //{id:0,completed:false, body:nil}
		if err := c.BodyParser(todo); err != nil {
			return err
		}
		if todo.Body == "" {
			return c.Status(400).JSON(fiber.Map{"error": "Todo  body required"})
		}
		todo.ID = len(todos)
		todos = append(todos, *todo)
		return c.Status(201).JSON(todo)

	})
	//Update a Todo
	app.Patch("/api/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		for i, todo := range todos {

			if fmt.Sprint(todo.ID) == id {
				todos[i].Completed = true
				return c.Status(200).JSON(todos[i])
			}

		}
		return c.Status(404).JSON(fiber.Map{"error":"Todo not found"})
	})

	//Delete a Todo

	app.Delete("/api/todos/:id",func(c *fiber.Ctx) error {
		id := c.Params("id")

		for i, todo := range todos{
			if fmt.Sprint(todo.ID)==id{
				todos=append(todos[:i],todos[i+1:]...)
			return c.Status(200).JSON(fiber.Map{"success":true})
			}
		}
		return  c.Status(404).JSON(fiber.Map{"error":"Todo not found"})
	})



	log.Fatal(app.Listen(":"+PORT))

}
