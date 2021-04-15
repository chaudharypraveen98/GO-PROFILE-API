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
	Topic               []ProjectTopic `gorm:"foreignKey:NameID" json:"topics"`
	ProgrammingLanguage string         `json:"programming_language"`
	Stars               int32          `json:"stars"`
	Forks               int32          `json:"forks"`
	LastUpdated         string         `json:"last_updated"`
	Link                string         `json:"link"`
}

type ProjectTopic struct {
	gorm.Model
	NameID string
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
