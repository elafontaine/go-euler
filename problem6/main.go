package main

func main() {
	acc1, acc2 := 0, 0

	for i := 0; i <= 100; i++ {
		acc1 += i * i
		acc2 += i
	}
	print(acc2*acc2 - acc1)
}
