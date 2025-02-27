package main

import "fmt"

func star(x int) {
	//ขาขึ้น
	for i := 1; i <= x; i++ {
		fmt.Print("*")
		for j := 1; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	//ขาลง
	for i := 1; i <= x-1; i++ {
		fmt.Print("*")
		for j := 1; j < x-i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

}
func main() {
	var x int
	fmt.Print("จำนวณ : ")
	fmt.Scan(&x)
	star(x)

}
