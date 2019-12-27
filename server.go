package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"gitlab.odds.team/internship/sec-kinkao/goback/config"
	"gitlab.odds.team/internship/sec-kinkao/goback/route"
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

	e.GET("/_ah/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "sec kinkao goback : ok!")
	})

	route.Init(e)

	e.Logger.Fatal(e.Start(s.APIPort))
}
