package main

import ("
		github.com/gofiber/fiber/v2"
		"codebrains.io/todolist/database"
		"codebrains.io/todolist/models"
		"gorm.io/driver/postgres"
		"gorm.io/gorm"
	)

func helloWorld(c *fiber.Ctx) error{
	return c.SendString("Welcome")
}
func initDatabase(){
	
}
func main(){
	app:=fiber.New();
	app.Get("/",helloWorld);
	app.Listen(":8000");
}
