package main

import (
	"fmt"
	"math"
)

func main() {
	n := 5
	x := float64(n/2)
	if n % 2 == 0 {
		fmt.Println("Angka Harus Ganjil");
	} else {
		for i := 0; i < n; i++ {
			str := ""

			for j := 0; j < n; j++ {
				if j == 0 || j == n - 1 {
					str += "* ";
				} else if float64(i) == math.Floor(x) {
					str += "* ";
				} else {
					str += "= ";
				}
			}
			fmt.Println(str)
		}
	}
}