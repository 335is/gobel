/// This program calculates the first 13 2-Gobel Sequence values.
/// https://mathworld.wolfram.com/GoebelsSequence.html
/// https://en.wikipedia.org/wiki/G%C3%B6bel%27s_sequence
///
/// 2-Gobel sequence values (squared terms)
/// k:     0, 1, 2,  3,  4,   5,   6,     7,        8, ...
/// sum:   1, 2, 6, 15, 40, 140, 924, 24640, 12415040, ...
/// value: 1, 2, 3,  5, 10,  28, 154,  3520,  1551880, ...

package main

import (
	"fmt"
)

func main() {
	// start := time.Now()

	var val float64 = 0.0

	var sum float64 = 0.0
	for k := 0; k < 13; k++ {
		val, sum = calcK(k, sum)
		fmt.Printf("gobel(%d) = %.f\n", k, val)
		// fmt.Printf("gobel(%d) = %.f, sum = %.f\n", k, val, sum)
	}

	// dur := time.Since(start)
}

// k = iteration number
// sum = sum from previous iteration
// returns: kth gobel value, new sum for next iteration
func calcK(k int, sum float64) (float64, float64) {
	// the k=0 term is a hard value, the rest are calculated
	if k == 0 {
		return 1.0, 2.0
	}

	// new value is the previous sum divided by k
	t := sum / float64(k)

	// calculate the new sum for next time
	newsum := (t * t) + sum

	return t, newsum
}
