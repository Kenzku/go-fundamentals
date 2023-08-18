package main

import (
	"fmt"
)

func main() {
	testScores := map[string]float64{
		"Harry":    87.3,
		"Hermione": 105,
		"Ronald":   63.5,
		"Neville":  27,
	}
	c := clone(testScores)

	fmt.Println(c)
}

// Type Constraints: "K comparable: indicates K has to be a type of "comparable"
// because in map, the key has to be comparable
func clone[K comparable, V any](m map[K]V) map[K]V {
	result := make(map[K]V, len(m))
	for k, v := range m {
		result[k] = v
	}

	return result
}
