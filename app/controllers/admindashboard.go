package controllers

import (
	"fmt"
	"strconv"
	"warehouse-app/app/helpers"
	"warehouse-app/app/models"

	"github.com/revel/revel"
)

var (
	errors helpers.Error
)

type AdminDashboard struct {
	*revel.Controller
}

func (c AdminDashboard) Index() revel.Result {
	// Retrieve the staff ID from the session
	staffIDInterface, err := c.Session.Get("StaffID")
	if err != nil {
		// Handle the case where StaffID is not found in session
		return c.RenderJSON(map[string]interface{}{"error": "StaffID not found in session"})
	}

	staffIDStr, ok := staffIDInterface.(string)
	if !ok {
		// Handle the case where StaffID is not a string in session
		return c.RenderJSON(map[string]interface{}{"error": "Invalid StaffID format in session"})
	}

	staffID, err := strconv.ParseInt(staffIDStr, 10, 64)
	if err != nil {
		// Handle the case where StaffID is invalid
		return c.RenderJSON(map[string]interface{}{"error": "Invalid StaffID in session"})
	}

	// Fetch the warehouse ID based on the staff ID from the database
	warehouse, err := models.GetWarehouseByStaffID(staffID)
	if err != nil {
		// Handle the case where warehouse ID retrieval fails
		return c.RenderJSON(map[string]interface{}{"error": "Failed to retrieve warehouse ID"})
	}

	// Store the warehouse ID in the session
	c.Session["WarehouseID"] = strconv.Itoa(int(warehouse.ID))

	// Retrieve the warehouse ID from the session after storing it
	warehouseIDInterface, err := c.Session.Get("WarehouseID")
	if err != nil {
		// Handle the case where WarehouseID is not found in session
		return c.RenderJSON(map[string]interface{}{"error": "WarehouseID not found in session"})
	}

	// Render the AdminDashboard view with the staff ID and warehouse ID
	return c.Render(staffIDInterface, warehouseIDInterface)
}

func (c *AdminDashboard) parseBodyClient(clients *[]models.Client) error {
	err := c.Params.BindJSON(clients)
	return err
}

func (c *AdminDashboard) parseBodyStaff(staffs *[]models.Staff) error {
	err := c.Params.BindJSON(staffs)
	return err
}

func (c *AdminDashboard) parseBodyUnit(units *[]models.Unit) error {
	err := c.Params.BindJSON(units)
	return err
}

func (c *AdminDashboard) parseBodyWarehouse(warehouses *[]models.Warehouse) error {
	err := c.Params.BindJSON(warehouses)
	return err
}

func (c AdminDashboard) Logout() revel.Result {
	return c.Redirect("/")
}

func (c AdminDashboard) GetClient(id int64) revel.Result {
	client, err := models.GetClient(id)
	if err != nil {
		c.Response.Status = 404 // or appropriate error status code
		return nil
	}
	c.Response.Status = 200 // Success status code
	return c.RenderJSON(client)
}

func (c AdminDashboard) AddClient() revel.Result {
	var errors map[string]interface{}
	var client models.Client

	// Bind JSON payload to the client struct
	if err := c.Params.BindJSON(&client); err != nil {
		c.Response.Status = 400
		errors = map[string]interface{}{"error": "Invalid JSON format"}
		return c.RenderJSON(errors)
	}

	// Add the client
	if err := client.AddClient(); err != nil {
		c.Response.Status = 503
		errors = map[string]interface{}{"error": "Unable to service your request. Please try again later"}
		return c.RenderJSON(errors)
	}

	// Success
	c.Response.Status = 201
	return nil
}

