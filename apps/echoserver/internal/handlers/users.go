package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUsers(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "id:"+id+",team:"+team+", member:"+member)
}
