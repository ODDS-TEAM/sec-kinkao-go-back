package handlers

import "github.com/labstack/echo"

import "net/http"

func (h *Handler) Login(c echo.Context) (err error) {
	return c.JSON(http.StatusOK, {"message": "OK"})
}
