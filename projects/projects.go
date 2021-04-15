package projects

import (
	"github.com/chaudharypraveen98/GoProfileAPI/database"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Projects struct {
	Projects []SingleProject `json:"projects"`
}

type SingleProject struct {
	gorm.Model
	ID                  int32          `json:"id"`
	Title               string         `json:"title"`
	Desciption          string         `json:"desc"`
	LastUpdated         string         `json:"last_updated"`
	ProgrammingLanguage string         `json:"programming_language"`
	Link                string         `json:"link"`
	Topic               []ProjectTopic `gorm:"foreignKey:ProjectTopic" json:"topics"`
}

type ProjectTopic struct {
	gorm.Model
	Name string
}

func GetProjects(c *fiber.Ctx) error {
	db := database.DBConn
	var projects []Projects
	db.Find(&projects)
	return c.JSON(projects)
}

// function name must be capital

func GetSingleProject(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var SingleProject Projects
	db.Find(&SingleProject, id)
	return c.JSON(SingleProject)
}
