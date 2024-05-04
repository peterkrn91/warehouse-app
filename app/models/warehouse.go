package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/revel/revel"
)

type Client struct {
	ID   int64  `gorm:"primaryKey"`
	Name string `gorm:"type:varchar(255);default:NULL"`
}

type Staff struct {
	ID       int64  `gorm:"primaryKey"`
	Name     string `json:"name"`
	UserType string `json:"user_type"`
	Password string `json:"password"`
}

type Unit struct {
	ID          int64   `gorm:"primaryKey"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Kg          float64 `json:"kg"`
	WarehouseID int     `json:"warehouse_id"`
	ClientID    int     `json:"client_id"`
	Status      string  `json:"status"`
	OrderDate   string  `json:"order_date"`
}

type Warehouse struct {
	ID      int64  `gorm:"primaryKey"`
	Name    string `json:"name"`
	StaffID int    `json:"staff_id"`
}

// GetClient retrieves a client by ID.
func GetClient(id int64) (Client, error) {
	var client Client
	response := DB.First(&client, id)
	return client, response.Error
}

// AddClient adds a new client to the database.
func (client Client) AddClient() error {
	response := DB.Create(&client)
	return response.Error
}

// UpdateClient updates an existing client in the database.
func UpdateClient(client *Client) error {
	response := DB.Save(client)
	return response.Error
}

// DeleteClient deletes a client from the database based on the ID.
func DeleteClient(id int64) error {
	var client Client
	response := DB.Delete(&client, id)
	return response.Error
}

// ClientHasUnits checks if a client is associated with any units.
func ClientHasUnits(clientID int64) (bool, error) {
	var count int64
	response := DB.Model(&Unit{}).Where("client_id = ?", clientID).Count(&count)
	return count > 0, response.Error
}

// ListClients retrieves a list of all clients from the database.
func (client Client) ListClients() ([]Client, error) {
	clients := make([]Client, 0)
	responses := DB.Find(&clients)
	return clients, responses.Error
}

// ListClientsByWarehouses retrieves a list of clients (client ID and name) filtered by warehouse ID from the database using a join query.
func (unit Unit) ListClientsByWarehouses(warehouseID int) ([]map[string]interface{}, error) {
	// Define a struct to hold the join result
	type JoinResult struct {
		ClientID int64  `gorm:"column:ClientID" json:"client_id"`
		Name     string `gorm:"column:Name" json:"name"`
	}
	var joinResults []JoinResult

	// Perform a join query to fetch client IDs and names for the specified warehouse
	response := DB.Table("clients").
		Select("DISTINCT clients.id AS ClientID, clients.name AS Name").
		Joins("JOIN units ON units.client_id = clients.id").
		Where("units.warehouse_id = ?", warehouseID).
		Scan(&joinResults)
	if response.Error != nil {
		return nil, response.Error
	}

	// Convert join results to the desired format
	clients := make([]map[string]interface{}, 0)
	for _, result := range joinResults {
		clientData := map[string]interface{}{
			"client_id": result.ClientID,
			"name":      result.Name,
		}
		clients = append(clients, clientData)
	}

	return clients, nil
}

// AddStaff adds a new staff member to the database.
func (staff *Staff) AddStaff() error {
	response := DB.Create(staff)
	return response.Error
}

// GetStaff retrieves a staff member by ID.
func GetStaff(id int64) (Staff, error) {
	var staff Staff
	response := DB.First(&staff, id)
	return staff, response.Error
}

// UpdateStaff updates an existing staff member in the database.
func UpdateStaff(staff *Staff) error {
	response := DB.Save(staff)
	return response.Error
}

// DeleteStaff deletes a staff member from the database based on the ID.
func DeleteStaff(id int64) error {
	var staff Staff
	staff.ID = id
	response := DB.Delete(&staff)
	return response.Error
}

// ListStaff retrieves a list of all staff members from the database.
func (staff Staff) ListStaffs() ([]Staff, error) {
	staffs := make([]Staff, 0)
	response := DB.Find(&staffs)
	return staffs, response.Error
}

// ValidateLogin checks the login credentials and retrieves staff information by name and password
func ValidateLogin(name, password string) (*Staff, string, error) {
	var id int64
	var fetchedStaff Staff
	fetchedStaff.ID = id
	result := DB.Where("name = ? AND password = ?", name, password).First(&fetchedStaff)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return nil, "", result.Error
	}

	return &fetchedStaff, fetchedStaff.UserType, nil
}

func (staff *Staff) LogoutUser(c *revel.Controller) {
	// Clear session data
	c.Session.Del("SessionID")
	c.Redirect("/")
}

// AddUnit adds a new unit to the database.
func (unit *Unit) AddUnit() error {
	response := DB.Create(unit)
	return response.Error
}

// GetUnit retrieves a unit by ID.
func GetUnit(id int64) (Unit, error) {
	var unit Unit
	response := DB.First(&unit, id)
	return unit, response.Error
}

// UpdateUnit updates an existing unit in the database.
func UpdateUnit(unit *Unit) error {
	response := DB.Save(unit)
	return response.Error
}

// DeleteUnit deletes a unit from the database based on the ID.
func DeleteUnit(id int64) error {
	var unit Unit
	unit.ID = id
	response := DB.Delete(&unit)
	return response.Error
}

// ListLatestOrdersInEachWarehouses retrieves a list of units filtered by warehouse ID and ordered by order date descending from the database.
func (unit Unit) ListLatestOrdersByWarehouses(warehouseID int) ([]Unit, error) {
	units := make([]Unit, 0)
	response := DB.Where("warehouse_id = ?", warehouseID).Order("order_date DESC").Find(&units)
	return units, response.Error
}

// TotalSales gets total units and multiply it by a dollar
func (unit Unit) GetTotalSales() (float64, error) {
	var totalSales float64

	// Fetch completed units and calculate total sales
	result := DB.Model(&Unit{}).Where("status = ?", "completed").Find(&[]Unit{})
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return totalSales, result.Error
	}

	// Count the number of completed units and multiply by $100 each
	completedUnits := len(*result.Value.(*[]Unit))
	totalSales = float64(completedUnits) * 100.0

	return totalSales, nil
}

func (unit Unit) GetTotalSalesByWarehouse(warehouseID int) (float64, error) {
	var totalSales float64

	// Fetch completed units based on warehouse ID
	result := DB.Model(&Unit{}).Where("status = ? AND warehouse_id = ?", "completed", warehouseID).Find(&[]Unit{})
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return totalSales, result.Error
	}

	// Count the number of completed units and multiply by $100 each
	completedUnits := len(*result.Value.(*[]Unit))
	totalSales = float64(completedUnits) * 100.0

	return totalSales, nil
}

// GetOrderCompleted retrieves all units with status "completed"
func (unit Unit) GetOrderCompleted() ([]Unit, error) {
	var completedUnits []Unit

	// Fetch completed units
	result := DB.Where("status = ?", "completed").Find(&completedUnits)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return nil, result.Error
	}

	return completedUnits, nil
}

func (unit Unit) GetOrderCompletedByWarehouse(warehouseID int) ([]Unit, error) {
	completedUnits := make([]Unit, 0)
	response := DB.Where("status = ? AND warehouse_id = ?", "completed", warehouseID).Find(&completedUnits)
	return completedUnits, response.Error
}

// GetOrderCompleted retrieves all units with status "completed"
func (unit Unit) GetOrderPending() ([]Unit, error) {
	var pendingUnits []Unit

	// Fetch completed units
	result := DB.Where("status = ?", "pending").Find(&pendingUnits)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return nil, result.Error
	}

	return pendingUnits, nil
}

func (unit Unit) GetOrderPendingByWarehouse(warehouseID int) ([]Unit, error) {
	pendingUnits := make([]Unit, 0)
	response := DB.Where("status = ? AND warehouse_id = ?", "pending", warehouseID).Find(&pendingUnits)
	if response.Error != nil {
		fmt.Println("Error fetching pending orders:", response.Error)
	}
	return pendingUnits, response.Error
}

// GetLatestOrders retrieves the three latest orders
func (unit Unit) GetLatestOrders() ([]Unit, error) {
	var latestOrders []Unit

	// Fetch the three latest orders based on order date
	result := DB.Order("order_date DESC").Find(&latestOrders)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return nil, result.Error
	}

	return latestOrders, nil
}

// UpdateOrderStatus updates an existing unit's order status in the database (from pending -> completed) if it's currently pending.
func UpdateOrderStatus(unit *Unit) (Unit, error) {
	// Fetch the existing unit based on the provided ID
	existingUnit := &Unit{}
	result := DB.First(existingUnit, unit.ID)
	if result.Error != nil {
		return *existingUnit, result.Error
	}

	// Check if the existing unit's status is already "completed"
	if existingUnit.Status == "completed" {
		return *existingUnit, result.Error
	}

	// Update the status to "completed" and save the changes
	existingUnit.Status = "completed"
	result = DB.Save(existingUnit)
	if result.Error != nil {
		return *existingUnit, result.Error
	}

	return *existingUnit, result.Error // No error, indicating successful update
}

// GetClientUnits retrieves all units belonging to a specific client
func GetClientOrders(clientID int) ([]Unit, error) {
	var clientUnits []Unit

	// Fetch units belonging to the specified client
	result := DB.Where("client_id = ?", clientID).Find(&clientUnits)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return nil, result.Error
	}

	return clientUnits, nil
}

// AddWarehouse adds a new warehouse to the database.
func (warehouse *Warehouse) AddWarehouse() error {
	response := DB.Create(warehouse)
	return response.Error
}

// GetWarehouse retrieves a warehouse by ID.
func GetWarehouse(id int64) (Warehouse, error) {
	var warehouse Warehouse
	response := DB.First(&warehouse, id)
	return warehouse, response.Error
}

// GetWarehouseByStaffID retrieves a warehouse by staff ID.
func GetWarehouseByStaffID(staffID int64) (Warehouse, error) {
	var warehouse Warehouse
	response := DB.Where("staff_id = ?", staffID).First(&warehouse)
	return warehouse, response.Error
}

// UpdateWarehouse updates an existing warehouse in the database.
func UpdateWarehouse(warehouse *Warehouse) error {
	response := DB.Save(warehouse)
	return response.Error
}

// DeleteWarehouse deletes a warehouse from the database based on the ID.
func DeleteWarehouse(id int64) error {
	var warehouse Warehouse
	warehouse.ID = id
	response := DB.Delete(&warehouse)
	return response.Error
}

// ListWarehouses retrieves a list of all warehouses from the database.
func (warehouse Warehouse) ListWarehouses() ([]Warehouse, error) {
	warehouses := make([]Warehouse, 0)
	response := DB.Find(&warehouses)
	return warehouses, response.Error
}
