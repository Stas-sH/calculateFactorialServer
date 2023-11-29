package factorial

func CalculateFactorial(n int, ch chan<- int) {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	ch <- result
}
