package main

import (
	"fmt"

	"github.com/pkumza/coding/q02/s01/keypress"
)

func main() {
	fmt.Println("N\tPress\tSequence")
	for i := 0; i < 100; i++ {
		press, sequence := keypress.KeyPress(i)
		fmt.Println(i, "\t\033[1;92m", press, "\t", string(sequence), "\033[0m")
	}
}
