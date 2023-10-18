package set

import (
	"fmt"
	"testing"
)

func TestGenerateSet(t *testing.T) {
	properInput := "-1, 1, 0, 2, 3"
	set, err := GenerateSet(properInput)
	if err != nil {
		t.Errorf("expected proper generation of set, got error: %v", err)
	}
	for i, elem := range set.Elems {
		if elem != i-1 {
			t.Errorf("expected %v, got: %v", elem, i-1)
		}
	}
	wrongInputs := []string{"1, 2,3", "1, a", ",,,"}
	for _, input := range wrongInputs {
		set, err = GenerateSet(input)
		expected := fmt.Errorf("check that the inserted numbers are separated by a comma and a space")
		if err == nil {
			t.Errorf("expected error: %v, got: %v", expected, nil)
		}
	}
}

func TestMakesElemsUnique(t *testing.T) {
	set := Set{Elems: []int{1, 1, 1, 2, 3, 4, 5, 6, 1, 1, 1, 1, 2, 3}}
	set.MakesElemsUnique()
	if len(set.Elems) != 6 {
		t.Errorf("expected set: %v, got: %v", []int{1, 2, 3, 4, 5, 6}, set.Elems)
	}
	for i, elem := range set.Elems {
		if elem != i+1 {
			t.Errorf("expected: %v at set.Elems[%v], got: %v", i+1, i, elem)
		}
	}
}

func TestGenerateCombinations(t *testing.T) {
	set := Set{Elems: []int{1, 2, 3}}
	set.GenerateCombinations()
	expectedCombinations := []Combination{Combination{}, Combination{1}, Combination{1, 2}, Combination{1, 2}, Combination{3}, Combination{1, 3}, Combination{2, 3}, Combination{1, 2, 3}}
	for i, elem := range set.Combinations {
		if len(elem) != len(expectedCombinations[i]) {
			t.Error("the length of the subsets obtained in a given step should be the same")
		}
		for j, element := range elem {
			if element != expectedCombinations[i][j] {
				t.Errorf("subset obtained in step %v should be: %v, got: %v", i, expectedCombinations[i], elem)
			}
		}
	}
}

func TestIsMobile(t *testing.T) {
	permutationsWithoutMobileInteger := []DirectedPermutation{DirectedPermutation{DirectedNumber{Number: 1, Direction: "right"}, DirectedNumber{Number: 2, Direction: "right"}, DirectedNumber{Number: 3, Direction: "right"}},
		DirectedPermutation{DirectedNumber{Number: 2, Direction: "left"}, DirectedNumber{Number: 1, Direction: "left"}, DirectedNumber{Number: 3, Direction: "right"}}}
	for _, elem := range permutationsWithoutMobileInteger {
		for i := range elem {
			isMobile := elem.IsMobile(i)
			if isMobile {
				t.Errorf("set: %v doesn't contain any mobile integer", elem)
			}
		}
	}
	permutationWithMobileInteger := DirectedPermutation{DirectedNumber{Number: 1, Direction: "left"}, DirectedNumber{Number: 2, Direction: "left"}, DirectedNumber{Number: 3, Direction: "left"}}
	anyMobile := false
	for i := 0; i < len(permutationWithMobileInteger); i++ {
		if permutationWithMobileInteger.IsMobile(i) {
			anyMobile = true
			break
		}
	}
	if !anyMobile {
		t.Errorf("set: %v contains atleast one mobile integer", permutationWithMobileInteger)
	}
}

func TestFindIndexOfGreatestMobileNumber(t *testing.T) {
	wrongPerm := DirectedPermutation{DirectedNumber{Number: 1, Direction: "right"}, DirectedNumber{Number: 2, Direction: "right"}, DirectedNumber{Number: 3, Direction: "right"}}
	_, err := wrongPerm.FindIndexOfGreatestMobileNumber()
	if err == nil {
		t.Error("there is no mobile integer, expected error")
	}
	perms := []DirectedPermutation{DirectedPermutation{DirectedNumber{Number: 7, Direction: "left"}, DirectedNumber{Number: 3, Direction: "right"}, DirectedNumber{Number: 1, Direction: "left"}},
		DirectedPermutation{DirectedNumber{Number: 3, Direction: "right"}, DirectedNumber{Number: 1, Direction: "left"}, DirectedNumber{Number: 7, Direction: "left"}, DirectedNumber{Number: 9, Direction: "right"}}}
	for i, elem := range perms {
		index, err := elem.FindIndexOfGreatestMobileNumber()
		if err != nil {
			t.Errorf("did not expect any error, got %v", err)
		}
		if index != i+1 {
			t.Errorf("index of greates mobile integer is %v, got %v", i+1, index)
		}
	}
}

func TestSwap(t *testing.T) {
	perm := DirectedPermutation{DirectedNumber{Number: 3, Direction: "right"}, DirectedNumber{Number: 1, Direction: "left"}, DirectedNumber{Number: 7, Direction: "left"}, DirectedNumber{Number: 9, Direction: "right"}}
	swapped := perm.Swap(0, perm[0])
	if swapped[0].Number != 1 || swapped[0].Direction != "left" {
		t.Errorf("after swap first elem in permutation should be: %v, got: %v", perm[1], swapped[0])
	}
	if swapped[1].Number != 3 || swapped[1].Direction != "right" {
		t.Errorf("after swap second elem in permutation should be: %v, got: %v", perm[0], swapped[1])
	}
	swapped = perm.Swap(2, perm[2])
	if swapped[1].Number != 7 || swapped[1].Direction != "left" {
		t.Errorf("after swap second elem in permutation should be: %v, got: %v", perm[2], swapped[1])
	}
	if swapped[2].Number != 1 || swapped[2].Direction != "left" {
		t.Errorf("after swap third elem in permutation should be: %v, got: %v", perm[1], swapped[2])
	}
}

func TestChangeDirection(t *testing.T) {
	perm := DirectedPermutation{DirectedNumber{Number: 3, Direction: "right"}, DirectedNumber{Number: 1, Direction: "left"}, DirectedNumber{Number: 7, Direction: "left"}, DirectedNumber{Number: 9, Direction: "right"}}
	perm.ChangeDirection(3)
	// here we expect that direction of every larger number than perm[0].Number will change to opposite
	expectedOutput := DirectedPermutation{DirectedNumber{Number: 3, Direction: "right"}, DirectedNumber{Number: 1, Direction: "left"}, DirectedNumber{Number: 7, Direction: "right"}, DirectedNumber{Number: 9, Direction: "left"}}
	for i, elem := range perm {
		if elem.Direction != expectedOutput[i].Direction {
			t.Errorf("expect direction: %v got: %v", expectedOutput[i].Direction, perm[i].Direction)
		}
	}
}

func TestCreateFirstDirectedPermutation(t *testing.T) {
	set := Set{Elems: []int{1, 2, 3, 4, 5, 6}}
	set.CreateFirstDirectedPermutation()
	if len(set.Permutations[0]) != len(set.Elems) {
		t.Errorf("set conatins %v elems so permutation should contain same number of elems, got %v elems", len(set.Permutations[0]), len(set.Elems))
	}
}
