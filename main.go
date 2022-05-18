package main

import "fmt"

func bubbleSort(n []int) []int {
	for i, _ := range n {
		for j := 0; j < len(n)-i-1; j++ {
			if n[j] > n[j+1] {
				n[j], n[j+1] = n[j+1], n[j]
			}
		}
	}
	return n
}

func main() {
	usorted := []int{5, 4, 3, 15, 13, 20, 49, 2, 1}
	bubbleSort(usorted)
	fmt.Println(bubbleSort(usorted))

}
