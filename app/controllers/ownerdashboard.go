package controllers

import (
	"github.com/revel/revel"
)

type OwnerDashboard struct {
	*revel.Controller
}

func (c OwnerDashboard) Index(ownerID int64) revel.Result {
	// Retrieve the OwnerID from the session after storing it
	ownerIDInterface, err := c.Session.Get("OwnerID")
	if err != nil {
		// Handle the case where OwnerID is not found in session
		return c.RenderJSON(map[string]interface{}{"error": "OwnerID not found in session"})
	}

	// Render the OwnerDashboard view with the OwnerID
	return c.Render(ownerIDInterface)
}
