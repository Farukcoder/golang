package main

import "fmt"

//function

func displayMessage(countryName string) {
	fmt.Printf("I love my mother land %v \n", countryName)
}

func square(number1, number2 int) {
	result := number1 + number2
	fmt.Println(result)
}

func main() {

	/* print */

	// Println new line every Println
	//fmt.Println("Hello Faruk")
	//fmt.Println("Hello Faruk how are you")

	// Print new line not ad
	//fmt.Print("Hello Faruk")
	//fmt.Print("Hello Faruk how are you")

	/* comment */

	// single line comment

	/*
		multiline
		Comment
	*/

	/* scape sequence */
	//fmt.Println("Hello Faruk")
	//fmt.Println("Hello Faruk \n how are you")
	//
	//fmt.Println("Name \t\t Age")
	//fmt.Println("Omar Faruk \t 25")
	//fmt.Println("Omar Faruk \\ 25")
	//
	//fmt.Println("\"Hello Faruk\"")

	/* data type */

	// "Md Omar Faruk" -> string
	// 25 -> int
	// 15.09 -> float
	// true/false -> boolean

	/* variable */

	//!!! variable naming rules  Allowed!!!//

	// 1. Letters
	// 2. Digits
	// 3. _ (underscore)

	// Not Allowed
	// 1. Keyword
	// 2. variable cannot have space
	// 3. variable cannot start with digit and _ (underscore)
	// 4. variable name case-sensitive

	//variable declaration
	//var fullName, country string
	//var age int
	//var gpa float32

	//variable initialization
	//fullName = "Omar Faruk"
	//country = "Bangladesh"
	//age = 25
	//gpa = 3.25

	//fmt.Println("fullName=", fullName)
	//fmt.Println("country=", country)
	//fmt.Println("age=", age)
	//fmt.Println("gpa=", gpa)
	//
	//fmt.Println("fullName=", fullName, "is", "country=", country, "this", "age=", age, "gpa=", gpa)

	//variable declaration & initialization

	//var newName = "Omar Faruk1"
	//var newCountry = "Bangladesh1"
	//
	//fmt.Println("Name", newName)
	//fmt.Println("Country", newCountry)

	//variable declaration & initialization special

	//lNewName := "Omar Faruk1"
	//lNewCountry := "Bangladesh1"
	//
	//fmt.Println("Name", lNewName)
	//fmt.Println("Country", lNewCountry)

	//formating

	//aNewName := "Omar Faruk1"
	//aNewCountry := "Bangladesh1"
	//
	//fmt.Printf("Name \n %v \n Country \n %v", aNewName, aNewCountry)

	//constant Variable declare

	//fullName := "Faruk"
	//
	//fullName = "nasir"
	//
	//fmt.Println(fullName)

	//get input from user

	//var name, country string
	//var age int
	//var gpa float32

	//get input function
	//Scan, Scanln, Scanf

	//fmt.Print("What is your name? ")
	//fmt.Scan(&name)
	//fmt.Print("What is your age? ")
	//fmt.Scan(&age)
	//fmt.Print("What is your ssc GPA? ")
	//fmt.Scan(&gpa)
	//fmt.Print("What is your Country name? ")
	//fmt.Scan(&country)

	// get input another way

	//fmt.Scanf("%v", &name)

	//get output
	//Print, Println, Printf

	//fmt.Printf("%v is a student \n", name)
	//fmt.Printf("%v  is %v Years Old \n", name, age)
	//fmt.Printf("%v  has got gpa %v in ssc \n", name, gpa)
	//fmt.Printf("%v  is Oreginally from %v \n", name, country)

	// operator +, -, *, /
	//Types of operators
	//1. Arithmetic Operator
	//2. Assignment Operator
	//3. Unary Operator
	//4. Relational Operator
	//5. Logical Operator
	//6. Bitwise Operator
	//7. Special Operator

	//1. Arithmetic Operator -> +, -, *, /, %

	//var num1, num2 int
	//
	//fmt.Printf("num1 =")
	//fmt.Scan(&num1)
	//
	//fmt.Printf("num2 =")
	//fmt.Scan(&num2)
	//
	//result := num1 + num2
	//fmt.Printf("%v + %v = %v \n", num1, num2, result)
	//
	//result = num1 - num2
	//fmt.Printf("%v - %v = %v \n", num1, num2, result)
	//
	//result = num1 * num2
	//fmt.Printf("%v * %v = %v \n", num1, num2, result)
	//
	//var result2 = float32(num1) / float32(num2)
	//fmt.Printf("%v / %v = %.2f \n", num1, num2, result2)
	//
	//result = num1 % num2
	//fmt.Printf("%v %% %v = %v \n", num1, num2, result)

	//area of triangle = 0.5 * base * height

	//var base, height float32
	//
	//fmt.Printf("base: ")
	//fmt.Scan(&base)
	//
	//fmt.Printf("height: ")
	//fmt.Scan(&height)
	//
	//result := 0.5 * base * height
	//
	//fmt.Printf("area of Triangle: %v and %v = %.2f\n", base, height, result)

	//area of circle = 3.1416 * radius * radius

	//var radius float32
	//
	//fmt.Printf("radius:")
	//fmt.Scan(&radius)

	//result := math.Sqrt(float64(radius))

	//result := 3.1416 * radius * radius
	//
	//fmt.Printf("area of circle for %v = %v", radius, result)

	// 2. Assignment Operator -> =, +=, -=, *=, /=, %=

	//x := 10
	//x += 5
	//fmt.Println(x)
	//
	//xy := 10
	//xy -= 5
	//fmt.Println(xy)

	//3. Unary Operator -> ++, --

	//x := 10
	//
	//x++
	//x--
	//fmt.Printf("Result =  %v", x)

	//4. Relational Operator // Comparison Operator

	//x := 5
	//y := 6
	//
	//fmt.Printf("x = %v \n", x > y)
	//fmt.Printf("x = %v \n", x < y)
	//fmt.Printf("x = %v \n", x <= y)
	//fmt.Printf("x = %v \n", x > y)
	//fmt.Printf("x = %v \n", x >= y)

	//5. Logical Operator -> &&, ||,!

	//x := 3
	//y := 4
	//z := x > 3 && y > 4
	//z = x > 2 || y > 4
	//z = !(x > 2 || y > 4)
	//fmt.Printf("%v", z)

	//6. Bitwise Operator -> &, |, ^, >> <<
	//
	//x := 18       //18  = 10010
	//y := 17       // 17 = 10001
	//and := x & y  //    = 10000 = 16
	//or := x | y   //    = 10011 = 19
	//exor := x ^ y //    = 0001 = 3
	//
	//fmt.Printf("x & y: %v \n", and)
	//fmt.Printf("x | y: %v \n", or)
	//fmt.Printf("x ^ y: %v \n", exor)

	// Conditional control stetmant
	//conditional -> if, elseif, else
	//var number int
	//fmt.Printf("Enter a number: ")
	//fmt.Scanf("%d", &number)
	//
	//if number > 0 {
	//	fmt.Printf("Positive")
	//} else if number < 0 {
	//	fmt.Printf("Negative")
	//} else {
	//	fmt.Printf("Zero")
	//}

	//a number even or odd checker
	//var number int
	//
	//fmt.Printf("Enter a number: ")
	//
	//fmt.Scanf("%d", &number)
	//
	//if number%2 == 0 {
	//	fmt.Println("The number is even")
	//} else {
	//	fmt.Println("The number is odd")
	//}

	//find large number between 2 numbers
	//var number1, number2 int
	//
	//fmt.Printf("Enter tow number: ")
	//fmt.Scan(&number1, &number2)
	//if number1 > number2 {
	//	fmt.Printf("Larger number is %v \n", number1)
	//} else if number1 < number2 {
	//	fmt.Printf("Larger number is %v \n", number2)
	//} else {
	//	fmt.Printf("Numbers are equal")
	//}

	//find large number among 3 numbers
	//var number1, number2, number3 int
	//
	//fmt.Printf("Enter 3 numbers: ")
	//fmt.Scan(&number1, &number2, &number3)

	//if number1 > number2 {
	//	if number1 > number3 {
	//		fmt.Printf("Large number is %v", number1)
	//	} else {
	//		fmt.Printf("Larger  number is %v", number3)
	//	}
	//} else if number2 > number1 {
	//	if number2 > number3 {
	//		fmt.Printf("Larger Number is %v", number2)
	//	} else {
	//		fmt.Printf("Larger Number is %v", number3)
	//	}
	//}

	//find large number using logical operators

	//if number1 < number2 && number1 < number3 {
	//	fmt.Printf("Large Number is %v \n", number3)
	//} else if number2 > number1 && number2 > number3 {
	//	fmt.Printf("Large Number is %v \n", number2)
	//} else if number3 < number1 && number3 > number2 {
	//	fmt.Printf("Large Number is %v \n", number1)
	//} else {
	//	fmt.Printf("Numbers Are equal \n")
	//}

	//digit spelling program using if, else

	//var digit int
	//fmt.Printf("Enter a Digit from 0-9 : ")
	//fmt.Scan(&digit)

	//if digit == 0 {
	//	fmt.Println("Zero")
	//} else if digit == 1 {
	//	fmt.Println("One")
	//} else if digit == 2 {
	//	fmt.Println("Two")
	//} else if digit == 3 {
	//	fmt.Println("Three")
	//} else if digit == 4 {
	//	fmt.Println("Four")
	//} else if digit == 5 {
	//	fmt.Println("Five")
	//} else if digit == 6 {
	//	fmt.Println("Six")
	//} else if digit == 7 {
	//	fmt.Println("Seven")
	//} else if digit == 8 {
	//	fmt.Println("Eight")
	//} else if digit == 9 {
	//	fmt.Println("Nine")
	//} else {
	//	fmt.Println("Invalid Digit")
	//}

	//switch catch

	//switch digit {
	//	case 0:
	//		fmt.Println("Zero")
	//	case 1:
	//		fmt.Println("One")
	//	case 2:
	//		fmt.Println("Two")
	//	case 3:
	//		fmt.Println("Three")
	//	case 4:
	//		fmt.Println("Four")
	//	case 5:
	//		fmt.Println("Five")
	//	case 6:
	//		fmt.Println("Six")
	//	case 7:
	//		fmt.Println("Seven")
	//	case 8:
	//		fmt.Println("Eight")
	//	case 9:
	//		fmt.Println("Nine")
	//	case 10:
	//		fmt.Println("Ten")
	//	default:
	//		fmt.Println("Invalid Digit")
	//}

	//multiple catch
	//switch digit {
	//case 0, 1, 2, 3, 4, 5:
	//	fmt.Println("Primary")
	//case 6, 7, 8, 9, 10:
	//	fmt.Println("Secondary")
	//case 11, 12:
	//	fmt.Println("Hsc")
	//default:
	//	fmt.Println("Invalid")
	//}

	//loop control statement

	//for i := 1; i < 6; i++ {
	//	fmt.Println("I love Bangladesh")
	//}

	//for i := 10; i <= 100; i = i + 5 {
	//	fmt.Println(i)
	//}

	//i := 100
	//
	//for i >= 1 {
	//	fmt.Printf("%v is less than 100\n", i)
	//	i--
	//}

	// print even number from 1 to 100, 2, 4, 6

	//for x := 2; x <= 100; x = x + 2 {
	//	fmt.Println(x)
	//}

	//for x := 1; x <= 100; x++ {
	//	if x%2 == 0 {
	//		fmt.Printf("%d is even\n", x)
	//	}
	//}

	//series program | 2 + 4 + 6 + ... + n

	//series program | 1 + 2 + 3 + ... + 10

	//var startNumber, endNumber, distanceNumber int
	//fmt.Printf("Enter the starting number of series: ")
	//fmt.Scan(&startNumber)
	//fmt.Printf("Enter the ending number of series: ")
	//fmt.Scan(&endNumber)
	//fmt.Printf("Enter the distance number of series: ")
	//fmt.Scan(&distanceNumber)
	//sum := 0
	//for i := startNumber; i <= endNumber; i = i + distanceNumber {
	//	sum += i
	//	fmt.Printf("%v+", i)
	//}
	//fmt.Printf("The sum of series is %v \n", sum)

	// continue

	//for i := 1; i <= 100; i++ {
	//	if i%2 == 0 {
	//		continue
	//	}
	//	fmt.Println(i)
	//}

	//break
	//for i := 1; i <= 100; i++ {
	//	if i%3 == 0 {
	//		break
	//	}
	//	fmt.Println(i)
	//}

	//function
	//displayMessage("Bangladesh")
	//square(2, 2)

	// pointer

	//var p *int // pointer variable store another variable memory address
	//x := 1
	//p = &x
	//*p -= 1
	//fmt.Println("x:", &x)
	//fmt.Println("x:", x)
	//fmt.Println("p:", p)
	//fmt.Println("p:", *p)

	//x := 10
	//chaneValue(x)
	//fmt.Println(x)
	//
	//callByRef(&x)
	//fmt.Println(x)

	//structure
	//Rahim -> id, name, age
	//Karim -> id, name, age

	//class Student{
	//	id
	//	name
	//	age
	//}
	//
	//rahim is a student

	rahim := Student{100, "Bangladeshi", 25}
	karim := Student{101, "Bangladeshis", 31}

	rahim.increasesAge(6)

	fmt.Println("Rahim Details")
	displayInfo(rahim)
	fmt.Println("Rahim Details")
	displayInfo(karim)

}

func displayInfo(s Student) {
	//fmt.Println(s.id)
	//fmt.Println(s.address)
	//fmt.Println(s.age)

	fmt.Printf("Id: %v \n Address: %v \n Age: %v \n", s.id, s.address, s.age)
}

type Student struct {
	id      int
	address string
	age     int
}

func (x *Student) increasesAge(val int) {
	x.age += val
}

//func chaneValue(x int) {
//	x = 25
//}
//
//func callByRef(val *int) {
//	*val = 25
//}
