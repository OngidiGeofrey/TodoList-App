package models

import(
	"codebrains.io/todolist/database"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// define the toDo model 
type ToDo struct{
	ID uint `gorm:"primarykey" json:"id"`
	Title string `json:"title"`
	Completed bool `json:"completed"`

}
//define function to get toDos
func GetToDos(c *fiber.Ctx) error{
db:=database.DBConn
var todos []ToDo
db.Find(&todos);
return c.JSON(fiber.Map{"responseCode":1, "responseMessage":"Tasks Fetched Successfully", "data":&todos});
}
//define function to get toDos
func GetToDoById(c *fiber.Ctx) error{
	id:=c.Params("id")
	if id == "" {
		return c.Status(400).JSON(fiber.Map{"responseCode": 0, "responseMessage": "Id is required"})
	}
	db:=database.DBConn
	var todo ToDo
	err:=db.Find(&todo,id).Error
	if(err!=nil){
		return c.Status(500).JSON(fiber.Map{"responseCode":0, "responseMessage":"Something went Wrong"});
	 } 
	return c.JSON(fiber.Map{"responseCode":1, "responseMessage":"Task Fetched Successfully", "data":&todo});
	}

	// update task by Id
func UpdateToDoById(c *fiber.Ctx) error{
		id:=c.Params("id")
		// task structure
		type UpdateToDo struct{
			Title string `json:"title"`
			Completed bool `json:"completed"`
		}
		if id == "" {
			return c.Status(400).JSON(fiber.Map{"responseCode": 0, "responseMessage": "Id is required"})
		}
		db:=database.DBConn
		var todo ToDo

		// find task by id
		err:=db.Find(&todo,id).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return c.Status(404).JSON(fiber.Map{"responseCode": 0, "responseMessage": "Task not found"})
			}
			return c.Status(500).JSON(fiber.Map{"responseCode": 0, "responseMessage": "Something went wrong"})
		}
		 var updateToDo UpdateToDo 
		 // validate inputs
		err=c.BodyParser(&updateToDo)
		if(err!=nil){
			return c.Status(500).JSON(fiber.Map{"responseCode":0, "responseMessage":"Missing Parameters"});
		}
		// set new values
		todo.Title=updateToDo.Title
		todo.Completed=updateToDo.Completed

		//save changes
		db.Save(&todo)
		return c.JSON(fiber.Map{"responseCode":1, "responseMessage":"Task Updated Successfully", "data":&todo})
		}
//define function to create toDos
func CreateToDos(c *fiber.Ctx) error{
	db:=database.DBConn
	todo:=new(ToDo)
	err:=c.BodyParser(todo)
	if(err!=nil){
		return c.Status(500).JSON(fiber.Map{"responseCode":0, "responseMessage":"Missing Parameters"});
	}
	// Input validation
	if todo.Title == "" {
		return c.Status(400).JSON(fiber.Map{"responseCode": 0, "responseMessage": "Title is required"})
	}
	
	if len(todo.Title) < 3 {
		return c.Status(400).JSON(fiber.Map{"responseCode": 0, "responseMessage": "Title must be at least 3 characters long"})
	}
	
	 err=db.Create(&todo).Error
	 if(err!=nil){
		return c.Status(500).JSON(fiber.Map{"responseCode":0, "responseMessage":"Something went Wrong"});
	 } 
	 return c.JSON(fiber.Map{"responseCode":1, "responseMessage":"Task Created Successfully", "data":&todo});
}

func DeleteToDoById(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(400).JSON(fiber.Map{"responseCode": 0, "responseMessage": "Id is required"})
	}
	db := database.DBConn
	var todo ToDo
	err := db.First(&todo, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(404).JSON(fiber.Map{"responseCode": 0, "responseMessage": "Task not found"})
		}
		return c.Status(500).JSON(fiber.Map{"responseCode": 0, "responseMessage": "Something went wrong"})
	}

	err = db.Delete(&todo).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"responseCode": 0, "responseMessage": "Something went wrong"})
	}
	return c.Status(200).JSON(fiber.Map{"responseCode": 1, "responseMessage": "Task deleted successfully"})
}