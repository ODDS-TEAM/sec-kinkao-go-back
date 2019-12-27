package api

import (
	"fmt"

	"gitlab.odds.team/internship/sec-kinkao/goback/config"
	"gopkg.in/mgo.v2"
)

type (
	MongoDB struct {
		Conn       *mgo.Session
		SellerUCol *mgo.Collection
	}
)

func NewMongoDB() (*MongoDB, error) {
	s := config.Spec()
	conn, err := mgo.Dial(s.DBHost)
	fmt.Println("mongo: dial at ", s.DBHost)
	if err != nil {
		return nil, fmt.Errorf("mongo: could not dial on %s: %v", s.DBHost, err)
	}

	return &MongoDB{
		Conn:       conn,
		SellerUCol: conn.DB(s.DBName).C(s.DBSellerUsersCol),
	}, nil
}

func (db *MongoDB) Close() {
	db.Conn.Close()
}
