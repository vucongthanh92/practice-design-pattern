package pipeline

import "fmt"

func ExecPipeLine(max int) {
	ranNums := make([]int, 0)
	for i := 0; i <= max; i++ {
		ranNums = append(ranNums, i)
	}

	inputChan := generatePipeline(ranNums)

	c1 := fanOut(inputChan)
	c2 := fanOut(inputChan)
	c3 := fanOut(inputChan)
	c4 := fanOut(inputChan)

	sum := fanIn(c1, c2, c3, c4)

	fmt.Println("total sum: ", sum)
}
