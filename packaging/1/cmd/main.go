package main

import (
	"fmt"
	"packaging/1/math"
)

func main() {
	m := math.NewMath(1, 2)
	//m := math.Math{}
	fmt.Println(m)
	fmt.Println(m.Add())
}
