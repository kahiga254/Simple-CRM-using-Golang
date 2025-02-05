package lead

import (
	"Simple-CRM-using-Golang/database"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Lead struct {
	gorm.Model
    Name    string     `json:"name"`
	Company	string	`json:"company"`
	Email	string	`json:"email"`
	Phone	string	`json:"phone"`

}
// Get all leads
func GetLeads(c *fiber.Ctx )error{
	var leads [] Lead
	database.DB.Find(&leads)
	return c.JSON(leads)

}
// Get a single lead by ID
func GetLead (c *fiber.Ctx)error{
	id := c.Params("id")
	var lead Lead
	database.DB.Find(&lead, id)
	if lead.Name == "" {
		return c.Status(404).SendString("lead not found")
	}
	return c.JSON(lead)
}

// Create a new lead
func NewLead(c *fiber.Ctx) error {
	lead := new(Lead)
	if err := c.BodyParser(lead); err != nil {
		return	c.Status(503).SendString(err.Error())
	}
	database.DB.Create(&lead)
	return c.JSON(lead)
}

// Delete a lead
func DeleteLead(c *fiber.Ctx)error{
	id := c.Params ("id")
	var lead Lead

	database.DB.First(&lead, id)
	if lead.Name == ""{
	return	c.Status(500).SendString("No lead found with ID")
	}
	database.DB.Delete(&lead)
	return c.SendString("lead successfully deleted")
}