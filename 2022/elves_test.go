package main

import (
	"bytes"
	"sort"
	"testing"
)

func TestElf(t *testing.T) {
	t.Run("EmptyTotalCalories", func(t *testing.T) {
		e := new(Elf)
		if e.TotalCalories() != 0 {
			t.Errorf("New elf should have 0 calories")
		}
	})
	t.Run("AddCaloriesAndTotalCalories", func(t *testing.T) {
		e := new(Elf)
		e.AddCalories(1, 2, 3)
		if e.TotalCalories() != 6 {
			t.Errorf("AddCalories")
		}
	})
}

type ElvesByCaloriesTest struct {
	Elves []*Elf
	Len   int
}

func TestElvesByCalories(t *testing.T) {
	t.Run("Len", func(t *testing.T) {
		for _, elves := range []ElvesByCaloriesTest{
			ElvesByCaloriesTest{
				Elves: []*Elf{},
				Len:   0,
			},
			ElvesByCaloriesTest{
				Elves: []*Elf{
					&Elf{calories: []int64{1, 2, 3}},
				},
				Len: 1,
			},
		} {
			el := ElvesByCalories(elves.Elves)
			if l := el.Len(); l != elves.Len {
				t.Errorf("Expected %d, got %d", elves.Len, l)
			}
		}
	})
}

func TestReadCalorieList(t *testing.T) {
	input := bytes.NewBufferString(`1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`)
	elves := ReadCalorieList(input)
	if l := len(elves); l != 5 {
		t.Errorf("Expected 5 elves, got %d", l)
	}
	sort.Sort(ElvesByCalories(elves))
	if m := elves[len(elves)-1].TotalCalories(); m != 24000 {
		t.Errorf("Expected max TotalCalories of 24000; got %d", m)
	}
	if m := ElvesByCalories(elves[len(elves)-2:]).TotalCalories(); m != 35000 {
		t.Errorf("Expected list TotalCalories of 35000; got %d", m)
	}
}
