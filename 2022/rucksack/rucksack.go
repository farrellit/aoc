package rucksack

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type Rucksack struct {
	Compartments [2]string
}

func (r *Rucksack) SetContents(s string) {
	r.Compartments = [2]string{
		s[0 : len(s)/2],
		s[len(s)/2:],
	}
}

func (r *Rucksack) Duplicate() rune {
	for _, o := range r.Compartments[1] {
		for _, p := range r.Compartments[0] {
			if o == p {
				return o
			}
		}
	}
	return 0
}

func Intersect(s1, s2 []rune) (result []rune) {
	result = make([]rune, 0)
outer:
	for idx, i := range s1 {
		// avoid duplicates
		for _, isdup := range s1[0:idx] {
			if isdup == i {
				continue outer
			}
		}
		for _, j := range s2 {
			if i == j {
				result = append(result, i)
				continue outer
			}
		}
	}
	return
}

func (r *Rucksack) Contents() []rune {
	return []rune(r.Compartments[0] + r.Compartments[1])
}

func RucksackIntersect(rucksacks ...*Rucksack) (candidates []rune) {
	candidates = rucksacks[0].Contents()
	for _, s := range rucksacks[1:] {
		candidates = Intersect(candidates, s.Contents())
	}
	return
}

func Priority(r rune) int {
	if r >= 'a' && r <= 'z' {
		return 1 + int(r-'a')
	}
	if r >= 'A' && r <= 'Z' {
		return 27 + int(r-'A')
	}
	return 0
}

type Priorities []int

func (p Priorities) Sum() (s int) {
	if p != nil {
		for _, p := range p {
			s += p
		}
	}
	return
}

func GroupBadges(groups [][]*Rucksack) (res []rune) {
	res = make([]rune, len(groups))
	for grpid, rucksacks := range groups {
		shared := RucksackIntersect(rucksacks...)
		if len(shared) != 1 {
			panic(fmt.Errorf(
				"Unexpected duplicate shared runes between group of rucksacks: %s",
				string(shared),
			))
		}
		res[grpid] = shared[0]
	}
	return
}

func GroupRucksacks(rs []*Rucksack, size int) (res [][]*Rucksack) {
	for i := 0; i < len(rs); i += size {
		lim := i + 3
		if lim > len(rs) {
			lim = len(rs)
		}
		res = append(res, rs[i:lim])
	}
	return
}

type RucksackList []*Rucksack

func (r *RucksackList) Read(input io.Reader) {
	scan := bufio.NewScanner(input)
	for scan.Scan() {
		line := strings.TrimSpace(scan.Text())
		s := new(Rucksack)
		s.SetContents(line)
		*r = append(*r, s)
	}
}

type RuneList []rune

func (r RuneList) Priorities() (ps Priorities) {
	for _, r := range r {
		ps = append(ps, Priority(r))
	}
	return
}

func (rlist RucksackList) Priorities() Priorities {
	var dups RuneList
	for _, ruck := range rlist {
		dups = append(dups, ruck.Duplicate())
	}
	return dups.Priorities()
}
