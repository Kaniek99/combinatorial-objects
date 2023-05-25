package combinations

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

type Combination map[string]bool

type Set struct {
	Cardinality int
	Set         []string
	Subsets     []Combination
}

func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}

func GenerateSet(input string) (Set, error) {
	defer timer("GenerateSet")()
	cardinality, err := strconv.Atoi(input)
	if err != nil || cardinality < 0 {
		return Set{}, fmt.Errorf("Insert integer greater than or equal to zero.")
	}
	var set []string
	for i := cardinality - 1; i >= 0; i-- {
		set = append(set, fmt.Sprint("x_", i))
	}
	return Set{Cardinality: cardinality, Set: set}, nil
}

func (set *Set) GenerateSubsets() {
	defer timer("GenerateSubsets")()
	for i := 0; i < int(math.Pow(2.0, float64(set.Cardinality))); i++ {
		combination := make(Combination)

		for j := 0; j < set.Cardinality; j++ {
			elemName := fmt.Sprint("x_", j)
			combination[elemName] = false
		} // now every combination is equal to empty set

		binaryRepresentation := fmt.Sprintf("%b", i)
		for j, elem := range binaryRepresentation {
			if elem == '1' {
				combination[fmt.Sprint("x_", len(binaryRepresentation)-j-1)] = true
			} else {
				combination[fmt.Sprint("x_", len(binaryRepresentation)-j-1)] = false
			}
		}
		set.Subsets = append(set.Subsets, combination)
	}
}

func (set *Set) GetSet() string {
	output := "{"
	for i, elem := range set.Set {
		output += elem
		if i < set.Cardinality-1 {
			output += ", "
		}
	}
	output += "}"
	return output
}

func (set *Set) GetCombinations() string {
	combinations := ""
	for i, subset := range set.Subsets {
		combination := []string{}
		for elem, presence := range subset {
			if presence {
				combination = append(combination, elem)
			}
		}
		CombinationAsSet := Set{Cardinality: len(combination), Set: combination}
		combinations += CombinationAsSet.GetSet()
		if i < len(set.Subsets)-1 {
			combinations += ", "
		}
	}
	return combinations
}
