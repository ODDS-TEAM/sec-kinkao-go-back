package model

import "gopkg.in/mgo.v2/bson"

type (
	FoodMenu struct {
		ID       bson.ObjectId
		SellerID bson.ObjectId
		Menus    []Menu
	}
	Menu struct {
		ID       bson.ObjectId
		Name     string
		Describe string
		Price    int
		ImgUrl   string
		Options  []Option
	}
	Option struct {
		ID          bson.ObjectId
		Title       string
		Require     bool
		MultiChoice bool
		Choices     []Choice
	}
	Choice struct {
		ID    bson.Setter
		Name  string
		Price int
	}
)
