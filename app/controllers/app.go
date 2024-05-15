package controllers

import (
	"strconv"
	"warehouse-app/app/models"

	"github.com/revel/revel"
)

var (
	client models.Client
	staff  models.Staff
	unit   models.Unit
)

type LoginPage struct {
	*revel.Controller
}

func (c LoginPage) Index() revel.Result {
	return c.Render()
}

func (c LoginPage) UserLogin() revel.Result {
	var loginCredentials models.Staff

	// Validate the JSON format and bind it to the loginCredentials struct
	if err := c.Params.BindJSON(&loginCredentials); err != nil {
		return c.RenderJSON(map[string]interface{}{"error": "Invalid JSON format"})
	}

	// Check if the name and password fields are not empty
	if loginCredentials.Name == "" || loginCredentials.Password == "" {
		return c.RenderJSON(map[string]interface{}{"error": "Name and password are required"})
	}

	// Validate login credentials and retrieve staff information along with user type
	loggedInStaff, userType, err := models.ValidateLogin(loginCredentials.Name, loginCredentials.Password)
	if err != nil {
		return c.RenderJSON(map[string]interface{}{"error": "Invalid credentials"})
	}

	// Log the logged-in staff information
	revel.AppLog.Info("Logged in Staff ID:", loggedInStaff.ID, "Name:", loggedInStaff.Name)

	// Check if the logged-in user is an owner or staff
	if userType == "Owner" {
		// Store the owner ID in session for owners
		c.Session["OwnerID"] = strconv.Itoa(int(loggedInStaff.ID))
		// Return success response for owners
		return c.RenderJSON(map[string]interface{}{
			"message":   "Login successful as Owner",
			"user_id":   loggedInStaff.ID,
			"user_type": "Owner",
		})
	} else if userType == "Staff" {
		// Store the staff ID in session for staff
		c.Session["StaffID"] = strconv.Itoa(int(loggedInStaff.ID))
		// Return success response for staff
		return c.RenderJSON(map[string]interface{}{
			"message":   "Login successful as Staff",
			"user_id":   loggedInStaff.ID,
			"user_type": "Staff",
		})
	}

	// If the user type is neither "Owner" nor "Staff", return an error response
	return c.RenderJSON(map[string]interface{}{"error": "Invalid user type"})
}
