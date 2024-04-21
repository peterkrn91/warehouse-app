package controllers

import (
	"github.com/revel/revel"
)

type OwnerDashboard struct {
	*revel.Controller
}

func (c OwnerDashboard) Index() revel.Result {
	return c.Render()
}
