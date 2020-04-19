package main

import "fmt"

func main() {
	primes := []int{2}
	sum := 2
	for i := 3; i <= 2_000_000; i++ {
		if isPrime(primes, i) {
			primes = append(primes, i)
			sum += i
		}
	}
	fmt.Printf("%d", sum)

}

func isPrime(primes []int, i int) bool {
	for _, y := range primes {
		if i%y == 0 {
			return false
		}
	}
	return true
}
