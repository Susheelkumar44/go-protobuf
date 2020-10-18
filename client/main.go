package main

import (
	readdata "Benz-assignment/client/readhandler"
	postdata "Benz-assignment/client/savehandler"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.POST("/save/:filetype", postdata.SaveHandler)
	e.GET("/read/:filetype", readdata.ReadHandler)
	e.Start(":8081")

}
