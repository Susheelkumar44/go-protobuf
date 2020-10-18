package main

import (
	model "Benz-assignment/server/model"
	"context"
	"encoding/csv"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strconv"

	pb "Benz-assignment/proto"

	"google.golang.org/grpc"
)

const port = ":8080"

type server struct {
	pb.UnimplementedStoreDataServer
}

func (s *server) Add(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	log.Println("calling server")
	personModel := model.Person{}
	personModel.EmpId = in.GetEmpid()
	personModel.Name = in.GetName()
	personModel.DOB = in.GetDob()
	personModel.Salary = in.GetSalary()
	personModel.Age = in.GetAge()

	jsonModel, _ := json.Marshal(personModel)
	fmt.Println(jsonModel)

	fileType := in.GetFiletype()
	if fileType == "CSV" {
		var d interface{}
		err := json.Unmarshal(jsonModel, &d)
		if err != nil {
			log.Fatal("Failed to unmarshal")
		}
		values := decodeJson(d.(map[string]interface{}))
		fmt.Println(values)

		f, err := os.OpenFile("outputfile.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal("Failed to create outputfile.csv")
			return &pb.Response{Response: "Failed to create outputfile.csv"}, nil
		}
		defer f.Close()
		w := csv.NewWriter(f)
		errWrite := w.Write(values)
		if errWrite != nil {
			log.Fatal("Failed to write to file")
			return &pb.Response{Response: "Failed to write to file"}, nil
		}
		w.Flush()
		if err := w.Error(); err != nil {
			log.Fatal("Failed to flush outputfile.csv")
			return &pb.Response{Response: "Failed to flush outputfile.csv"}, nil
		}

	} else if fileType == "XML" {
		err := json.Unmarshal(jsonModel, &personModel)
		if err != nil {
			log.Fatal("Failed to unmarshal")
		}
		out, err := xml.Marshal(personModel)
		f, err := os.OpenFile("outputfile.xml", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal("Failed to create outputfile.xml")
			return &pb.Response{Response: "Failed to create outputfile.xml"}, nil
		}
		defer f.Close()
		err = ioutil.WriteFile("outputfile.xml", out, 0644)
		if err != nil {
			return &pb.Response{Response: "Failed to write to outputfile.xml"}, nil
		}
	}
	return &pb.Response{Response: "Successfully Stored on file"}, nil
}

func main() {
	log.Println("calling server")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterStoreDataServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func decodeJson(m map[string]interface{}) []string {
	values := make([]string, 0, len(m))
	for _, v := range m {
		switch vv := v.(type) {
		case map[string]interface{}:
			for _, value := range decodeJson(vv) {
				values = append(values, value)
			}
		case string:
			values = append(values, vv)
		case float64:
			values = append(values, strconv.FormatFloat(vv, 'f', -1, 64))
		case []interface{}:
		case bool:
			values = append(values, strconv.FormatBool(vv))
		case nil:
			values = append(values, "nil")
		}
	}
	return values
}
