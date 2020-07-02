package adder

// Add adds two numbers together
func Add(n1, n2, sum []int) {
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
func Plus(n1, n2, sum []int) {
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
