package employee

//Employee represents employee info
type Employee struct {
	Name string
	Age  uint8
}

var employees = []Employee{
	Employee{"Josh", 8},
	Employee{"Lyle", 39},
	Employee{"Mike", 51},
	Employee{"Nina", 25},
	Employee{"Sepp", 18},
}

//GetEmployeesOlderThan18 return the array of employees whose age older than 18
func GetEmployeesOlderThan18() []Employee {
	results := []Employee{}
	for _, emp := range employees {
		if emp.Age > 18 {
			results = append(results, emp)
		}
	}
	return results
}

//GetAllEmployees return all employees
func GetAllEmployees() []Employee {
	return employees
}
