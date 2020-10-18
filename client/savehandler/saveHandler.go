package savehandler

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo"

	model "Benz-assignment/client/model"
	pb "Benz-assignment/proto"

	"encoding/json"

	"google.golang.org/grpc"
)

const (
	address = "localhost:8080"
)

func randGenerator() string {
	n := 5
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		fmt.Println(err)
	}
	s := fmt.Sprintf("%X", b)
	return s
}

func SaveHandler(cntx echo.Context) error {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	cntx.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewStoreDataClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	defer cntx.Request().Body.Close()

	personModel := model.Person{}
	errDecode := json.NewDecoder(cntx.Request().Body).Decode(&personModel)
	if errDecode != nil {
		ResponseMapper(http.StatusMethodNotAllowed, "error in decoding json", cntx)
	}
	fileType := cntx.Param("filetype")
	if fileType != "CSV" {
		log.Println(fileType)
		if fileType != "XML" {
			return ResponseMapper(http.StatusMethodNotAllowed, "Invalid file type", cntx)
		}
	}

	personModel.EmpId = randGenerator()

	r, err := c.Add(ctx, &pb.Request{Empid: personModel.EmpId, Name: personModel.Name, Dob: personModel.DOB, Salary: personModel.Salary, Age: personModel.Age, Filetype: fileType})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetResponse())
	return json.NewEncoder(cntx.Response()).Encode(r.GetResponse())
}

//ResponseMapper to handle response
func ResponseMapper(code int, message string, c echo.Context) error {

	response := model.APIResponse{}

	response.Code = code
	response.Type = http.StatusText(code)
	response.Message = message

	c.Response().WriteHeader(code)
	return json.NewEncoder(c.Response()).Encode(response)
}
