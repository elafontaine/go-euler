package main

func main() {
	value := fibonacci(1)
	index := 2
	sum := 0

	for value < 4000000 {
		value = fibonacci(index)
		if value%2 == 0 {
			sum += value
		}
		index++
	}
	print(sum)
}

func fibonacci(x int) int {
	if x <= 1 {
		return 1
	}
	return fibonacci(x-1) + fibonacci(x-2)
}
