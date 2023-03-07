package combinations

import (
	"fmt"
	"log"
	"math"
)

type Combination map[string]bool

type Set struct {
	Cardinality int
	Set         []string
	Subsets     []Combination
}

func (set *Set) GenerateSubsets() {
	for i := 0; i < int(math.Pow(2.0, float64(set.Cardinality))); i++ {
		combination := make(Combination)

		for j := 0; j < set.Cardinality; j++ { // what with empty set
			elemName := fmt.Sprint("x_", j)
			combination[elemName] = false
		} // now every combination is equal to empty set

		test := fmt.Sprintf("%b", i)
		fmt.Println(test)
		for j, elem := range test {
			if elem == '1' {
				combination[fmt.Sprint("x_", len(test)-j-1)] = true
			} else {
				combination[fmt.Sprint("x_", len(test)-j-1)] = false
			}
		}
		set.Subsets = append(set.Subsets, combination)
	}
}

func (set *Set) PrintSet() {
	// for _, elem := range set.Set {

	// }
	fmt.Println(set.Set)
}

func (set *Set) GetCombinations() [][]string {
	combinations := [][]string{}
	for _, combination := range set.Subsets {
		subset := []string{}
		for elem, presence := range combination {
			if presence == true {
				subset = append(subset, elem)
			}
		}
		combinations = append(combinations, subset)
	}
	return combinations
	// fmt.Println(set.Subsets)
}

func GenerateSet(cardinality int) Set {
	// var cardinality int
	// fmt.Scan(&cardinality)
	// if cardinality < 0 {
	// retry, cardinality need to be uint
	// }
	log.Println("cardinality is: ", cardinality)
	var set []string
	for i := cardinality - 1; i >= 0; i-- {
		set = append(set, fmt.Sprint("x_", i))
	}
	return Set{Cardinality: cardinality, Set: set}
}
