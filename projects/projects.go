package projects

import (
	"errors"
	"fmt"

	"github.com/chaudharypraveen98/GoProfileAPI/database"
	"github.com/gofiber/fiber/v2"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Projects struct {
	gorm.Model
	Projects datatypes.JSON `json:"projects" gorm:"foreignKey:ID"`
}

// all projects serializer
type SerialProject struct {
	Projects []SingleProject `json:"projects" gorm:"foreignKey:ProjectID"`
}

type SingleProject struct {
	gorm.Model
	ProjectID           int32  `gorm:"primaryKey" json:"id"`
	Title               string `json:"title"`
	Desciption          string `json:"desc"`
	ProgrammingLanguage string `json:"programming_language"`
	Stars               int32  `gorm:"default:null" json:"stars"`
	Forks               int32  `gorm:"default:null" json:"forks"`
	LastUpdated         string `json:"last_updated"`
	Link                string `json:"link"`
}

// Function to return all Projects
func GetProjects(c *fiber.Ctx) error {
	db := database.DBConn
	var projects []Projects
	err := db.Find(&projects).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.SendStatus(404)
	}
	return c.JSON(projects)
}

// It will create a project with incremented id.
func CreateSingleProject(c *fiber.Ctx) error {
	db := database.DBConn
	inputData := SingleProject{}
	if err := c.BodyParser(&inputData); err != nil {
		fmt.Println("error = ", err)
		return c.SendStatus(500)
	}
	lastProject := SingleProject{}
	db.Last(&lastProject)
	inputData.ID = lastProject.ID + 1
	db.Create(&inputData)
	return c.SendStatus(201)
}

func UpdateSingleProject(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var SingleProject SingleProject
	db.First(&SingleProject, id)
	inputData := SingleProject
	if err := c.BodyParser(&inputData); err != nil {
		fmt.Println("error = ", err)
		return c.SendStatus(200)
	}
	fmt.Println("--------")
	fmt.Println(inputData)
	db.Save(inputData)
	return c.SendStatus(204)
}

func DeleteProject(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var SingleProject SingleProject
	err := db.Delete(&SingleProject, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.SendStatus(404)
	}
	return c.SendStatus(202)
}

func GetSingleProject(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var SingleProject SingleProject
	err := db.First(&SingleProject, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.SendStatus(404)
	}
	return c.JSON(SingleProject)
}
