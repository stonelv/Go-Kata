package employee

import "testing"

/*
User Story->
As shop owner I want to view a list of all employees, which are older than 18 years, so that I know who is allowed to work on Sundays.
As shop owner I want the list of employees to be sorted by their name, so I can find employees easier.
As shop owner I want the list of employees to be capitalized, so I can read it better.
As shop owner I want the employees to be sorted by their names descending instead of ascending.
*/

func TestShouldReturnEmployeeOlderThan18(t *testing.T) {
	employees := GetEmployeesOlderThan18()
	for _, emp := range employees {
		if emp.Age <= 18 {
			t.Errorf("employee %s (age:%d) is not older than 18", emp.Name, emp.Age)
		}
	}
}

func TestEmployeesShouldSortedByName(t *testing.T) {
	employees := GetAllEmployees()
	var currentChar byte
	for _, emp := range employees {
		preChar := emp.Name[0]
		if currentChar != 0 && preChar < currentChar {
			t.Errorf("Catch not ordered employee, Name:%s.", emp.Name)
		}
		currentChar = preChar
	}
}
