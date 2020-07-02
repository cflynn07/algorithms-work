package sum

import "testing"

func BenchmarkAdd(b *testing.B) {
	n1 := []int{9, 9, 9, 9}
	n2 := []int{9, 9, 9, 9}
	sum := make([]int, len(n1)+1)

	for i := 0; i < b.N; i++ {
		add(n1, n2, sum)
	}
}

func BenchmarkPlus(b *testing.B) {
	n1 := []int{9, 9, 9, 9}
	n2 := []int{9, 9, 9, 9}
	sum := make([]int, len(n1)+1)

	for i := 0; i < b.N; i++ {
		plus(n1, n2, sum)
	}
}
