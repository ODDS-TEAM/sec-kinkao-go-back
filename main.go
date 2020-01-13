package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"github.com/xesina/golang-echo-realworld-example-app/handler"
	"gitlab.odds.team/internship/sec-kinkao/goback/config"
	"gopkg.in/mgo.v2"
)

func main() {
	e := echo.New()
	s := config.Spec()

	e.Logger.SetLevel(log.ERROR)
	e.Use(
		middleware.CORS(),
		middleware.Recover(),
		middleware.Logger(),
	)

	db, err := mgo.Dial("localhost")
	if err != nil {
		e.Logger.Fatal(err)
	}

	h := &handler.Handler{DB: db}

	e.POST("/login", h.Login)

	e.Logger.Fatal(e.Start(s.APIPort))
}
