package main

import (
	"fmt"
)

func main() {
	kelvin := 373
	celsius := kelvin - 273

	fmt.Printf("O ponto de ebulição da água na escala Kelvin é %d K, ", kelvin)
	fmt.Printf("o que corresponde a %d° C na escala Celsius", celsius)
}
