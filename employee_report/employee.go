package employee

//Employee represents employee info
type Employee struct {
	Name string
	Age  uint8
}

var employees = []Employee{
	{"Josh", 8},
	{"Lyle", 39},
	{"Mike", 51},
	{"Nina", 25},
	{"Sepp", 18},
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

//GetAllEmployeesInDescend return all employees sorted by names in descending
func GetAllEmployeesInDescend() []Employee {
	return SortEmployeeInDescend(employees)
}

//SortEmployeeInDescend used to sort the employee array by name in descending
func SortEmployeeInDescend(emps []Employee) []Employee {
	return QuickSort(emps, 0, len(emps)-1)
}

//QuickSort used to do the sorting with quick sort method
func QuickSort(emps []Employee, left, right int) []Employee {
	if left < right {
		partitionIndex := Partition(emps, left, right)
		QuickSort(emps, left, partitionIndex-1)
		QuickSort(emps, partitionIndex+1, right)
	}
	return emps
}

//Partition is a helper method, used to find the middle index of the array.
func Partition(emps []Employee, left, right int) int {
	pivot := left
	index := pivot + 1

	for i := index; i <= right; i++ {
		if emps[i].Name[0] > emps[pivot].Name[0] {
			Swap(emps, i, index)
			index++
		}
	}

	Swap(emps, pivot, index-1)
	return index - 1
}

//Swap used to swap the two items by index
func Swap(emps []Employee, i, j int) {
	emps[i], emps[j] = emps[j], emps[i]
}
