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


type tApp struct {}
var App tApp


func (_ tApp) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Index", args).URL
}

func (_ tApp) Logout(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.Logout", args).URL
}

func (_ tApp) GetClient(
		id int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("App.GetClient", args).URL
}

func (_ tApp) AddClient(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.AddClient", args).URL
}

func (_ tApp) UpdateClient(
		id int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("App.UpdateClient", args).URL
}

func (_ tApp) DeleteClient(
		id int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("App.DeleteClient", args).URL
}

func (_ tApp) ListClients(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.ListClients", args).URL
}

func (_ tApp) AddStaff(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.AddStaff", args).URL
}

func (_ tApp) GetStaff(
		id int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("App.GetStaff", args).URL
}

func (_ tApp) UpdateStaff(
		id int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("App.UpdateStaff", args).URL
}

func (_ tApp) DeleteStaff(
		id int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("App.DeleteStaff", args).URL
}

func (_ tApp) ListStaffs(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.ListStaffs", args).URL
}

func (_ tApp) GetLoginStatus(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.GetLoginStatus", args).URL
}

func (_ tApp) AssignUnitsToClient(
		clientID int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "clientID", clientID)
	return revel.MainRouter.Reverse("App.AssignUnitsToClient", args).URL
}

func (_ tApp) AddUnit(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.AddUnit", args).URL
}

func (_ tApp) GetUnit(
		id int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("App.GetUnit", args).URL
}

func (_ tApp) UpdateUnit(
		id int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("App.UpdateUnit", args).URL
}

func (_ tApp) DeleteUnit(
		id int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("App.DeleteUnit", args).URL
}

func (_ tApp) ListUnits(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.ListUnits", args).URL
}

func (_ tApp) GetTotalSales(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.GetTotalSales", args).URL
}

func (_ tApp) GetOrderCompleted(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.GetOrderCompleted", args).URL
}

func (_ tApp) GetOrderPending(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.GetOrderPending", args).URL
}

func (_ tApp) GetLatestOrders(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.GetLatestOrders", args).URL
}

func (_ tApp) AddWarehouse(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.AddWarehouse", args).URL
}

func (_ tApp) GetWarehouse(
		id int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("App.GetWarehouse", args).URL
}

func (_ tApp) UpdateWarehouse(
		id int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("App.UpdateWarehouse", args).URL
}

func (_ tApp) DeleteWarehouse(
		id int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "id", id)
	return revel.MainRouter.Reverse("App.DeleteWarehouse", args).URL
}

func (_ tApp) ListWarehouses(
		page int64,
		perPage int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "page", page)
	revel.Unbind(args, "perPage", perPage)
	return revel.MainRouter.Reverse("App.ListWarehouses", args).URL
}


type tOwnerDashboard struct {}
var OwnerDashboard tOwnerDashboard


func (_ tOwnerDashboard) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("OwnerDashboard.Index", args).URL
}


