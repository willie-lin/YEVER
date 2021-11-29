package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/willie-lin/YEVER/pkg/database/ent"
	"net/http"
)

func (controller *Controller) CreateUser(c echo.Context) error {

	var user ent.User

	// bind request to user struct
	if err := c.Bind(&user); err != nil {
		c.Logger().Error("Bind:", err)
		return c.String(http.StatusBadRequest, "Bind: "+err.Error())
	}

	// insert record

	cc := controller.client.User.Create().SetDescription(user.Description)
	if user.Name != "" {
		cc.SetName(user.Name)
	}
	if user.Password != "" {
		cc.SetPassword(user.Password)

	}
	if user.Email != "" {
		cc.SetEmail(user.Email)
	}
	if user.Age >= 0 {
		cc.SetAge(user.Age)

	}

	return nil
}
