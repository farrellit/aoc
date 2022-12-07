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
				if c, exp := string(r.Contents()), test.input; c != exp {
					t.Errorf("Wanted Contents() %s got %s", exp, c)
				}
			})
		}
	})
}

type RucksackIntersectTest struct {
	start, end int
	exp        []rune
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
	t.Run("RucksackIntersect", func(t *testing.T) {
		for _, test := range []RucksackIntersectTest{
			RucksackIntersectTest{0, 3, []rune{'r'}},
			RucksackIntersectTest{3, 6, []rune{'Z'}},
		} {
			t.Run(
				fmt.Sprintf("%s %s",
					func() (s string) {
						for _, i := range rl[test.start:test.end] {
							if s != "" {
								s += ","
							}
							s += string(i.Contents())
						}
						return
					}(),
					string(test.exp),
				),
				func(t *testing.T) {
					res, exp := RucksackIntersect(rl[test.start:test.end]...), test.exp
					if !reflect.DeepEqual(res, exp) {
						t.Errorf("Expected %s got %s", string(exp), string(res))
					}
				},
			)
		}
	})
	t.Run("GroupRucksacks", func(t *testing.T) {
		groups := GroupRucksacks(rl[0:5], 3)
		if l, exp := len(groups), 2; l != exp {
			t.Errorf("Expected 2, got %d", l)
		}
		if res, exp := groups[0], rl[0:3]; !reflect.DeepEqual([]*Rucksack(exp), res) {
			t.Errorf("Expected %#v, got %#v", exp, res)
		}
		if res, exp := groups[1], rl[3:5]; !reflect.DeepEqual([]*Rucksack(exp), res) {
			t.Errorf("Expected %#v, got %#v", exp, res)
		}
	})
}

type IntersectTest struct {
	s1, s2 []rune
	exp    []rune
}

func TestIntersect(t *testing.T) {
	for _, test := range []IntersectTest{
		IntersectTest{[]rune("abc"), []rune("cdef"), []rune("c")},
		IntersectTest{[]rune("rk4X"), []rune("Jm5dccXej4"), []rune("4X")},
	} {
		t.Run(fmt.Sprintf("%v", test), func(t *testing.T) {
			if res, exp := Intersect(test.s1, test.s2), test.exp; !reflect.DeepEqual(res, exp) {
				t.Errorf("Expected %v got %v", exp, res)
			}
		})
	}

}
