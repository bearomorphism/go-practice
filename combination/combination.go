package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var n int
	var m int

	fmt.Scanf("%d %d", &n, &m)
	if n < m {
		n, m = m, n
	}

	arr := make([]int, n)
	for i := range arr {
		arr[i] = i
	}
	fmt.Println(combination(m, arr))
}

func combination(a int, arr []int) []int {
	// fmt.Println(arr[a])
	for i := 0; i < a; i++ {
		j := rand.Intn(len(arr)-i) + i
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr[:a]
}
