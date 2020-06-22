package main
import "fmt"

func main() {
	var balance = [] float32 {1000.0, 2.0, 3.4, 7.0, 50.0}

	for i, a := range balance {
		fmt.Printf("index: %d, value: %.2f\n", i, a)
	}
}