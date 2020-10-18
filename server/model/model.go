package model

type Person struct {
	EmpId  string `xml:"EmployeeID" json:"EmployeeID"`
	Name   string `xml:"Name" json:"Name"`
	DOB    string `xml:"DOB" json:"DOB"`
	Salary string `xml:"Salary" json:"Salary"`
	Age    int32  `xml:"Age" json:"Age"`
}
