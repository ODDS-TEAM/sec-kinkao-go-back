package models

import "gopkg.in/mgo.v2/bson"

type (
	User struct {
		ID       bson.ObjectId `bson:"_id" json:"_id"`
		Username string        `bson:"username" json:"username"`
		Password string        `bson:"password" json:"password"`
		Email    string        `bson:"email" json:"email"`
	}
	// CreatedAt time.Time     `bson:"createAt" json:"createAt"`
	// StoreInfo Info          `bson:"storeInfo" json:"storeInfo"`
	// Info struct {
	// 	Name        string `bson:"name" json:"name"`
	// 	Owner       string `bson:"owner" json:"owner"`
	// 	Describe    string `bson:"describe" json:"describe"`
	// 	PhoneNumber string `bson:"phoneNumber" json:"phoneNumber"`
	// 	ImgUrl      string `bson:"imgUrl" json:"imgUrl"`
	// }
)
