package main

import "fmt"

type employee struct {
	employeeID   string
	employeeName string
	phone        string
}

func main() {

	employeeList := []employee{}
	employee1 := employee{
		employeeID:   "101",
		employeeName: "yuth",
		phone:        "099999999",
	}
	employee2 := employee{
		employeeID:   "102",
		employeeName: "sura",
		phone:        "099999998",
	}
	employee3 := employee{
		employeeID:   "103",
		employeeName: "nave",
		phone:        "099999997",
	}
	employeeList = append(employeeList, employee1)
	employeeList = append(employeeList, employee2)

	employeeList = append(employeeList, employee3)
	fmt.Println("employee=", employeeList)

}
