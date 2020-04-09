package main

func main() {
	acc := 0
	for i := 0; i < 1000; i++ {
		if multiple_of(i, 3) || multiple_of(i, 5) {
			acc += i
		}
	}
	print(acc)
}

func multiple_of(i, x int) bool {
	return i%x == 0
}
