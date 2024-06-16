package models

import(
	"codebrains.io/todolist/database"
	"github.com/gofiber/fiber/v2"
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
//define function to create toDos
func CreateToDos(c *fiber.Ctx) error{
	db:=database.DBConn
	todo:=new(ToDo)
	err:=c.BodyParser(todo)
	if(err!=nil){
		return c.Status(500).JSON(fiber.Map{"responseCode":0, "responseMessage":"Missing Parameters"});
	}
	 err=db.Create(&todo).Error
	 if(err!=nil){
		return c.Status(500).JSON(fiber.Map{"responseCode":0, "responseMessage":"Something went Wrong"});
	 } 
	return c.JSON(&todo)
}