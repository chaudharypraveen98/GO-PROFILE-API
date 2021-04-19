package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/chaudharypraveen98/GoProfileAPI/database"
	"github.com/chaudharypraveen98/GoProfileAPI/projects"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func loadData() {
	// initiate the database connection
	db := database.DBConn

	// Open our jsonFile
	jsonFile, err := os.Open("data/projects.json")

	// if we os.Open returns an error then handle while opening
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully Opened Projects.json")
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var SerialProjectData projects.SerialProject
	json.Unmarshal(byteValue, &SerialProjectData)

	// delete tables if they exit already
	db.Migrator().DropTable(&projects.Projects{})
	db.Migrator().DropTable(&projects.SingleProject{})

	// create table using struct
	db.Migrator().CreateTable(&projects.Projects{})
	db.Migrator().CreateTable(&projects.SingleProject{})

	var allProjects projects.Projects
	json.Unmarshal(byteValue, &allProjects)

	// creating collections of all projects
	db.Create(&allProjects)

	// creating single project
	for _, project := range SerialProjectData.Projects {
		var singleProject projects.SingleProject
		singleProject.ID = project.ID
		singleProject.Title = project.Title
		singleProject.Desciption = project.Desciption
		singleProject.Forks = project.Forks
		singleProject.LastUpdated = project.LastUpdated
		singleProject.Link = project.Link
		singleProject.ProgrammingLanguage = project.ProgrammingLanguage
		db.Create(&singleProject)
	}
	defer jsonFile.Close()
}

// This func handles database connection
func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open(sqlite.Open("projects.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Database successfully opened")
}

// router to handles url endpoints
func setRoutes(app *fiber.App) {
	app.Get("/api/v1/projects", projects.GetProjects)
	app.Get("/api/v1/projects/:id", projects.GetSingleProject)
	app.Post("/api/v1/projects/:id/update", projects.UpdateSingleProject)
	app.Post("/api/v1/projects/create", projects.CreateSingleProject)
	app.Post("/api/v1/projects/:id/delete", projects.DeleteProject)
}

func main() {
	app := fiber.New()
	initDatabase()
	loadData()
	database.DBConn.AutoMigrate(&projects.SingleProject{}, &projects.Projects{})
	fmt.Println("Database migrated successfully")
	fmt.Println("---------------------------")
	setRoutes(app)
	app.Listen(":3000")
}
