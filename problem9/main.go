package main

import (
	"log"
	"math"
)

func main() {
	for b := 2; b <= 500; b++ {
		for a := 1; a < b; a++ {
			c := math.Sqrt(math.Pow(float64(a), 2) + math.Pow(float64(b), 2))
			if c == math.Trunc(c) && a+b+int(c) == 1000 {
				log.Printf("%d, %d, %f \n", a, b, c)
				log.Printf("%d \n", a*b*int(c))

			}
		}
	}
}
