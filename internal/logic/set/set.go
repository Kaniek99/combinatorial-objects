package set

import (
	"fmt"
	"sort"
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
	set.UniqueElems()
	sort.Ints(set.Elems)
	set.Cardinality = len(set.Elems)
	return set, nil
}

func (set *Set) UniqueElems() {
	keys := make(map[int]bool)
	elems := []int{}
	for _, entry := range set.Elems {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			elems = append(elems, entry)
		}
	}
	set.Elems = elems
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

func (perm *DirectedPermutation) IsMobile(i int) bool {
	if (i == 0 && (*perm)[i].Direction == "left") || (i == len(*perm)-1 && (*perm)[i].Direction == "right") {
		return false
	}
	var indexOfNeighbour int
	if (*perm)[i].Direction == "left" {
		indexOfNeighbour = i - 1
	} else {
		indexOfNeighbour = i + 1
	}
	if (*perm)[i].Number > (*perm)[indexOfNeighbour].Number {
		return true
	}
	return false
}

func (perm *DirectedPermutation) FindIndexOfGreatestMobileNumber() (int, error) {
	indexOfGMNumber := -1
	for i := range *perm {
		if perm.IsMobile(i) {
			indexOfGMNumber = i
			break
		}
	}
	if indexOfGMNumber == -1 {
		err := fmt.Errorf("there are no mobile numbers")
		return indexOfGMNumber, err
	}
	for i := indexOfGMNumber + 1; i < len(*perm); i++ {
		if perm.IsMobile(i) {
			if (*perm)[i].Number > (*perm)[indexOfGMNumber].Number {
				indexOfGMNumber = i
			}
		}
	}
	return indexOfGMNumber, nil
}

func (perm DirectedPermutation) Swap(index int, mobileNumber DirectedNumber) DirectedPermutation {
	newPerm := make(DirectedPermutation, len(perm))
	copy(newPerm, perm)
	if mobileNumber.Direction == "left" {
		newPerm[index] = perm[index-1]
		newPerm[index-1] = mobileNumber
		newPerm.ChangeDirection(mobileNumber.Number)
		return newPerm
	}
	newPerm[index] = perm[index+1]
	newPerm[index+1] = mobileNumber
	newPerm.ChangeDirection(mobileNumber.Number)
	return newPerm
}

func (perm *DirectedPermutation) ChangeDirection(greatestMobile int) {
	for i, elem := range *perm {
		if elem.Number > greatestMobile {
			if elem.Direction == "left" {
				(*perm)[i].Direction = "right"
			} else {
				(*perm)[i].Direction = "left"
			}
		}
	}
}

func (set *Set) CreateFirstDirectedPermutation() DirectedPermutation {
	var directedPerm DirectedPermutation
	for _, elem := range set.Elems {
		directedNum := DirectedNumber{Number: elem, Direction: "left"}
		directedPerm = append(directedPerm, directedNum)
	}
	set.Permutations = append(set.Permutations, directedPerm)
	return directedPerm
}

func (set *Set) GenerateAllPermutations() {
	perm := set.CreateFirstDirectedPermutation()
	fact := Factorial(set.Cardinality)
	for i := 1; i < fact; i++ {
		index, err := perm.FindIndexOfGreatestMobileNumber()
		if err != nil {
			fmt.Println("???")
			return
		}
		perm = perm.Swap(index, perm[index])
		permutation := make(DirectedPermutation, len(perm))
		copy(permutation, perm)
		set.Permutations = append(set.Permutations, permutation)
	}
}

func Factorial(n int) int { // common part
	res := 1
	for i := 1; i <= n; i++ {
		res = res * i
	}
	return res
}

// maybe it should not be here and I need to change files structure
func CreateInversionSequence(input string) (InversionSequence, error) {
	seq := strings.Split(input, ", ")
	invSeq := []int{}
	for i, elem := range seq {
		el, err := strconv.Atoi(elem)
		if err != nil {
			return []int{}, fmt.Errorf("check that the numbers entered are separated by a comma and a space")
		}
		if el > len(seq)-i-1 {
			return []int{}, fmt.Errorf("elem with index %v does not meet the assumption of inversion sequence %v < %v", i, el, len(seq)-i-1)
		}
		if el < 0 {
			return []int{}, fmt.Errorf("elems must be non-negative")
		}
		invSeq = append(invSeq, el)
	}
	return invSeq, nil
}

func (seq *InversionSequence) GeneratePermutation() Permutation {
	permutation := make([]int, len(*seq))
	remaining := []int{}
	for i := 0; i < len(*seq); i++ {
		remaining = append(remaining, i)
	}
	for i, elem := range *seq {
		permutation[remaining[elem]] = i + 1
		remaining = Remove(remaining, elem)
	}
	return permutation
}

func Remove(slice []int, i int) []int {
	if i == len(slice)-1 {
		return slice[:i]
	}
	return append(slice[:i], slice[i+1:]...)
}
