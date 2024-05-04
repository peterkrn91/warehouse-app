// GENERATED CODE - DO NOT EDIT
// This file provides a way of creating URL's based on all the actions
// found in all the controllers.
package routes

import "github.com/revel/revel"


type tStatic struct {}
var Static tStatic


func (_ tStatic) Serve(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.Serve", args).URL
}

func (_ tStatic) ServeDir(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeDir", args).URL
}

func (_ tStatic) ServeModule(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModule", args).URL
}

func (_ tStatic) ServeModuleDir(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModuleDir", args).URL
}


type tTestRunner struct {}
var TestRunner tTestRunner


func (_ tTestRunner) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.Index", args).URL
}

func (_ tTestRunner) Suite(
		suite string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	return revel.MainRouter.Reverse("TestRunner.Suite", args).URL
}

func (_ tTestRunner) Run(
		suite string,
		test string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	revel.Unbind(args, "test", test)
	return revel.MainRouter.Reverse("TestRunner.Run", args).URL
}

func (_ tTestRunner) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.List", args).URL
}


type tAdminDashboard struct {}
var AdminDashboard tAdminDashboard


func (_ tAdminDashboard) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("AdminDashboard.Index", args).URL
}

func (_ tAdminDashboard) Logout(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("AdminDashboard.Logout", args).URL
}

func (_ tAdminDashboard) GetClient(
		id int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("AdminDashboard.GetClient", args).URL
}

func (_ tAdminDashboard) AddClient(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("AdminDashboard.AddClient", args).URL
}

func (_ tAdminDashboard) UpdateClient(
		id int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("AdminDashboard.UpdateClient", args).URL
}

func (_ tAdminDashboard) DeleteClient(
		id int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("AdminDashboard.DeleteClient", args).URL
}

func (_ tAdminDashboard) ListClients(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("AdminDashboard.ListClients", args).URL
}

func (_ tAdminDashboard) ListClientsByWarehouses(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("AdminDashboard.ListClientsByWarehouses", args).URL
}

func (_ tAdminDashboard) AddStaff(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("AdminDashboard.AddStaff", args).URL
}

func (_ tAdminDashboard) GetStaff(
		id int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("AdminDashboard.GetStaff", args).URL
}

func (_ tAdminDashboard) UpdateStaff(
		id int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("AdminDashboard.UpdateStaff", args).URL
}

func (_ tAdminDashboard) DeleteStaff(
		id int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("AdminDashboard.DeleteStaff", args).URL
}

func (_ tAdminDashboard) ListStaffs(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("AdminDashboard.ListStaffs", args).URL
}

func (_ tAdminDashboard) AddUnit(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("AdminDashboard.AddUnit", args).URL
}

func (_ tAdminDashboard) GetUnit(
		id int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("AdminDashboard.GetUnit", args).URL
}

func (_ tAdminDashboard) UpdateUnit(
		id int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("AdminDashboard.UpdateUnit", args).URL
}

func (_ tAdminDashboard) DeleteUnit(
		id int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("AdminDashboard.DeleteUnit", args).URL
}

func (_ tAdminDashboard) ListLatestOrdersByWarehouses(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("AdminDashboard.ListLatestOrdersByWarehouses", args).URL
}

func (_ tAdminDashboard) GetTotalSales(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("AdminDashboard.GetTotalSales", args).URL
}

func (_ tAdminDashboard) GetTotalSalesByWarehouse(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("AdminDashboard.GetTotalSalesByWarehouse", args).URL
}

func (_ tAdminDashboard) GetOrderCompleted(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("AdminDashboard.GetOrderCompleted", args).URL
}

func (_ tAdminDashboard) GetOrderCompletedByWarehouse(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("AdminDashboard.GetOrderCompletedByWarehouse", args).URL
}

func (_ tAdminDashboard) GetOrderPending(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("AdminDashboard.GetOrderPending", args).URL
}

func (_ tAdminDashboard) GetOrderPendingByWarehouse(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("AdminDashboard.GetOrderPendingByWarehouse", args).URL
}

func (_ tAdminDashboard) GetLatestOrders(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("AdminDashboard.GetLatestOrders", args).URL
}

func (_ tAdminDashboard) UpdateOrderStatus(
		id int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("AdminDashboard.UpdateOrderStatus", args).URL
}

func (_ tAdminDashboard) AddWarehouse(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("AdminDashboard.AddWarehouse", args).URL
}

func (_ tAdminDashboard) GetWarehouse(
		id int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("AdminDashboard.GetWarehouse", args).URL
}

func (_ tAdminDashboard) GetWarehouseByStaffID(
		id int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("AdminDashboard.GetWarehouseByStaffID", args).URL
}

func (_ tAdminDashboard) UpdateWarehouse(
		id int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("AdminDashboard.UpdateWarehouse", args).URL
}

func (_ tAdminDashboard) DeleteWarehouse(
		id int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("AdminDashboard.DeleteWarehouse", args).URL
}

func (_ tAdminDashboard) ListWarehouses(
		page int64,
		perPage int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "page", page)
	revel.Unbind(args, "perPage", perPage)
	return revel.MainRouter.Reverse("AdminDashboard.ListWarehouses", args).URL
}


type tLoginPage struct {}
var LoginPage tLoginPage


func (_ tLoginPage) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("LoginPage.Index", args).URL
}

func (_ tLoginPage) UserLogin(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("LoginPage.UserLogin", args).URL
}


type tOwnerDashboard struct {}
var OwnerDashboard tOwnerDashboard


func (_ tOwnerDashboard) Index(
		ownerID int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "ownerID", ownerID)
	return revel.MainRouter.Reverse("OwnerDashboard.Index", args).URL
}