func (c AdminDashboard) UpdateClient(id int64) revel.Result {
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

func (c AdminDashboard) DeleteClient(id int64) revel.Result {
	// Check if the client is associated with any units
	hasUnits, err := models.ClientHasUnits(id)
	if err != nil {
		c.Response.Status = 500 // Internal server error status code
		return nil
	}
	if hasUnits {
		c.Response.Status = 400 // Bad request status code
		return c.RenderJSON(map[string]interface{}{"error": "Cannot delete client with associated units"})
	}

	if err := models.DeleteClient(id); err != nil {
		c.Response.Status = 500 // Internal server error status code
		return nil
	}
	c.Response.Status = 204 // No content status code
	return nil
}

func (c AdminDashboard) ListClients() revel.Result {
	clientModel := models.Client{}
	data, err := clientModel.ListClients()
	if err != nil {
		c.Response.Status = 204
		return nil
	}
	c.Response.Status = 200
	return c.RenderJSON(data)
}

func (c AdminDashboard) ListClientsByWarehouses() revel.Result {
	unitModel := models.Unit{}

	// Retrieve the warehouse ID from the session
	warehouseIDInterface := c.Session["WarehouseID"]
	warehouseIDStr, ok := warehouseIDInterface.(string)
	fmt.Println("Session WarehouseID:", warehouseIDInterface)

	if !ok {
		c.Response.Status = 400
		return c.RenderJSON(map[string]interface{}{"error": "Invalid WarehouseID in session"})
	}

	warehouseID, err := strconv.Atoi(warehouseIDStr)
	if err != nil {
		c.Response.Status = 400
		return c.RenderJSON(map[string]interface{}{"error": "Invalid WarehouseID format in session"})
	}

	// Call the model method to retrieve clients by warehouse
	data, err := unitModel.ListClientsByWarehouses(warehouseID)
	if err != nil {
		c.Response.Status = 500 // Internal Server Error
		fmt.Println("Error fetching clients by warehouse:", err)
		return c.RenderJSON(map[string]interface{}{"error": "Failed to fetch clients"})
	}

	c.Response.Status = 200
	return c.RenderJSON(data)
}

func (c AdminDashboard) AddStaff() revel.Result {
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

func (c AdminDashboard) GetStaff(id int64) revel.Result {
	staff, err := models.GetStaff(id)
	if err != nil {
		c.Response.Status = 404 // Not found status code
		return nil
	}
	c.Response.Status = 200 // Success status code
	return c.RenderJSON(staff)
}

func (c AdminDashboard) UpdateStaff(id int64) revel.Result {
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

func (c AdminDashboard) DeleteStaff(id int64) revel.Result {
	if err := models.DeleteStaff(id); err != nil {
		c.Response.Status = 503 // Service unavailable status code
		errors.Error = "Unable to service your request. Please try again later"
		return c.RenderJSON(errors)
	}
	c.Response.Status = 204 // No content status code
	return nil
}

func (c AdminDashboard) ListStaffs() revel.Result {
	staffModel := models.Staff{}
	data, err := staffModel.ListStaffs()
	if err != nil {
		c.Response.Status = 204 // No content status code
		return nil
	}
	c.Response.Status = 200 // Success status code
	return c.RenderJSON(data)
}

func (c AdminDashboard) AddUnit() revel.Result {
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

func (c AdminDashboard) GetUnit(id int64) revel.Result {
	unit, err := models.GetUnit(id)
	if err != nil {
		c.Response.Status = 404 // or appropriate error status code
		return nil
	}
	c.Response.Status = 200 // Success status code
	return c.RenderJSON(unit)
}

func (c AdminDashboard) UpdateUnit(id int64) revel.Result {
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

func (c AdminDashboard) DeleteUnit(id int64) revel.Result {
	if err := models.DeleteUnit(id); err != nil {
		c.Response.Status = 500 // Internal server error status code
		return nil
	}
	c.Response.Status = 204 // No content status code
	return nil
}

func (c AdminDashboard) ListLatestOrdersByWarehouses() revel.Result {
	unitModel := models.Unit{}

	// Retrieve the warehouse ID from the session
	warehouseIDInterface, err := c.Session.Get("WarehouseID")
	if err != nil {
		c.Response.Status = 400
		return c.RenderJSON(map[string]interface{}{"error": "WarehouseID not found in session"})
	}

	warehouseIDStr, ok := warehouseIDInterface.(string)
	if !ok {
		c.Response.Status = 400
		return c.RenderJSON(map[string]interface{}{"error": "Invalid WarehouseID format in session"})
	}

	warehouseID, err := strconv.Atoi(warehouseIDStr)
	if err != nil {
		c.Response.Status = 400
		return c.RenderJSON(map[string]interface{}{"error": "Invalid WarehouseID format in session"})
	}

	data, err := unitModel.ListLatestOrdersByWarehouses(warehouseID)
	if err != nil {
		c.Response.Status = 500 // Internal Server Error
		return c.RenderJSON(map[string]interface{}{"error": "Failed to fetch latest orders"})
	}
	c.Response.Status = 200
	return c.RenderJSON(data)
}

func (c AdminDashboard) GetTotalSales() revel.Result {
	unitModel := models.Unit{}
	data, err := unitModel.GetTotalSales()
	if err != nil {
		c.Response.Status = 204
		return nil
	}
	c.Response.Status = 200
	return c.RenderJSON(data)
}

func (c AdminDashboard) GetTotalSalesByWarehouse() revel.Result {
	unitModel := models.Unit{}

	// Retrieve the warehouse ID from the session
	warehouseIDInterface, err := c.Session.Get("WarehouseID")
	warehouseIDStr, ok := warehouseIDInterface.(string)
	fmt.Println("Session WarehouseID:", warehouseIDInterface)

	if !ok {
		c.Response.Status = 400
		return c.RenderJSON(map[string]interface{}{"error": "Invalid WarehouseID in session"})
	}

	warehouseID, err := strconv.Atoi(warehouseIDStr)
	if err != nil {
		c.Response.Status = 400
		return c.RenderJSON(map[string]interface{}{"error": "Invalid WarehouseID format in session"})
	}

	// Call the model method to retrieve total sales by warehouse
	data, err := unitModel.GetTotalSalesByWarehouse(warehouseID)
	if err != nil {
		c.Response.Status = 500 // Internal Server Error
		fmt.Println("Error fetching total sales by warehouse:", err)
		return c.RenderJSON(map[string]interface{}{"error": "Failed to fetch total sales"})
	}

	c.Response.Status = 200
	return c.RenderJSON(data)
}

func (c AdminDashboard) GetOrderCompleted() revel.Result {
	unitModel := models.Unit{}
	data, err := unitModel.GetOrderCompleted()
	if err != nil {
		c.Response.Status = 204
		return nil
	}
	c.Response.Status = 200
	return c.RenderJSON(data)
}

func (c AdminDashboard) GetOrderCompletedByWarehouse() revel.Result {
	unitModel := models.Unit{}

	// Retrieve the warehouse ID from the session
	warehouseIDInterface, err := c.Session.Get("WarehouseID")
	warehouseIDStr, ok := warehouseIDInterface.(string)
	fmt.Println("Session WarehouseID:", warehouseIDInterface)

	if !ok {
		c.Response.Status = 400
		return c.RenderJSON(map[string]interface{}{"error": "Invalid WarehouseID in session"})
	}

	warehouseID, err := strconv.Atoi(warehouseIDStr)
	if err != nil {
		c.Response.Status = 400
		return c.RenderJSON(map[string]interface{}{"error": "Invalid WarehouseID format in session"})
	}

	// Call the model method to retrieve completed orders by warehouse
	data, err := unitModel.GetOrderCompletedByWarehouse(warehouseID)
	if err != nil {
		c.Response.Status = 500 // Internal Server Error
		fmt.Println("Error fetching completed orders by warehouse:", err)
		return c.RenderJSON(map[string]interface{}{"error": "Failed to fetch completed orders"})
	}

	c.Response.Status = 200
	return c.RenderJSON(data)
}

func (c AdminDashboard) GetOrderPending() revel.Result {
	unitModel := models.Unit{}
	data, err := unitModel.GetOrderPending()
	if err != nil {
		c.Response.Status = 204
		return nil
	}
	c.Response.Status = 200
	return c.RenderJSON(data)
}

func (c AdminDashboard) GetOrderPendingByWarehouse() revel.Result {
	unitModel := models.Unit{}

	// Retrieve the warehouse ID from the session
	warehouseIDInterface, err := c.Session.Get("WarehouseID")
	warehouseIDStr, ok := warehouseIDInterface.(string)
	fmt.Println("Session WarehouseID:", warehouseIDInterface)

	if !ok {
		c.Response.Status = 400
		return c.RenderJSON(map[string]interface{}{"error": "Invalid WarehouseID in session"})
	}

	warehouseID, err := strconv.Atoi(warehouseIDStr)
	if err != nil {
		c.Response.Status = 400
		return c.RenderJSON(map[string]interface{}{"error": "Invalid WarehouseID format in session"})
	}

	// Call the model method to retrieve pending orders by warehouse
	data, err := unitModel.GetOrderPendingByWarehouse(warehouseID)
	if err != nil {
		c.Response.Status = 500 // Internal Server Error
		fmt.Println("Error fetching pending orders by warehouse:", err)
		return c.RenderJSON(map[string]interface{}{"error": "Failed to fetch pending orders"})
	}

	c.Response.Status = 200
	return c.RenderJSON(data)
}

func (c AdminDashboard) GetLatestOrders() revel.Result {
	unitModel := models.Unit{}
	data, err := unitModel.GetLatestOrders()
	if err != nil {
		c.Response.Status = 204
		return nil
	}
	c.Response.Status = 200
	return c.RenderJSON(data)
}

func (c AdminDashboard) UpdateOrderStatus(id int64) revel.Result {
	var unit models.Unit

	// Bind the JSON payload to the unit struct
	if err := c.Params.BindJSON(&unit); err != nil {
		c.Response.Status = 400
		return c.RenderJSON(map[string]interface{}{"error": "Invalid JSON format"})
	}

	// Update the order status in the database and check the result
	updatedUnit, err := models.UpdateOrderStatus(&unit)
	if err != nil {
		if err.Error() == "order is already completed" {
			c.Response.Status = 400 // Bad request status code
			return c.RenderJSON(map[string]interface{}{"error": "Order is already completed"})
		}
		c.Response.Status = 500 // Internal server error status code
		return c.RenderJSON(map[string]interface{}{"error": "Failed to update order status"})
	}

	// Create a response map with the updated ID and status
	response := map[string]interface{}{
		"id":     updatedUnit.ID,
		"status": updatedUnit.Status,
	}

	// Return the updated response as JSON
	c.Response.Status = 200 // Success status code
	return c.RenderJSON(response)
}

func (c AdminDashboard) AddWarehouse() revel.Result {
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

func (c AdminDashboard) GetWarehouse(id int64) revel.Result {
	warehouse, err := models.GetWarehouse(id)
	if err != nil {
		c.Response.Status = 404
		return nil
	}
	c.Response.Status = 200 // Success status code
	return c.RenderJSON(warehouse)
}

func (c AdminDashboard) GetWarehouseByStaffID(id int64) revel.Result {
	warehouse, err := models.GetWarehouseByStaffID(id)
	if err != nil {
		c.Response.Status = 404
		return nil
	}

	// Create a map or struct to hold the JSON data
	responseData := map[string]interface{}{
		"warehouse_id": warehouse.ID,
	}

	c.Response.Status = 200 // Success status code
	return c.RenderJSON(responseData)
}

func (c AdminDashboard) UpdateWarehouse(id int64) revel.Result {
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

func (c AdminDashboard) DeleteWarehouse(id int64) revel.Result {
	if err := models.DeleteWarehouse(id); err != nil {
		c.Response.Status = 500 // Internal server error status code
		return nil
	}
	c.Response.Status = 204 // No content status code
	return nil
}

func (c AdminDashboard) ListWarehouses(page int64, perPage int64) revel.Result {
	warehouseModel := models.Warehouse{}
	data, err := warehouseModel.ListWarehouses()
	if err != nil {
		c.Response.Status = 204
		return nil
	}
	c.Response.Status = 200
	return c.RenderJSON(data)
}
