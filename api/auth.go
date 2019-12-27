package api

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
	"gitlab.odds.team/internship/sec-kinkao/goback/model"
	"gopkg.in/mgo.v2/bson"
)

func (db *MongoDB) Register(c echo.Context) error {
	su := &model.SellerUser{
		ID:        bson.NewObjectId(),
		CreatedAt: time.Now(),
	}

	if err := c.Bind(su); err != nil {
		return err
	}

	db.CreateSellerUser(su)

	return c.JSON(http.StatusCreated, su)
}
