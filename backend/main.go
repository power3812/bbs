package main

import (
	"backend/helper"
	"backend/router"
)

func main() {
	err := helper.SetJst()
	if err != nil {
		panic(err)
	}

	router.Init()
}
