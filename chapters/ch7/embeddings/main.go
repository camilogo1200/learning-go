package main

import "fmt"

type Employee struct {
	Name string
	Id   string
}

func (e Employee) Description() string {
	return fmt.Sprintf("%s  (%s)", e.Name, e.Id)
}

type Manager struct {
	Employee
	Reports []Employee
}

func (m Manager) FindNewEmployee() []Employee {
	return m.Reports
}

func main() {

	m := Manager{
		Employee: Employee{
			Name: "Bob Bobson",
			Id:   "1234",
		},
		Reports: []Employee{},
	}
	fmt.Println(m.Id)
	fmt.Println(m.Description())

	//Embedding is not inheritance
	//var eFail Employee = m
	var eOk Employee = m.Employee

	//fmt.Println(eFail)
	fmt.Println(eOk)
}
