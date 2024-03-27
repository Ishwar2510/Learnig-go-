package main

import "fmt"

func main() {
	// for loop

	// simple for loop

	// for i := 1; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// for as while loop

	// i := 1
	// for ; i < 5; i++ {
	// 	fmt.Println(i)
	// }

	// infinte loop
	// for {
	// 	if i == 20 {
	// 		return
	// 	}
	// 	fmt.Println(i)
	// 	i++
	// }

	// if and else statement and switch stament

	// var day = 2

	// if day == 1 {
	// 	fmt.Print("one")
	// } else {
	// 	fmt.Println("not one")
	// }

	// best way of writing if ans else
	// switch {
	// case day == 1:
	// 	fmt.Print("first day")
	// case day == 2:
	// 	fmt.Println("second day")
	// }

	// using defer to reverse  print  an array
	// defer will use stack to save the function
	// array := []int{
	// 	1, 2, 3, 4, 5, 6, 7,
	// }
	// for _, value := range array {
	// 	fmt.Println(value)
	// 	defer fmt.Println(value)
	// }

	//Pointer
	// & ---> this is refrence i.e pass by refrence
	// * ---> use it as

	// i := 10
	// // to generate pointer of it use &
	// // amphersand is used to create pointer or say address
	// fmt.Println(&i)
	// j := &i
	// fmt.Println(j)
	// // to acces it value use *
	// k := *j
	// fmt.Println(k)

	// struct
	// A struct is a collection of field
	// define an array of type objct having fname and last name

	// type Obj struct {
	// 	fname string
	// 	lname string
	// }

	// emp := []Obj{
	// 	{fname: "ishwar", lname: "Kumar"},
	// 	{fname: "gayatri", lname: "kumar"},
	// }
	// emp = append(emp, Obj{fname: "rakshita", lname: "daughter"})
	// fmt.Println(emp)

	// maps

	// type Vertex struct {
	// 	Lat, Long float64
	// }
	// var m map[string]Vertex

	// fmt.Println(fib(10))
	// fmt.Println(fibCalled)

	// variadic function

	// myFunc := variab(10)
	// fmt.Println(myFunc())
	// fmt.Println(myFunc())
	// fmt.Println(myFunc())
	// fmt.Println(myFunc())
	// fmt.Println(myFunc())

	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}
}

// func variab(initial int) func() int {
// 	num := initial

// 	return func() int {
// 		onum := num
// 		num++
// 		return onum
// 	}
// }
