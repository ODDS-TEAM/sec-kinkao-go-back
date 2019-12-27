package model

import "gopkg.in/mgo.v2/bson"

type (
	SellerUser struct {
		ID        bson.ObjectId
		Username  string
		Password  string
		Email     string
		StoreInfo StoreInfo
	}
	StoreInfo struct {
		Name        string
		Owner       string
		Describe    string
		PhoneNumber string
		ImgUrl      string
	}
)
