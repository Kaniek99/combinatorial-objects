package set

import (
	"fmt"
	"strconv"
	"strings"
)

type Permutation []int
type InversionSequence []int
type DirectedPermutation []DirectedNumber
type Combination []int

type DirectedNumber struct {
	Number    int
	Direction string
}

type Set struct {
	Cardinality  int
	Elems        []int
	Combinations []Combination
	Permutations []DirectedPermutation
}

func GenerateSet(input string) (Set, error) {
	elemsAsStrings := strings.Split(input, ", ")
	set := Set{}
	for _, elem := range elemsAsStrings {
		el, err := strconv.Atoi(elem)
		if err != nil {
			e := fmt.Errorf("check that the inserted numbers are separated by a comma and a space")
			return Set{}, e
		}
		set.Elems = append(set.Elems, el)
	}
	set.Cardinality = len(set.Elems)
	return set, nil
}

// GenerateCombinations generates all combinations of a set using binary sequences
func (set *Set) GenerateCombinations() {
	max := 1 << set.Cardinality
	for i := 0; i < max; i++ {
		combination := Combination{}
		for j := 0; j < set.Cardinality; j++ {
			if (i & (1 << j)) != 0 {
				combination = append(combination, set.Elems[j])
			}
		}
		set.Combinations = append(set.Combinations, combination)
	}
}
