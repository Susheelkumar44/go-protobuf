package main

import (
	"Benz-assignment/readserver/readhandler"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/read/:filetype", readhandler.ReadHandler)
	e.Start(":8082")

}
