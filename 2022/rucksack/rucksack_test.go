package rucksack

import (
	"fmt"
	"reflect"
	"testing"
)

type SetContentsTest struct {
	input              string
	expected           [2]string
	expected_duplicate rune
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
