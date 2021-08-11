package counter

import "fmt"

func CountSteps(n uint64) (steps, values []uint64) {
	var step uint64
	steps = append(steps, 0)
	values = append(values, n)
	fmt.Printf("\n\nStart : %4d\tn = %4d\n", step, n)
	for n != 1 {
		step++
		if n%2 == 0 {
			n /= 2
		} else {
			n = (n * 3) + 1
		}
		fmt.Printf("Step  : %4d\tn = %4d\n", step, n)
		steps = append(steps, step)
		values = append(values, n)
	}
	return
}
