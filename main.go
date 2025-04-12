// Filename: main.go
// "waitgroups" are just another way to syncronize your goroutines
package main

//demonstrate flags
import (
	"flag"
	"fmt"
	"strings"
)


func main() {


	//set the flags
	msg := flag.String("msg", "Howdy, stranger!", "the message to display")
	num := flag.Int("num",1 , "how many time to display the message")
	caps := flag.Bool("U", false,"specific whether to uppercase the string(true or false)")
	reverse := flag.Bool("r", false, "reverse the string (true or false)")
	flag.Parse()
	//check if the user set the flag
	//fmt.Println(*msg)
	//fmt.Println(*num)

	//check if it should be uppercase before printing it  
	if *caps {
		*msg = strings.ToUpper(*msg)
	}
	//when the user issues a flag reverse the string 
	//check if we should reverse the string 
	if *reverse {
		//create and empty string 
		//2.iterate over the string that you want to reverse
		var result string 
		for _, value := range *msg {
			result = string(value) + result
	
		}
		//write the reverse string to *msg
		*msg = result;
	}

	//print the s string 
	for i := 0; i < *num; i++ {
		fmt.Println(*msg)
	}
}
///how do i run the go file in vscode
//what is going on now