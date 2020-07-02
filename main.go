package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	n1 := []int{1, 2, 3}

	fmt.Println(n1)
	fmt.Println(rotateL(n1), "rotated left")
	fmt.Println(rotateR(n1), "rotated right")

	generateTable()
}

// generate table comparing each approach in time
func generateTable() {
	const MaxSize = 1048576
	const NumTrials = 10000
	n := 256

	for n <= MaxSize {
		fmt.Println("Trying " + strconv.Itoa(n) + "...")
		n1 := make([]int, n)
		n2 := make([]int, n)
		sum := make([]int, n+1)

		randomNumber(n1, n)
		randomNumber(n2, n)

		// copy n1/n2?

		baseS := time.Now()
		fmt.Println("start: " + strconv.Itoa(n))
		for i := 0; i < NumTrials; i++ {
			// circular shift (n1 left, n2 right)
			n2 = rotateR(n2)
			n1 = rotateL(n1)
		}
		baseE := time.Now()

		fmt.Println(time.Since(baseS))
		// baseE := time.Now()
		n = n * 2
	}
}

// rotates a slice to the left
// ex {1, 2, 3} -> {2, 3, 1}
func rotateL(n1 []int) []int {
	return append(n1[1:], n1[0])
}

// rotates a slice to the right
// ex {1, 2, 3} -> {3, 1, 2}
func rotateR(n1 []int) []int {
	return append(n1[len(n1)-1:], n1[:len(n1)-1]...)
}

func randomNumber(num []int, n int) {
	for i := 0; i < n; i++ {
		num[i] = rand.Intn(10)
	}
}

// Add adds two numbers together
func add(n1, n2, sum []int) {
	p := len(n1) - 1
	carry := 0
	for p >= 0 {
		total := n1[p] + n2[p] + carry
		sum[p+1] = total % 10
		if total > 9 {
			carry = 1
		} else {
			carry = 0
		}
		p--
	}
	sum[0] = carry
}

// Plus adds two numbers together
func plus(n1, n2, sum []int) {
	p := len(n1)
	carry := 0
	for p--; p >= 0; p-- {
		total := n1[p] + n2[p] + carry
		if total > 9 {
			sum[p+1] = total - 10
			carry = 1
		} else {
			sum[p+1] = total
			carry = 0
		}
	}
	sum[0] = carry
}
