package age

import (
	"fmt"
)

func Run() {
	fmt.Println("Runing Age package")

	age := 32
	fmt.Println("age: ", age)
	agePointer := &age
	increment(agePointer)
	fmt.Println("increased age", age)

}
func increment(age *int) {
	*age = *age + 1
}
