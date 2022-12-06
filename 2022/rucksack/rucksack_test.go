package rucksack

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"
)

type SetContentsTest struct {
	input              string
	expected           [2]string
	expected_duplicate rune
}

type PriorityTest struct {
	r rune
	e int
}

func TestPriority(t *testing.T) {
	tests := []PriorityTest{
		PriorityTest{'p', 16},
		PriorityTest{'L', 38},
		PriorityTest{'P', 42},
		PriorityTest{'v', 22},
		PriorityTest{'t', 20},
		PriorityTest{'s', 19},
	}
	var priorities []int
	for _, test := range tests {
		priorities = append(priorities, Priority(test.r))
		t.Run(fmt.Sprintf("%c_%d", test.r, test.e), func(t *testing.T) {
			if res, exp := Priority(test.r), test.e; res != exp {
				t.Errorf("Wanted %c got %c ", exp, res)
			}
		})
	}
	t.Run("Sum==157", func(t *testing.T) {
		if res, exp := Priorities(priorities).Sum(), 157; res != exp {
			t.Errorf("Expected %d got %d", exp, res)
		}
	})
}

func TestRucksack(t *testing.T) {
	tests := []SetContentsTest{
		SetContentsTest{
			"vJrwpWtwJgWrhcsFMMfFFhFp",
			[2]string{"vJrwpWtwJgWr", "hcsFMMfFFhFp"},
			'p',
		},
		SetContentsTest{
			"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
			[2]string{"jqHRNqRjqzjGDLGL", "rsFMfFZSrLrFZsSL"},
			'L',
		},
		SetContentsTest{
			"PmmdzqPrVvPwwTWBwg",
			[2]string{"PmmdzqPrV", "vPwwTWBwg"},
			'P',
		},
		SetContentsTest{
			"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
			[2]string{"wMqvLMZHhHMvwLH", "jbvcjnnSBnvTQFn"},
			'v',
		},
		SetContentsTest{
			"ttgJtRGJQctTZtZT",
			[2]string{"ttgJtRGJ", "QctTZtZT"},
			't',
		},
		SetContentsTest{
			"CrZsJsPPZsGzwwsLwLmpwMDw",
			[2]string{"CrZsJsPPZsGz", "wwsLwLmpwMDw"},
			's',
		},
	}
	t.Run("SetContents", func(t *testing.T) {
		for _, test := range tests {
			t.Run(fmt.Sprintf("%v", test), func(t *testing.T) {
				var r Rucksack
				r.SetContents(test.input)
				if !reflect.DeepEqual(r.Compartments, test.expected) {
					t.Errorf("Wanted %v, got %#v", test.expected, r.Compartments)
				}
				if dup, exp := r.Duplicate(), test.expected_duplicate; exp != dup {
					t.Errorf("Wanted duplicate %s got %s", string(exp), string(dup))
				}
			})
		}
	})
}
func TestRucksackList(t *testing.T) {
	input := `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`
	var rl RucksackList
	rl.Read(bytes.NewBufferString(input))
	if s, exp := rl.Priorities().Sum(), 157; s != exp {
		t.Errorf("Expected sum %d got %d", exp, s)
	}
}
