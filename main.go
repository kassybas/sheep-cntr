package main

import "fmt"
import "os"
import "strconv"
import "math"

func main() {
	var i uint64 = 1
	x, err := strconv.Atoi(os.Args[1])
	var limit uint64 = uint64(math.Pow10(x))
	if err != nil {
		fmt.Println(err, "Parmater is not a number")
	}
	for i < limit+1 {
		fmt.Printf("%d sheep\n", i)
		i++
	}
}
