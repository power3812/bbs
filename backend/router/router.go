package router

import (
	"fmt"
	"net/http"
	"os"

	controller "backend/controller/api/v1"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Init() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		fmt.Fprintf(os.Stderr, "Request: %v\n", string(reqBody))
	}))

	e.GET("/healthcheck", func(c echo.Context) error {
		response := map[string]string{"status": "ok"}
		return c.JSON(http.StatusOK, response)
	})

	api := e.Group("/api/v1")
	api.GET("/posts", controller.Index)
	api.POST("/posts", controller.Create)

	e.Logger.Fatal(e.Start(":5000"))
}
