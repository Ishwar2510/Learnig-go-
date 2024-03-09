package main 
 import "fmt"
 import  "rsc.io/quote"
 import "example.com/greetings"
 import (
	"log"
 )

func main(){
	fmt.Println("Hello world")
	fmt.Println(quote.Go())
	message,err := (greetings.Hello(""))
	if err != nil {
		log.Fatal(fmt.Sprintf("Error is: %v", err))
		return
	}
	fmt.Println(message)
}