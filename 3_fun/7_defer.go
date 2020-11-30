package main

import "fmt"


type Car struct {
  model string
}

// DeLorean DMC-12
func (c Car) PrintModel() {
  fmt.Println(c.model)
}

// 传指针
// Chevrolet Impala
// func (c *Car) PrintModel() {
//   fmt.Println(c.model)
// }

func main() {
  c := Car{model: "DeLorean DMC-12"}
  defer c.PrintModel()
  c.model = "Chevrolet Impala"
}