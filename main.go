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
	caps := flag.Bool("caps", false,"should the string print in caps")
	flag.Parse()
	//check if the user set the flag
	//fmt.Println(*msg)
	//fmt.Println(*num)
	if *caps {
		*msg = strings.ToUpper(*msg)
	}
	
	for i := 0; i < *num; i++ {
		fmt.Println(*msg)
	}
}