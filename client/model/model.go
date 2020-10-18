package model

//APIResponse struct defines structure for the responses generated
type APIResponse struct {
	Code    int    `json:"code"`
	Type    string `json:"type"`
	Message string `json:"message"`
}

type Person struct {
	EmpId  string `json:"EmployeeID"`
	Name   string `json:"Name"`
	DOB    string `json:"DOB"`
	Salary string `json:"Salary"`
	Age    int32  `json:"Age"`
}
