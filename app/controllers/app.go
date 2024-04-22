package controllers

import (
	"strconv"
	"warehouse-app/app/helpers"
	"warehouse-app/app/models"

	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

var (
	client    models.Client
	staff     models.Staff
	unit      models.Unit
	warehouse models.Warehouse
	errors    helpers.Error
)

func (c App) Index() revel.Result {
	return c.Render()
}

func (c *App) parseBodyClient(clients *[]models.Client) error {
	err := c.Params.BindJSON(clients)
	return err
}

func (c *App) parseBodyStaff(staffs *[]models.Staff) error {
	err := c.Params.BindJSON(staffs)
	return err
}

func (c *App) parseBodyUnit(units *[]models.Unit) error {
	err := c.Params.BindJSON(units)
	return err
}

func (c *App) parseBodyWarehouse(warehouses *[]models.Warehouse) error {
	err := c.Params.BindJSON(warehouses)
	return err
}

func (c App) Logout() revel.Result {
	return c.Redirect("/")
}

func (c App) GetClient(id int64) revel.Result {
	client, err := models.GetClient(id)
	if err != nil {
		c.Response.Status = 404 // or appropriate error status code
		return nil
	}
	c.Response.Status = 200 // Success status code
	return c.RenderJSON(client)
}

func (c App) AddClient() revel.Result {
	var errors map[string]interface{}
	var clients []models.Client

	// Parse JSON body
	err := c.parseBodyClient(&clients)
	if err != nil {
		c.Response.Status = 400
		errors = map[string]interface{}{"error": "Invalid JSON format"}
		return c.RenderJSON(errors)
	}

	// Add each client in the array
	for _, client := range clients {
		if err := client.AddClient(); err != nil {
			c.Response.Status = 503
			errors = map[string]interface{}{"error": "Unable to service your request. Please try again later"}
			return c.RenderJSON(errors)
		}
	}

	// Success
	c.Response.Status = 201
	return nil
}

func (c App) UpdateClient(id int64) revel.Result {
	var client models.Client
	if err := c.Params.BindJSON(&client); err != nil {
		c.Response.Status = 400 // Bad request status code
		return nil
	}
	client.ID = id
	if err := models.UpdateClient(&client); err != nil {
		c.Response.Status = 500 // Internal server error status code
		return nil
	}
	c.Response.Status = 200 // Success status code
	return c.RenderJSON(client)
}

func (c App) DeleteClient(id int64) revel.Result {
	if err := models.DeleteClient(id); err != nil {
		c.Response.Status = 500 // Internal server error status code
		return nil
	}
	c.Response.Status = 204 // No content status code
	return nil
}

func (c App) ListClients() revel.Result {
	clientModel := models.Client{}
	data, err := clientModel.ListClients()
	if err != nil {
		c.Response.Status = 204
		return nil
	}
	c.Response.Status = 200
	return c.RenderJSON(data)
}

func (c App) AddStaff() revel.Result {
	var errors map[string]interface{}
	var staffs []models.Staff

	// Parse JSON body
	err := c.parseBodyStaff(&staffs)
	if err != nil {
		c.Response.Status = 400
		errors = map[string]interface{}{"error": "Invalid JSON format"}
		return c.RenderJSON(errors)
	}

	// Add each staff member in the array
	for _, s := range staffs {
		if err := s.AddStaff(); err != nil {
			c.Response.Status = 503
			errors = map[string]interface{}{"error": "Unable to service your request. Please try again later"}
			return c.RenderJSON(errors)
		}
	}

	// Success
	c.Response.Status = 201
	return nil
}

func (c App) GetStaff(id int64) revel.Result {
	staff, err := models.GetStaff(id)
	if err != nil {
		c.Response.Status = 404 // Not found status code
		return nil
	}
	c.Response.Status = 200 // Success status code
	return c.RenderJSON(staff)
}

func (c App) UpdateStaff(id int64) revel.Result {
	var staff models.Staff
	if err := c.Params.BindJSON(&staff); err != nil {
		c.Response.Status = 400 // Bad request status code
		return nil
	}
	staff.ID = id
	if err := models.UpdateStaff(&staff); err != nil {
		c.Response.Status = 503 // Service unavailable status code
		errors.Error = "Unable to service your request. Please try again later"
		return c.RenderJSON(errors)
	}
	c.Response.Status = 200 // Success status code
	return nil
}

func (c App) DeleteStaff(id int64) revel.Result {
	if err := models.DeleteStaff(id); err != nil {
		c.Response.Status = 503 // Service unavailable status code
		errors.Error = "Unable to service your request. Please try again later"
		return c.RenderJSON(errors)
	}
	c.Response.Status = 204 // No content status code
	return nil
}

func (c App) ListStaffs() revel.Result {
	staffModel := models.Staff{}
	data, err := staffModel.ListStaffs()
	if err != nil {
		c.Response.Status = 204 // No content status code
		return nil
	}
	c.Response.Status = 200 // Success status code
	return c.RenderJSON(data)
}

func (c App) GetLoginStatus() revel.Result {
	var staff models.Staff
	if err := c.Params.BindJSON(&staff); err != nil {
		return c.RenderJSON(map[string]interface{}{"error": "Invalid JSON format"})
	}

	// Check login credentials and retrieve staff information
	loggedInStaff, err := staff.GetLoginStatus()
	if err != nil {
		return c.RenderJSON(map[string]interface{}{"error": err.Error()})
	}

	// Store only the staff ID in session
	c.Session["ID"] = strconv.Itoa(int(loggedInStaff.ID))
	c.Session.SetDefaultExpiration()

	return c.RenderJSON(map[string]interface{}{"ID": loggedInStaff.ID})
}

func (c App) AssignUnitsToClient(clientID int) revel.Result {
	// Parse JSON body to get units data
	var units []models.Unit
	if err := c.parseBodyUnit(&units); err != nil {
		c.Flash.Error("Invalid JSON format")
		return c.Redirect("/admin-dashboard")
	}

	// Assign each unit to the client
	for _, unit := range units {
		unit.ClientID = clientID // Set the client ID for the unit
		// Add unit to the database
		if err := unit.AddUnit(); err != nil {
			c.Flash.Error("Error assigning units to client")
			return c.Redirect("/admin-dashboard")
		}
	}

	// Units assigned successfully, show success message or redirect
	c.Flash.Success("Units assigned successfully")
	return c.Redirect("/admin-dashboard")
}

func (c App) AddUnit() revel.Result {
	var errors map[string]interface{}
	var units []models.Unit

	// Parse JSON body
	err := c.parseBodyUnit(&units)
	if err != nil {
		c.Response.Status = 400
		errors = map[string]interface{}{"error": "Invalid JSON format"}
		return c.RenderJSON(errors)
	}

	// Add each staff member in the array
	for _, s := range units {
		if err := s.AddUnit(); err != nil {
			c.Response.Status = 503
			errors = map[string]interface{}{"error": "Unable to service your request. Please try again later"}
			return c.RenderJSON(errors)
		}
	}

	// Success
	c.Response.Status = 201
	return nil
}

func (c App) GetUnit(id int64) revel.Result {
	unit, err := models.GetUnit(id)
	if err != nil {
		c.Response.Status = 404 // or appropriate error status code
		return nil
	}
	c.Response.Status = 200 // Success status code
	return c.RenderJSON(unit)
}

func (c App) UpdateUnit(id int64) revel.Result {
	var unit models.Unit
	if err := c.Params.BindJSON(&unit); err != nil {
		c.Response.Status = 400 // Bad request status code
		return nil
	}
	unit.ID = id
	if err := models.UpdateUnit(&unit); err != nil {
		c.Response.Status = 500 // Internal server error status code
		return nil
	}
	c.Response.Status = 200 // Success status code
	return c.RenderJSON(unit)
}

func (c App) DeleteUnit(id int64) revel.Result {
	if err := models.DeleteUnit(id); err != nil {
		c.Response.Status = 500 // Internal server error status code
		return nil
	}
	c.Response.Status = 204 // No content status code
	return nil
}

func (c App) ListUnits() revel.Result {
	unitModel := models.Unit{}
	data, err := unitModel.ListUnits()
	if err != nil {
		c.Response.Status = 204
		return nil
	}
	c.Response.Status = 200
	return c.RenderJSON(data)
}

// func (c App) GetClientUnits(clientID int) revel.Result {
// 	units, err := models.GetClientUnits(clientID)
// 	if err != nil {
// 		c.Response.Status = 500 // Internal Server Error
// 		return c.RenderJSON(map[string]interface{}{
// 			"error": "Failed to fetch client units",
// 		})
// 	}

// 	c.Response.Status = 200 // OK
// 	return c.RenderJSON(units)
// }

func (c App) GetTotalSales() revel.Result {
	unitModel := models.Unit{}
	data, err := unitModel.GetTotalSales()
	if err != nil {
		c.Response.Status = 204
		return nil
	}
	c.Response.Status = 200
	return c.RenderJSON(data)
}

func (c App) GetOrderCompleted() revel.Result {
	unitModel := models.Unit{}
	data, err := unitModel.GetOrderCompleted()
	if err != nil {
		c.Response.Status = 204
		return nil
	}
	c.Response.Status = 200
	return c.RenderJSON(data)
}

func (c App) GetOrderPending() revel.Result {
	unitModel := models.Unit{}
	data, err := unitModel.GetOrderPending()
	if err != nil {
		c.Response.Status = 204
		return nil
	}
	c.Response.Status = 200
	return c.RenderJSON(data)
}

func (c App) GetLatestOrders() revel.Result {
	unitModel := models.Unit{}
	data, err := unitModel.GetLatestOrders()
	if err != nil {
		c.Response.Status = 204
		return nil
	}
	c.Response.Status = 200
	return c.RenderJSON(data)
}

func (c App) AddWarehouse() revel.Result {
	var errors map[string]interface{}
	var warehouses []models.Warehouse

	// Parse JSON body
	err := c.parseBodyWarehouse(&warehouses)
	if err != nil {
		c.Response.Status = 400
		errors = map[string]interface{}{"error": "Invalid JSON format"}
		return c.RenderJSON(errors)
	}

	// Add each staff member in the array
	for _, s := range warehouses {
		if err := s.AddWarehouse(); err != nil {
			c.Response.Status = 503
			errors = map[string]interface{}{"error": "Unable to service your request. Please try again later"}
			return c.RenderJSON(errors)
		}
	}

	// Success
	c.Response.Status = 201
	return nil
}

func (c App) GetWarehouse(id int64) revel.Result {
	warehouse, err := models.GetWarehouse(id)
	if err != nil {
		c.Response.Status = 404
		return nil
	}
	c.Response.Status = 200 // Success status code
	return c.RenderJSON(warehouse)
}

func (c App) UpdateWarehouse(id int64) revel.Result {
	var warehouse models.Warehouse
	if err := c.Params.BindJSON(&warehouse); err != nil {
		c.Response.Status = 400 // Bad request status code
		return nil
	}
	warehouse.ID = id
	if err := models.UpdateWarehouse(&warehouse); err != nil {
		c.Response.Status = 500 // Internal server error status code
		return nil
	}
	c.Response.Status = 200 // Success status code
	return c.RenderJSON(client)
}

func (c App) DeleteWarehouse(id int64) revel.Result {
	if err := models.DeleteWarehouse(id); err != nil {
		c.Response.Status = 500 // Internal server error status code
		return nil
	}
	c.Response.Status = 204 // No content status code
	return nil
}

func (c App) ListWarehouses(page int64, perPage int64) revel.Result {
	warehouseModel := models.Warehouse{}
	data, err := warehouseModel.ListWarehouses()
	if err != nil {
		c.Response.Status = 204
		return nil
	}
	c.Response.Status = 200
	return c.RenderJSON(data)
}
