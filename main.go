package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"codebrains.io/todolist/database"
	"codebrains.io/todolist/models"
	"gorm.io/gorm"
	)

func welcomeMessage(c *fiber.Ctx) error{
	return c.SendString("Welcome 😊")

}
func initDatabase() {
	var err error
	dsn := "root:@tcp(localhost:3306)/todoapp?charset=utf8mb4&parseTime=True&loc=Local"
	database.DBConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to db")
	}
	fmt.Println("Database connected!")
	database.DBConn.AutoMigrate(&models.ToDo{})
	fmt.Println("Migrated DB!")
}
func setupRoutes(app *fiber.App){
	app.Post("api/todos/get-all",models.GetToDos);
	app.Post("api/todos/create",models.CreateToDos);
	app.Post("api/todos/get/:id",models.GetToDoById);
	app.Post("api/todos/update/:id",models.UpdateToDoById);
	app.Post("api/todos/delete/:id",models.DeleteToDoById);
}
func main(){
	app:=fiber.New();
	initDatabase()
	app.Get("/",welcomeMessage);
	setupRoutes(app)
	app.Listen(":8000");
}
