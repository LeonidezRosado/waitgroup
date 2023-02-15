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
	flag.Parse()
	//check if the user set the flag
	//fmt.Println(*msg)
	//fmt.Println(*num)


	//when the user issues a flag reverse the string 
	
	//check if it should be uppercase before printing it  
	if *caps {
		*msg = strings.ToUpper(*msg)
	}

	//print the s string 
	for i := 0; i < *num; i++ {
		fmt.Println(*msg)
	}
}