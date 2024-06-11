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

