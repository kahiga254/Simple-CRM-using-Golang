package main

import (
	"fmt"
	"log"

	"Simple-CRM-using-Golang/database" // ✅ Import the database package (local)
	"Simple-CRM-using-Golang/lead"    // ✅ Import the leads package (local)

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite" // ✅ SQLite driver for GORM
	"gorm.io/gorm"          // ✅ Correct GORM version
)


func setupRoutes( app *fiber.App) {
	app.Get("/api/v1/lead",lead.GetLeads)
	app.Get("/api/v1/lead/:id",lead.GetLead)
	app.Post("/api/v1/lead",lead.NewLead)
	app.Delete("/api/v1/lead/:id",lead.DeleteLead)
}

func initDatabase()(*gorm.DB, error){
	db, err := gorm.Open(sqlite.Open("./gorm.db"), &gorm.Config{})
	if err != nil{
		return nil, err
	}

	//Auto-migrate the lead model
	err = db.AutoMigrate(&lead.Lead{})
	if err != nil {
		return nil, err
	}
	fmt.Println("Database successfully connected and migrated")
	return db, nil
}
func main () {
	app := fiber.New()

	db, err := initDatabase()
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}
	database.DB = db // ✅ Store the database instance in the global variable
	initDatabase()
	setupRoutes(app)

	//start the fiber server
	port := ":3000"
	log.Println("Server running on port", port)
	if err := app.Listen(port); err != nil{
		log.Fatal("server failed to start:", err)
	}
	
	// Properly close the database connection when exiting
	sqlDB, err := database.DB.DB()
	if err == nil{
		defer sqlDB.Close()
	} else {
		log.Println("Warning: Failed to retrieve sql.DB for closing.")
	}

}