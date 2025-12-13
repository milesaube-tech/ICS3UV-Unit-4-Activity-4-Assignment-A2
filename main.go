/**
 * Author: Miles Aube
 * Version: 1.0.0
 * Date: 2025-12-13
 * Fileoverview: program to create calculator
 */


package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	//create buffer reader
	reader := bufio.NewReader(os.Stdin)

	// variables
	var userOperations string

	// get input for operation 
	fmt.Print("Welcome to my calculator program.\nwhich operation would you like to perform today? (select by typing the letter in the front of the operation.)\na.add\nb.subtract\nc.multiply\nd.divide\ne.absolute value\nf.round\ng.raise to an exponent\nh.square root\n")
  userOperations, _ = reader.ReadString('\n')
	userOperations = strings.TrimSpace(userOperations)

	// if statements for each operation
	if userOperations == "a"{
		var userValue1AsNumber int
		var userValue2AsNumber int
		var userValue1AsString string
		var userValue2AsString string
		var answer int
    fmt.Println("Please enter your first number")
		userValue1AsString, _ = reader.ReadString('\n')
		userValue1AsString = strings.TrimSpace(userValue1AsString)
		userValue1AsNumber, _ = strconv.Atoi(userValue1AsString)

    fmt.Println("Please enter your second number")
		userValue2AsString, _ = reader.ReadString('\n')
		userValue2AsString = strings.TrimSpace(userValue2AsString)
		userValue2AsNumber, _ = strconv.Atoi(userValue2AsString)
    
		answer = userValue1AsNumber + userValue2AsNumber 
		fmt.Println("The sum of ",userValue1AsNumber," and ",userValue2AsNumber, " is ",answer)
	}

	if userOperations == "b"{
		var userValue1AsNumber int
		var userValue2AsNumber int
		var userValue1AsString string
		var userValue2AsString string
		var answer int
    fmt.Println("Please enter your first number")
		userValue1AsString, _ = reader.ReadString('\n')
		userValue1AsString = strings.TrimSpace(userValue1AsString)
		userValue1AsNumber, _ = strconv.Atoi(userValue1AsString)

    fmt.Println("Please enter your second number")
		userValue2AsString, _ = reader.ReadString('\n')
		userValue2AsString = strings.TrimSpace(userValue2AsString)
		userValue2AsNumber, _ = strconv.Atoi(userValue2AsString)
    
		answer = userValue1AsNumber - userValue2AsNumber 
		fmt.Println("The difference of ",userValue1AsNumber," and ",userValue2AsNumber, " is ",answer)
	}

	if userOperations == "c"{
		var userValue1AsNumber int
		var userValue2AsNumber int
		var userValue1AsString string
		var userValue2AsString string
		var answer int
    fmt.Println("Please enter your first number")
		userValue1AsString, _ = reader.ReadString('\n')
		userValue1AsString = strings.TrimSpace(userValue1AsString)
		userValue1AsNumber, _ = strconv.Atoi(userValue1AsString)

    fmt.Println("Please enter your second number")
		userValue2AsString, _ = reader.ReadString('\n')
		userValue2AsString = strings.TrimSpace(userValue2AsString)
		userValue2AsNumber, _ = strconv.Atoi(userValue2AsString)
    
		answer = userValue1AsNumber * userValue2AsNumber 
		fmt.Println("The product of ",userValue1AsNumber," and ",userValue2AsNumber, " is ",answer)
	}
  
	if userOperations == "d"{
		var userValue1AsNumber int
		var userValue2AsNumber int
		var userValue1AsString string
		var userValue2AsString string
		var answer int
    fmt.Println("Please enter your first number")
		userValue1AsString, _ = reader.ReadString('\n')
		userValue1AsString = strings.TrimSpace(userValue1AsString)
		userValue1AsNumber, _ = strconv.Atoi(userValue1AsString)

    fmt.Println("Please enter your second number")
		userValue2AsString, _ = reader.ReadString('\n')
		userValue2AsString = strings.TrimSpace(userValue2AsString)
		userValue2AsNumber, _ = strconv.Atoi(userValue2AsString)
    
		answer = userValue1AsNumber / userValue2AsNumber 
		fmt.Println("The quotient of ",userValue1AsNumber," and ",userValue2AsNumber, " is ",answer)
	}
  
	if userOperations == "e"{
		var userValue1AsNumber float64
		var userValue1AsString string
		var answer float64
    fmt.Println("Please enter the number you would like to find the absolute for")
		userValue1AsString, _ = reader.ReadString('\n')
		userValue1AsString = strings.TrimSpace(userValue1AsString)
		userValue1AsNumber, _ = strconv.ParseFloat(userValue1AsString,64)
    
		answer = math.Abs(userValue1AsNumber) 
		fmt.Println("The absolute value of ",userValue1AsNumber," is ",answer)
	}

	if userOperations == "f"{
		var userValue1AsNumber float64
		var userValue1AsString string
		var answer float64
    fmt.Println("Please enter the number you would like to round ")
		userValue1AsString, _ = reader.ReadString('\n')
		userValue1AsString = strings.TrimSpace(userValue1AsString)
		userValue1AsNumber, _ = strconv.ParseFloat(userValue1AsString,64)
    
		answer = math.Round(userValue1AsNumber) 
		fmt.Println("The rounded value of ",userValue1AsNumber," is ",answer)
	}

	if userOperations == "g"{
		var userValue1AsNumber float64
		var userValue2AsNumber float64
		var userValue1AsString string
		var userValue2AsString string
		var answer float64
    fmt.Println("Please enter your base number")
		userValue1AsString, _ = reader.ReadString('\n')
		userValue1AsString = strings.TrimSpace(userValue1AsString)
		userValue1AsNumber, _ = strconv.ParseFloat(userValue1AsString,64)

    fmt.Println("Please enter your exponent number")
		userValue2AsString, _ = reader.ReadString('\n')
		userValue2AsString = strings.TrimSpace(userValue2AsString)
		userValue2AsNumber, _ = strconv.ParseFloat(userValue2AsString,64)
    
		answer = math.Pow(userValue1AsNumber,userValue2AsNumber) 
		fmt.Println("The power of ",userValue1AsNumber," to ",userValue2AsNumber, " is ",answer)
	}

	if userOperations == "h"{
		var userValue1AsNumber float64
		var userValue1AsString string
		var answer float64
    fmt.Println("Please enter the number you would like to find the square root of ")
		userValue1AsString, _ = reader.ReadString('\n')
		userValue1AsString = strings.TrimSpace(userValue1AsString)
		userValue1AsNumber, _ = strconv.ParseFloat(userValue1AsString,64)
    
		answer = math.Sqrt(userValue1AsNumber) 
		fmt.Println("The square root  of ",userValue1AsNumber," is ",answer)
	}
	fmt.Println("\nDone.")
}