package readhandler

import (
	"io/ioutil"

	"github.com/labstack/echo"

	"Benz-assignment/readserver/model"
	"encoding/csv"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func ReadHandler(context echo.Context) error {
	fileTypeToRead := context.Param("filetype")

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	var personDetails model.Person
	var allpersonDetails []model.Person

	if fileTypeToRead == "csv" {
		confFilePath := filepath.Dir(dir) + `\server\outputfile.csv`
		csvFile, err := os.Open(confFilePath)
		if err != nil {
			fmt.Println(err)
		}
		defer csvFile.Close()

		reader := csv.NewReader(csvFile)
		reader.FieldsPerRecord = -1

		csvData, err := reader.ReadAll()

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		for _, each := range csvData {
			personDetails.EmpId = each[0]
			personDetails.Name = each[1]
			personDetails.DOB = each[2]
			personDetails.Salary = each[3]
			Age, _ := strconv.ParseInt(each[4], 10, 32)
			personDetails.Age = int32(Age)
			allpersonDetails = append(allpersonDetails, personDetails)
		}
	} else if fileTypeToRead == "xml" {
		confFilePath := filepath.Dir(dir) + `\server\outputfile.xml`
		xmlFile, err := os.Open(confFilePath)
		if err != nil {
			fmt.Println(err)
		}
		defer xmlFile.Close()
		byteValue, _ := ioutil.ReadAll(xmlFile)
		xml.Unmarshal(byteValue, &allpersonDetails)
		fmt.Println(allpersonDetails)
	}
	return json.NewEncoder(context.Response()).Encode(allpersonDetails)
}
