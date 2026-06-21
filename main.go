package main

import (
	"errors"
	"fmt"
)

// Pi represents a foundational math constant.
const Pi = 3.14159

// Person defines a custom data type grouping a person's attributes.
type Person struct {
	Name string
	Age  int
}

// Greet returns a welcoming introduction string from the receiver.
func (p Person) Greet() string {
	return fmt.Sprintf("Hi, my name is %s and I am %d years old.", p.Name, p.Age)
}

func main() {
	// ==========================================
	// 1. Variables & Data Types
	// ==========================================
	var message string = "Learning Go is fun!"

	score := 95      // Type inferred as int
	isPassed := true // Type inferred as bool

	fmt.Println(message)
	fmt.Printf("Score: %d, Passed: %t\n", score, isPassed)
	fmt.Println("--------------------------------")

	// ==========================================
	// 2. Conditional Statements (if/else)
	// ==========================================
	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 70 {
		fmt.Println("Grade: B")
	} else {
		fmt.Println("Grade: C")
	}
	fmt.Println("--------------------------------")

	// ==========================================
	// 3. Arrays & Slices
	// ==========================================
	// Arrays have a fixed, unchangeable size
	var fixedArray [3]int = [3]int{10, 20, 30}

	// Slices are dynamic wrappers around arrays (highly preferred in Go)
	dynamicSlice := []string{"Apple", "Banana", "Cherry"}
	dynamicSlice = append(dynamicSlice, "Dragonfruit")

	fmt.Println("Array:", fixedArray)
	fmt.Println("Slice:", dynamicSlice)
	fmt.Println("--------------------------------")

	// ==========================================
	// 4. Loops (Go relies exclusively on 'for')
	// ==========================================
	fmt.Print("Standard loop counter: ")
	for i := 0; i < 3; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// Range loops are perfect for iterating over collections
	fmt.Println("Iterating over slice:")
	for index, value := range dynamicSlice {
		fmt.Printf("  Index %d: %s\n", index, value)
	}
	fmt.Println("--------------------------------")

	// ==========================================
	// 5. Maps (Key-Value associations)
	// ==========================================
	userRoles := make(map[string]string)
	userRoles["Alice"] = "Admin"
	userRoles["Bob"] = "Developer"

	fmt.Println("Bob's role is:", userRoles["Bob"])
	fmt.Println("--------------------------------")

	// ==========================================
	// 6. Explicit Error Handling & Multiple Returns
	// ==========================================
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Division Result:", result)
	}

	// Handling explicit errors cleanly
	_, err = divide(10, 0)
	if err != nil {
		fmt.Println("Caught expected error:", err)
	}
	fmt.Println("--------------------------------")

	// ==========================================
	// 7. Struct Instantiation & Method Invocation
	// ==========================================
	student := Person{Name: "Alex", Age: 21}
	fmt.Println(student.Greet())
}

// divide calculates the quotient of two values or returns an error on division by zero.
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}