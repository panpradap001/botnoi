// package main

// import (
// 	"fmt"
// )

// func star(x int) {
// 	// ดาวขาขึ้น
// 	for i := 1; i <= x; i++ {
// 		for j := 1; j <= i; j++ {
// 			fmt.Print("*")
// 		}
// 		fmt.Println()
// 	}
// 	// ดาวขาลง

// 	for i := 1; i <= x-1; i++ {
// 		for j := 1; j <= x-i; j++ {
// 			fmt.Print("*")
// 		}
// 		fmt.Println()

// 	}
// }

// func main() {
// 	var x int
// 	fmt.Scan(&x)
// 	star(x)
// }
