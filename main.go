// Filename: main.go
// "waitgroups" are just another way to syncronize your goroutines
package main

//demonstrate flags
import (
	"flag"
	"fmt"
	"strings"
)

//"fmt"
//"sync"
//"time"

/*func one (wg *sync.WaitGroup)  {
	defer wg.Done()
	fmt.Println("hola")
}
func two (wg *sync.WaitGroup)  {
	defer wg.Done()
	fmt.Println("ni hao")
}
func three (wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("hello")
}
*/

func main() {

	/*var wg sync.WaitGroup
	wg.Add(3)
	go one(&wg)
	go two(&wg)
	go three(&wg)
	//
	wg.Wait()*/

	//set the flags
	msg := flag.String("msg", "Howdy, stranger!", "the message to display")
	num := flag.Int("num",1 , "how many time to display the message")
	caps := flag.Bool("caps", false,"should the string print in caps")
	flag.Parse()
	//check if the user set the flag
	//fmt.Println(*msg)
	//fmt.Println(*num)
	
	for i := 0; i < *num; i++ {
		fmt.Println(*msg)
	}
}