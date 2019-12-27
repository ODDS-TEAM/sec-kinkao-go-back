package route

import (
	"github.com/labstack/echo"
	"gitlab.odds.team/internship/sec-kinkao/goback/api"
)

func Init(e *echo.Echo) {
	db, err := api.NewMongoDB()
	if err != nil {
		e.Logger.Fatal(err)
	}

	a := &api.MongoDB{
		Conn:       db.Conn,
		SellerUCol: db.SellerUCol,
	}

	e.POST("/register", a.Register)
}
