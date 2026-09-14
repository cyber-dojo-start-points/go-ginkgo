package hiker

import "fmt"

func answer() int {
	// The learner put a print inside a loop to see what was happening, and it
	// prints far more than the 50K the runner keeps.
	total := 0
	for i := 0; i < 5000; i++ {
		fmt.Printf("debug: i is %d, total is %d\n", i, total)
		total++
	}
	return 6 * 7
}
