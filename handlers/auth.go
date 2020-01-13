package handlers

import (
	"net/http"

	"github.com/labstack/echo"
	"gitlab.odds.team/internship/sec-kinkao/goback/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func (h *Handler) Login(c echo.Context) (err error) {
	u := &models.User{}
	if err = c.Bind(u); err != nil {
		return
	}

	db := h.DB.Clone()
	defer db.Close()
	if err = db.DB("twitter").C("users").
		Find(bson.M{"email": u.Email, "password": u.Password}).One(u); err != nil {
		if err == mgo.ErrNotFound {
			return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "invalid email or password"}
		}
		return
	}

	return c.JSON(http.StatusOK, u)
}
