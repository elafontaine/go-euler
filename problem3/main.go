package main

import big "math/big"

func main() {
	print(find_out_prime(600851475143))
}

func find_out_prime(number int) []int {
	prime_numbers := []int{}
	for i := 2; i < number/2; i++ {
		if number%i == 0 {
			if big.NewInt(int64(i)).ProbablyPrime(0) {
				prime_numbers = append(prime_numbers, i)

				print(i)
				if prime_numbers[0]*prime_numbers[len(prime_numbers)-1] >= number {
					return prime_numbers
				}
			}
		}
	}
	return prime_numbers
}
