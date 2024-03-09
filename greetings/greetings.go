package greetings

import (
	"fmt"
	"errors"
)
func Hello(name string)  (string,error) {
	if name =="" {
		return name,errors.New("empty name")
	}
	message := fmt.Sprintf("hi %v kaise ho", name)
	return message,nil
}
