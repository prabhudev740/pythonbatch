// // package main

// // import (
// // 	"fmt"
// // 	// "math"
// // 	// "sort"
// // 	"strings"
// // )

// // func main() {
// // 	// x := "1.1.1.1"
// // 	// // x = strings.ReplaceAll(x, ".", "[.]")
// // 	// x := []string{"please wait", "continue to fight", "continue to win"}
// // 	// res := fimdMax(x)
// // 	// fmt.Println(res)

// // 	// temp()
// // 	pattern2()
// // }

// // func sol(address string) string {
// // 	var res string

// // 	for _, v := range address {
// // 		if v == '.' {
// // 			res += "[.]"
// // 		} else {
// // 			res += string(v)
// // 		}
// // 	}

// // 	return res
// // }

// // func fimdMax(sentences []string) int {
// // 	res := 0

// // 	for _, v := range sentences {
// // 		c := len(strings.Split(v, " "))
// // 		if res < c {
// // 			res = c
// // 		}
// // 	}

// // 	return res
// // }

// // func temp() {
// // 	// res := strings.Join([]string{"a", "bc"}, "")
// // 	// res := []int{}
// // 	// strings.Contains()
// // 	// fmt.Println(m, p, g)

// // 	// if 10 && 10 {
// // 	// 	fmt.Println("abv")
// // 	// }
// // 	x := 10
// // 	y := 10
// // 	z := y
// // 	fmt.Println(&x, &y, &z)
// // }

// package main

// import (
// 	"fmt"
// 	"math"
// 	// "strconv"
// )

// func pattern2() {
// 	n := 5
// 	for i := 1; i <= n; i++ {
// 		for j := n - i; j >= 1; j-- {
// 			fmt.Printf(" ")
// 		}
// 		for j := i; j >= 1; j-- {
// 			fmt.Printf("%d", j)
// 		}

// 		fmt.Println("")
// 	}
// }

// func pattern() {
// 	n := 5
// 	for i := 1; i <= n; i++ {
// 		for j := n - i; j >= 1; j-- {
// 			fmt.Printf(" ")
// 		}

// 		for j := n - i + 1; j <= n; j++ {
// 			fmt.Printf("%d", j)
// 		}

// 		fmt.Println("")
// 	}
// }

// func main() {
// 	// pattern()
// 	// pattern2()
// 	// arr := []int{1, 2, 3, 4, 5}
// 	// index := 3
// 	// fmt.Println(arr)
// 	// arr = insert(arr, index, 600)
// 	// fmt.Println(arr)
// 	fmt.Println(isValid("([])"))
// }

// func insert(arr []int, pos, element int) []int {
// 	arr = append(arr, element)
// 	fmt.Println(arr)
// 	for i := len(arr) - 1; i > pos; i-- {
// 		arr[i], arr[i-1] = arr[i-1], arr[i]
// 		fmt.Println(arr)
// 	}
// 	return arr
// }

// func isValid(s string) bool {
// 	m := map[rune]rune{'}': '{', ']': '[', ')': '('}
// 	temp := []rune{}
// 	for _, v := range s {
// 		if _, ok := m[v]; ok {
// 			if len(temp) != 0 && m[v] == temp[len(temp)-1] {
// 				fmt.Println(temp)
// 				temp = temp[:len(temp)-1]
// 				fmt.Println(temp)
// 			} else {
// 				return false
// 			}
// 		} else {
// 			temp = append(temp, v)
// 		}
// 	}

// 	return len(temp) == 0
// }
