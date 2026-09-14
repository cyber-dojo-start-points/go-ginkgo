package hiker

import (
	"fmt"
	"os"
)

func answer() int {
	fmt.Fprintln(os.Stderr, "answer was called")
	return 6 * 7
}
