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

func TestUniqueElems(t *testing.T) {
	set := Set{Elems: []int{1, 1, 1, 2, 3, 4, 5, 6, 1, 1, 1, 1, 2, 3}}
	set.UniqueElems()
	if len(set.Elems) != 6 {
		t.Errorf("expected set: %v, got: %v", []int{1, 2, 3, 4, 5, 6}, set.Elems)
	}
	for i, elem := range set.Elems {
		if elem != i+1 {
			t.Errorf("expected: %v at set.Elems[%v], got: %v", i+1, i, elem)
		}
	}
}
