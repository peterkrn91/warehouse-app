package models

import (
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

// ListClients retrieves a list of all clients from the database.
func (client Client) ListClients() ([]Client, error) {
	clients := make([]Client, 0)
	responses := DB.Find(&clients)
	return clients, responses.Error
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

func (staff *Staff) GetLoginStatus() (*Staff, error) {
	var fetchedStaff Staff
	result := DB.Where("name = ? AND password = ?", staff.Name, staff.Password).First(&fetchedStaff)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return nil, result.Error
	}
	return &fetchedStaff, nil
}

func (staff *Staff) LogoutUser(c *revel.Controller) {
	// Clear session data
	c.Session["id"] = nil
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

// ListUnits retrieves a list of all units from the database.
func (unit Unit) ListUnits() ([]Unit, error) {
	units := make([]Unit, 0)
	response := DB.Find(&units)
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

	// Count the number of completed units and multiply by $5 each
	completedUnits := len(*result.Value.(*[]Unit))
	totalSales = float64(completedUnits) * 5.0 // Assuming each completed unit contributes $5 to total sales

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

// GetLatestOrders retrieves the three latest orders
func (unit Unit) GetLatestOrders() ([]Unit, error) {
	var latestOrders []Unit

	// Fetch the three latest orders based on order date
	result := DB.Order("order_date DESC").Limit(3).Find(&latestOrders)
	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return nil, result.Error
	}

	return latestOrders, nil
}

// UpdateUnit updates an existing unit in the database.
func UpdateOrderStatus(unit *Unit) error {
	response := DB.Where("id = ?", unit.ID).Save(unit)
	return response.Error
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
