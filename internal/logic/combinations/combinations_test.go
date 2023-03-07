package combinations

import "testing"

func TestGenerateSubsets(t *testing.T) {
	test := Set{Cardinality: 3, Set: []string{"x_2", "x_1", "x_0"}}
	test.GenerateSubsets()
}

func TestGenerateSet(t *testing.T) {
	// how to write unit tests with user input? maybe send input as a parameter?
}
