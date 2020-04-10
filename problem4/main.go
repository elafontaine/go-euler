package main

import (
	"math"
)

var digits = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

func main() {
	previous := 0
	for number1 := range forXDigits(3) {
		for number2 := range forXDigits(3) {
			if number1 > 0 && number2 > 0 {
				multiplication := number1 * number2
				if is_palindrome(multiplication) && multiplication > previous {
					print(multiplication)
					print("\n")
					previous = multiplication
				}
			}
		}
	}
}
func is_palindrome(number int) bool {
	numbers := []int{}
	for i := 1; i <= number; i *= 10 {
		digit := number / i % 10
		numbers = append(numbers, digit)
	}
	for i, _ := range numbers {
		if numbers[i] != numbers[len(numbers)-1-i] {
			return false
		}
	}
	return true
}

func forXDigits(number int) <-chan int {
	left := digitsGenerator()
	right := daisyChain(left, nextDigit, number-1)
	return right
}

func nextDigit(leftChannel chan int, rightChannel chan int, index int) {
	rightChannel2 := daisyChain(leftChannel, write_next_number(index), 1)
	for value := range rightChannel2 {
		rightChannel <- value
	}
	close(rightChannel)
}

func write_next_number(index int) func(left_channel, rightChannel chan int, repetition int) {
	return func(left_channel, rightChannel chan int, repetition int) {
		for inputNumber := range left_channel {
			digitChannel := digitsGenerator()
			for digit := range digitChannel {
				rightChannel <- digit*int(math.Pow10(index+1)) + inputNumber
			}
		}
		close(rightChannel)
	}
}

func daisyChain(chan1 chan int, chainFunc func(chan int, chan int, int), repetition int) chan int {
	left := chan1
	right := chan1
	for i := 0; i < repetition; i++ {
		right = make(chan int)
		go chainFunc(left, right, i)
		left = right
	}
	return right
}
func digitsGenerator() chan int {
	yield := make(chan int)
	go func() {
		for _, v := range digits {
			yield <- v
		}
		close(yield)
	}()
	return yield
}
