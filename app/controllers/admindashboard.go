package controllers

import (
	"github.com/revel/revel"
)

type AdminDashboard struct {
	*revel.Controller
}

func (c AdminDashboard) Index() revel.Result {
	return c.Render()
}
