package permutations

import (
	"fmt"
	"strconv"
	"strings"
)

type Permutation []int
type InversionSequence []int

func CreateInversionSequence(input string) (InversionSequence, error) {
	seq := strings.Split(input, ", ")
	invSeq := []int{}
	for i, elem := range seq {
		el, err := strconv.Atoi(elem)
		if err != nil {
			error := fmt.Errorf("check that the numbers entered are separated by a comma and a space")
			return []int{}, error
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
		remaining = remove(remaining, elem)
	}
	return permutation
}

func remove(slice []int, i int) []int {
	if i == len(slice)-1 {
		return slice[:i]
	}
	return append(slice[:i], slice[i+1:]...)
}
