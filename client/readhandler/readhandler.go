package readhandler

import (
	"Benz-assignment/client/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo"
)

const readServerAddress = "http://localhost:8082/read"

func ReadHandler(context echo.Context) error {
	fileTypeToRead := context.Param("filetype")

	var data []byte
	if fileTypeToRead == "csv" {
		fmt.Println(readServerAddress + "/csv")
		resp, err := http.Get(readServerAddress + "/csv")
		if err != nil {
			fmt.Println(err)
		}
		data, _ = ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		fmt.Println(string(data))
	} else if fileTypeToRead == "xml" {
		resp, err := http.Get(readServerAddress + "/xml")
		if err != nil {
			fmt.Println(err)
		}
		data, _ = ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		fmt.Println(string(data))
	}
	response := string(data)
	bytes := []byte(response)

	var allPersonDetails []model.Person
	json.Unmarshal(bytes, &allPersonDetails)
	//Encoding and sending the recorded response
	context.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	context.Response().WriteHeader(http.StatusOK)
	return json.NewEncoder(context.Response()).Encode(allPersonDetails)
}
