package pipeline

import "math"

// LaunchPipeline is the Entrance
func LaunchPipeline(amount int, pow float64) int {
	// 1st stage
	firstCh := generator(amount)
	// 2nd stage
	secondCh := power(firstCh, pow)
	// 3rd stage
	thirdCh := sum(secondCh)

	// read result
	result := <-thirdCh

	return result
}

func generator(max int) <-chan int {
	outChInt := make(chan int, 100)

	// generate number
	go func() {
		for i := 1; i <= max; i++ {
			outChInt <- i
		}

		close(outChInt)
	}()

	return outChInt
}

func power(in <-chan int, n float64) <-chan int {
	out := make(chan int, 100)

	// power of n
	go func() {
		for v := range in {
			out <- int(math.Pow(float64(v), n))
		}
		close(out)
	}()

	return out
}

func sum(in <-chan int) <-chan int {
	out := make(chan int, 100)

	// sum
	go func() {
		var sum int
		for v := range in {
			sum += v
		}
		out <- sum
		close(out)
	}()

	return out
}
