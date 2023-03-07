package main

import "fmt"



func addNumber(x int, y int)int {
	return x +y
	}

func subNumber(x int, y int)int {
		return x - y
		}
	

func main(){
	fmt.Println("Hello, World!")


	fmt.Println("This is my HB Assement assignment 4")
	
	
	result:= addNumber(5, 10)
	fmt.Printf("I created a func that adds 2 numbers together: %v\n", result)

	sum := subNumber(5, 10)
	fmt.Printf("I created a func that subtracts 2 numbers from each other: %v\n", sum)
}

