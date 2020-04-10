package main

func main() {
	i := 2
	for i > 0 {
		if is_divisible_by_all_(20, i) {
			print(i)
			print("\n")
			break
		}
		i++
	}

}

func is_divisible_by_all_(upToNumber, numberToDivide int) bool {
	for i := 2; i <= upToNumber; i++ {
		if numberToDivide%i != 0 {
			return false
		}
	}
	return true
}
