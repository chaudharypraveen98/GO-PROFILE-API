package projects

import (
	"github.com/gofiber/fiber/v2"
)

func getProjects(c *fiber.Ctx) error {
	return c.SendString("all projects")
}

func getSingleProject(c *fiber.Ctx) error {
	return c.SendString("single projects")
}
