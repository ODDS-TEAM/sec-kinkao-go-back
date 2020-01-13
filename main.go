package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"gitlab.odds.team/internship/sec-kinkao/goback/config"
	"gitlab.odds.team/internship/sec-kinkao/goback/handlers"
	"gopkg.in/mgo.v2"
)

func main() {
	e := echo.New()
	s := config.Spec()

	e.Logger.SetLevel(log.ERROR)
	// e.Use(
	// 	middleware.CORS(),
	// 	middleware.Recover(),
	// 	middleware.Logger(),
	// )

	db, err := mgo.Dial("localhost")
	if err != nil {
		e.Logger.Fatal(err)
	}

	h := &handlers.Handler{DB: db}

	e.POST("/login", h.Login)

	e.Logger.Fatal(e.Start(s.APIPort))
}
