package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func eval(a, b int, ops string) (int, error) {
	switch ops {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf("unsopported operation: %s", ops)
	}
}

func div(a, b int) (q, r int) {
	return a / b, a % b
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args " + "(%d, %d)\n", opName, a, b)
	return op(a, b)
}

func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

/*
func swap(a, b *int){
	*b, *a = *a,  *b
}
 */

func swap(a, b int) (int, int) {
	return b, a
}

func main() {
	if result, err := eval(3, 4, "*"); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}

	fmt.Println(div(13, 3))

	fmt.Println(pow(3, 4))

	//fmt.Println(apply(pow, 3, 4))

	fmt.Println(apply(func(a, b int) int {
		return int(math.Pow(float64(a), float64(b)))
	}, 3, 4))

	fmt.Println(sum(1, 2, 3, 4, 5))

	a, b := 3, 4
	/*
	swap(&a, &b)
	 */
	a, b = swap(a, b)
	fmt.Println(a, b)

}
