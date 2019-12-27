package api

import (
	"gitlab.odds.team/internship/sec-kinkao/goback/model"
)

func (db *MongoDB) CreateSellerUser(su *model.SellerUser) {
	if err := db.SellerUCol.Insert(&su); err != nil {
		return
	}
}
