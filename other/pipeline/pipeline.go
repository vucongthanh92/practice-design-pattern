package pipeline

func generatePipeline(nums []int) <-chan int {
	out := make(chan int)

	go func() {
		for _, num := range nums {
			out <- num
		}
	}()

	return out
}

func fanOut(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()

	return out
}

func fanIn(inputChan ...<-chan int) int {
	var sum = 0

	for _, c := range inputChan {
		sum += <-c
	}

	return sum
}
