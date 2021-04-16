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
	db := database.DBConn
	// Open our jsonFile
	jsonFile, err := os.Open("data/projects.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened Projects.json")
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var allProjects projects.Projects
	json.Unmarshal(byteValue, &allProjects)
	// db.Migrator().DropTable(&projects.Projects{})
	// db.Migrator().DropTable(&projects.SingleProject{})
	// db.Migrator().DropTable(&projects.ProjectTopic{})
	db.Migrator().CreateTable(&projects.Projects{})
	db.Migrator().CreateTable(&projects.SingleProject{})
	db.Migrator().CreateTable(&projects.ProjectTopic{})
	fmt.Println(len(allProjects.Projects))
	// fmt.Println(db.Migrator().HasTable(&projects.Projects{}))
	db.Create(&allProjects.Projects)
	// for i := 0; i < len(allProjects.Projects); i++ {
	// 	var SingleProject projects.SingleProject
	// 	SingleProject.ID = allProjects.Projects[i].ID
	// 	SingleProject.Title = allProjects.Projects[i].Title
	// 	SingleProject.Desciption = allProjects.Projects[i].Desciption
	// 	SingleProject.LastUpdated = allProjects.Projects[i].LastUpdated
	// 	SingleProject.ProgrammingLanguage = allProjects.Projects[i].ProgrammingLanguage
	// 	SingleProject.Link = allProjects.Projects[i].Link
	// 	fmt.Println(SingleProject.ProgrammingLanguage)
	// 	db.Create(&SingleProject)
	// }
	for _, project := range allProjects.Projects {
		db.Create(&project)
	}
	defer jsonFile.Close()
}
func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open(sqlite.Open("projects.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Database successfully opened")
}
func setRoutes(app *fiber.App) {
	app.Get("/api/v1/projects", projects.GetProjects)
	app.Get("/api/v1/projects/:id", projects.GetSingleProject)
}

func main() {
	app := fiber.New()
	initDatabase()
	loadData()
	database.DBConn.AutoMigrate(&projects.Projects{}, &projects.SingleProject{}, &projects.ProjectTopic{})
	fmt.Println("Database migrated successfully")
	setRoutes(app)
	app.Listen(":3000")
}
