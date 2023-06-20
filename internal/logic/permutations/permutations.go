package permutations

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Permutation []int
type InversionSequence []int
type DirectedPermutation []DirectedNumber

type DirectedNumber struct {
	Number    int
	Direction string
}

// TODO in combinations.go is already struct called Set, it should be unified and belongs to the common part
type NumericSet struct {
	Cardinality  int
	Elems        []int
	Permutations []DirectedPermutation
}

func CreateInversionSequence(input string) (InversionSequence, error) {
	seq := strings.Split(input, ", ")
	invSeq := []int{}
	for i, elem := range seq {
		el, err := strconv.Atoi(elem)
		if err != nil {
			e := fmt.Errorf("check that the numbers entered are separated by a comma and a space")
			return []int{}, e
		}
		if el > len(seq)-i-1 {
			err := fmt.Errorf("elem with index %v does not meet the assumption of inversion sequence %v < %v", i, el, len(seq)-i-1)
			return []int{}, err
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

func Remove(slice []int, i int) []int { //TODO this probably should be in common part
	if i == len(slice)-1 {
		return slice[:i]
	}
	return append(slice[:i], slice[i+1:]...)
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
	anyMobile := false
	indexOfGMNumber := -1
	for !anyMobile {
		for i := range *perm {
			if perm.IsMobile(i) {
				anyMobile = true
				indexOfGMNumber = i
			}
		}
	}
	if !anyMobile {
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
	if mobileNumber.Direction == "left" {
		perm[index] = perm[index-1]
		perm[index-1] = mobileNumber
		fmt.Println("gretest mobile number is:", mobileNumber)
		perm.ChangeDirection(mobileNumber.Number)
		return perm
	}
	perm[index] = perm[index+1]
	perm[index+1] = mobileNumber
	perm.ChangeDirection(mobileNumber.Number)
	fmt.Println("gretest mobile number is:", mobileNumber)
	return perm
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

func (set *NumericSet) CreateFirstDirectedPermutation() DirectedPermutation {
	sort.Ints(set.Elems) // sort while
	var directedPerm DirectedPermutation
	for _, elem := range set.Elems {
		directedNum := DirectedNumber{Number: elem, Direction: "left"}
		directedPerm = append(directedPerm, directedNum)
	}
	set.Permutations = append(set.Permutations, directedPerm)
	return directedPerm
}

func (set *NumericSet) GenerateAllPermutations() {
	perm := set.CreateFirstDirectedPermutation()
	fact := Factorial(set.Cardinality)
	for i := 1; i < fact; i++ {
		index, err := perm.FindIndexOfGreatestMobileNumber()
		if err != nil {
			fmt.Println("???")
			return
		}
		perm = perm.Swap(index, perm[index])
		set.Permutations = append(set.Permutations, perm)
	}
}

func (set *NumericSet) GetPermutationsAsString() string {
	output := "\n"
	for _, permutation := range set.Permutations {
		output += "("
		for i := 0; i < len(permutation); i++ {
			elemAsString := strconv.Itoa(permutation[i].Number)
			if i < len(permutation)-1 {
				output += elemAsString + ", "
			} else {
				output += elemAsString + ")\n"
			}
		}
	}
	return output
}

func GenerateSet(input string) (NumericSet, error) {
	elemsAsStrings := strings.Split(input, ", ")
	set := NumericSet{}
	for _, elem := range elemsAsStrings {
		el, err := strconv.Atoi(elem)
		if err != nil {
			e := fmt.Errorf("check that the inserted numbers are separated by a comma and a space")
			return NumericSet{}, e
		}
		set.Elems = append(set.Elems, el)
	}
	set.Cardinality = len(set.Elems)
	return set, nil
}

func Factorial(n int) int { // common part
	res := 1
	for i := 1; i <= n; i++ {
		res = res * i
	}
	return res
}
