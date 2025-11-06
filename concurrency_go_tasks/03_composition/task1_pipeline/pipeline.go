package pipeline

// Run строит конвейер из трёх стадий: квадрат, умножение на 2 и суммирование.
func Run(nums []int) int {
	input := make(chan int)
	squared := make(chan int)
	multiplied := make(chan int)
	resultChan := make(chan int)

	go func() {
		for n := range input {
			squared <- n * n
		}
		close(squared)
	}()

	go func() {
		for n := range squared {
			multiplied <- n * 2
		}
		close(multiplied)
	}()

	go func() {
		total := 0
		for n := range multiplied {
			total += n
		}
		resultChan <- total
		close(resultChan)
	}()

	for _, n := range nums {
		input <- n
	}
	close(input)

	return <-resultChan
}
