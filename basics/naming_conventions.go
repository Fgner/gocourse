package basics

import "fmt"

// Usa PascalCase para Structs, Interfaces, Enums: exemplo = CalculateArea, UserInfo
type EmployeeGoogle struct {
	FirstName string
	LastName  string
	Age       int
}

// Usa PascalCase para Structs, Interfaces, Enums: exemplo = CalculateArea, UserInfo
type EmployeeApple struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	// snake_case exemplo = user_id, user_name

	// UPPERCASE =  Exclusivamente usada para nomear constantes em Go
	const MAXRETRIES = 5
	// MixedCase = userName, userAge, sumNumbers
	var employeeID = 1001

	fmt.Println("Employee ID: ", employeeID)
}
