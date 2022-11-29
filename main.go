package main

import (
	"fmt"
)

func main() {
	// var n int
	// fmt.Scan(&n)
	// res := solution(n)
	// fmt.Println(res)
	arr := []int{5, 4, 3, 2, 1}
	mergeSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

// func solution(n int) int { // 133
// 	n++ // 134
// 	for {
// 		flag := true

// 		s := strconv.Itoa(n) // '134'

// 		for i := 0; i < len(s)-1; i++ {
// 			if s[i] == s[i+1] {
// 				flag = false
// 				break
// 			}
// 		}

// 		if flag {
// 			return n
// 		}
// 		n++
// 	}
// }

// func x() {
// 	var v int

// 	v = "abc"

// }
func mergeSort(arr []int, s, e int) {
	if s > e {
		return
	}
	mid := (s + e) / 2
	mergeSort(arr, s, mid)
	mergeSort(arr, mid+1, e)

	merge(arr, s, mid, e)
}

func merge(arr []int, s, mid, e int) {
	i := s
	j := mid + 1
	temp := []int{}

	for i <= mid && j <= e {
		if arr[i] <= arr[j] {
			temp = append(temp, arr[i])
			i++
		} else {
			temp = append(temp, arr[j])
			j++
		}
	}

	if i <= mid {
		temp = append(temp, arr[i:]...)
	}
	if j <= e {
		temp = append(temp, arr[j:]...)
	}

	for i := s; i <= e; i++ {
		arr[i] = temp[s-i]
	}

	// x := map[int][]int{10:[]int{10, 20}}

}
