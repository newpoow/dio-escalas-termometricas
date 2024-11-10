package main

import "fmt"

func exibirMultiplos3() {
	fmt.Println("Os seguintes valores entre 1 e 100 são divisíveis por 3:")
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Printf("%d,", i)
		}
	}
}

func pinPan() {
	for i := 1; i <= 100; i++ {
		multiplo3 := i%3 == 0
		multiplo5 := i%5 == 0

		if multiplo3 && multiplo5 {
			fmt.Print("PinPan,")
		} else if multiplo3 {
			fmt.Print("Pin,")
		} else if multiplo5 {
			fmt.Print("Pan,")
		} else {
			fmt.Printf("%d,", i)
		}
	}
}

func main() {
	fmt.Println("Primeiro desafio")
	fmt.Println("------------")
	exibirMultiplos3()

	fmt.Println()
	fmt.Println()

	fmt.Println("Segundo desafio")
	fmt.Println("------------")
	pinPan()
}
