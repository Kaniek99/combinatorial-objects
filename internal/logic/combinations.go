package combinations

import (
	"fmt"
	"math"
)

// type Combination []int
type Combination map[string]bool

type Set struct {
	Cardinality int
	Set         []string
	Subsets     []Combination
}

// type Combination struct {
// 	Cardinality int
// 	Subset      []int
// }

func (set *Set) generateSubsets() {
	combination := make(Combination)

	for i := 0; i < set.Cardinality; i++ { // what with empty set
		elemName := fmt.Sprint("x_", i)
		combination[elemName] = false
	} // now every combination is equal to empty set

	for i := 0; i < int(math.Pow(2.0, float64(set.Cardinality))); i++ {

	}

	for subsetCardinality := 0; subsetCardinality <= set.Cardinality; subsetCardinality++ {

	}
}
