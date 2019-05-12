package main

import (
	"fmt"

	"github.com/pkumza/coding/q02/s01/keypress"
)

func keyPressWithBs(N int) (int, string) {
	if N == 0 {
		return 1, "B"
	}
	type record struct {
		press int
		seq   string
	}
	var records = make([]record, 2*N+1)
	press, seq := keypress.KeyPress(N)

	records[N] = record{press, seq}
	for i := N + 1; i < N+records[N].press; i++ {
		press, seq := keypress.KeyPress(i)

		records[i] = record{press, seq}
		if records[N].press > records[i].press+(i-N) {
			records[N].press = records[i].press + (i - N)
			records[N].seq = records[i].seq
			for l := 0; l < (i - N); l++ {
				records[N].seq += "B"
			}
		}
	}
	return records[N].press, records[N].seq
}

func main() {
	fmt.Println("N\tPress\tSequence")
	for i := 0; i <= 100; i++ {
		press, sequence := keyPressWithBs(i)
		fmt.Println(i, "\t\033[1;92m", press, "\t", string(sequence), "\033[0m")
	}
}
