package projects

import (
	"github.com/chaudharypraveen98/GoProfileAPI/database"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type SingleProject struct {
	gorm.Model
	Id                  int32    `json:"id"`
	Title               string   `json:"title"`
	Desciption          string   `json:"desc"`
	LastUpdated         string   `json:"last_updated"`
	Topics              []string `json:"topics"`
	ProgrammingLanguage string   `json:"programming_language"`
	Link                string   `json:"link"`
}
type Projects struct {
	Projects []SingleProject `json:"projects"`
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
