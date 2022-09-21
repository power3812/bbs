package controller

import (
	"net/http"

	post "backend/model/post"
	"fmt"

	"github.com/labstack/echo"
)

func Index(c echo.Context) error {
	p, err := post.GetAll()

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusOK, p)
}

func Create(c echo.Context) error {
	p := &post.Post{}

	if err := c.Bind(&p); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	p, err := p.Create()

	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	fmt.Println(p)

	return c.JSON(http.StatusOK, p)
}
