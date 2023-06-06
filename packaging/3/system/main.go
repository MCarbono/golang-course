//go mod edit -replace packaging/3/math=../math
package main

import (
	"fmt"
	"packaging/3/math"
)

func main() {
	m := math.NewMath(1, 2)
	fmt.Println(m)
}
