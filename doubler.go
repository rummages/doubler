package main 

imprt "fmt"

func main() { // This creates the function

	fmt.Print("Enter Your Number: ") 	// Prints enter number  
	var input float64	// creates the var 'input' and sets as floating point 64
	fmt.Scanf("%f", &input)
	
/* Scanf scans text read from standard input, storing successive space-separated values into successive arguments as determined by the format. It returns the number of items successfully scanned. If that is less than the number of arguments, err will report why. Newlines in the input must match newlines in the format. The one exception: the verb %c always scans the next rune in the input, even if it is a space (or tab etc.) or newline. */ 

	output := input * 2 // Creates 'output' with the value of input times 2

	fmt.Println(output) 
} // This closes off the function  	
