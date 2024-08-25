package main

import "fmt"

func main(){
	var name string = "John" //explicit declaration & assignment
	var age = 38 // implicit one, compiler infers
	var isAdmin bool = true
	percentage:= 98.8 //shorthand notation

	var interestRate float32 = 10.45
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(isAdmin)
	fmt.Println(percentage)
	fmt.Println("The interest rate is = ",interestRate+10)
}