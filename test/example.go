package test

// Avg gets the average of
// numbers passed in
func Avg(num ...int) int {
	sum := 0
	for _, value := range num {
		sum += value
	}
	if sum == 0 {
		return 0
	}
	return sum / len(num)
}
