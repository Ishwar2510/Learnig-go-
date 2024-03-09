package main

// creating generic and using it
// with generic u can declare and use fucntions or type that are written to work
// witha ny of a set of types

// writing non-generic and then converting it to generic
// waf to add two values fro the map
import (
	"fmt"
)

func SumInt(m map[string]int64) int64 {
	var sum int64
	for _, v := range m {
		sum += v
	}
	return sum
}

func SumFloat(m map[string]float64) float64 {
	var sum float64
	for _, v := range m {
		sum += v
	}
	return sum
}

func main() {
	int_map := map[string]int64{
		"first":  1,
		"second": 2,
	}
	float_map := map[string]float64{
		"first":  1,
		"second": 2,
	}

	fmt.Printf("The ans are %v, %v", SumInt(int_map), SumFloat(float_map))

	fmt.Printf("The ans are %v, %v", Generic(int_map), Generic(float_map))

	fmt.Printf("The ans are %v, %v", Generic2(int_map), Generic2(float_map))

}

// now we are going to replace the above with generic function
// with type parameter inside a square box
// we receive arguments and we pass parameters

func Generic[K comparable, V int64 | float64](m map[K]V) V {

	var sum V
	for _, value := range m {
		sum += value
	}
	return sum

}

// declarug type constraints as interface

type Number interface {
	int64 | float64
}

func Generic2[K comparable, V Number](m map[K]V) V {
	var sum V
	for _, value := range m {
		sum += value
	}
	return sum

}
