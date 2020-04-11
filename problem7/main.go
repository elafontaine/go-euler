package main

func main() {
	primes := []int{2}
	i := 3
	for len(primes) <= 10_000 {
		if is_prime(i, primes) {
			primes = append(primes, i)
		}
		i++
	}
	print(primes[len(primes)-1])
}

func is_prime(i int, primes []int) bool {
	for _, p := range primes {
		if i%p == 0 {
			return false
		}
	}
	return true
}
