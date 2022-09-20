package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var n int
	fmt.Scanf("%d", &n)
	arr := make([]int, n)
	for i := range arr {
		arr[i] = rand.Int() % 100
	}
	fmt.Println(arr)
	rand.Shuffle(n, func(i, j int) { arr[i], arr[j] = arr[j], arr[i] })
	fmt.Println(arr)
	quicksort(arr, 0, n-1)
	fmt.Println(arr)
}

func quicksort(arr []int, l int, r int) {
	if l > r {
		return
	}
	i := l
	j := r
	pivot := arr[l]

	for i != j {
		for arr[j] > pivot && i < j {
			j--
		}
		for arr[i] <= pivot && i < j {
			i++
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[l] = arr[i]
	arr[i] = pivot
	quicksort(arr, l, i-1)
	quicksort(arr, i+1, r)
}
