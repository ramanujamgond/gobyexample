// for loop in go

package main

import "fmt"

func main() {
	i := 1

	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	k := 1

	for k <= 5 {
		fmt.Println("loop")
		k++
	}

	for {
		fmt.Println("Loop Data")
		break
	}

	for n := 0; n <= 25; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
