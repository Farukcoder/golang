package main

import "fmt"

func add(x, y float32) float32 {
	return x + y
}

func sub(x, y float32) float32 {
	return x - y
}

func mul(x, y float32) float32 {
	return x * y
}

func dev(x, y float32) float32 {
	return x / y
}

func main() {
	var number1, number2, result float32
	var option string
	i := true

	for i == true {
		fmt.Print("number1: ")
		fmt.Scan((&number1))
		fmt.Print("number2: ")
		fmt.Scan((&number2))

		fmt.Print("choose an option(+, -, *, /, %): ")
		fmt.Scan(&option)

		switch option {
		case "+":
			result = add(number1, number2)
		case "-":
			result = sub(number1, number2)
		case "*":
			result = mul(number1, number2)
		case "/":
			result = dev(number1, number2)
		default:
			fmt.Println("Invalid option")
			continue
		}

		fmt.Printf("Result: %v\n", result)
	}

}
