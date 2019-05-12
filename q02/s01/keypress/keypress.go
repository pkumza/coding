package keypress

// KeyPress accepts target number of repeat times on the screen,
// 	returns times of key pressing & key pressing sequence
func KeyPress(N int) (int, string) {
	var seq string
	if N <= 0 {
		return -1, seq
	}
	var result = 0
	for i := 2; i <= N; {
		if N%i == 0 {
			N = N / i
			result += i
			seq = seq + "C"
			for j := 1; j < i; j++ {
				seq = seq + "P"
			}
		} else {
			i++
		}
	}
	return result, seq
}
