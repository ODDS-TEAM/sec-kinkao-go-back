package model

import "gopkg.in/mgo.v2/bson"

import "time"

type (
	WeekMenu struct {
		ID   bson.ObjectId
		Week bson.ObjectId
		Days []Day
	}
	Day struct {
		ID     bson.ObjectId
		Day    string
		Date   time.Time
		Stores []Store
	}
	Store struct {
		ID       bson.ObjectId
		SellerID bson.ObjectId
		Name     string
		Menus    []DayMenu
	}
	DayMenu struct {
		ID     bson.ObjectId
		MenuID bson.ObjectId
		Name   string
		ImgUrl string
	}
)
