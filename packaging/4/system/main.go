//go mod edit -replace packaging/3/math=../math
package main

import (
	"fmt"
	"packaging/3/math"

	"github.com/google/uuid"
)

func main() {
	m := math.NewMath(1, 2)
	fmt.Println(m)
	fmt.Println(uuid.New().String())
}
