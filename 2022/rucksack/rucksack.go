package rucksack

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

func Priority(r rune) int {
	if r >= 'a' && r <= 'z' {
		return 1 + int(r - 'a')
	}
	if r >= 'A' && r <= 'Z' {
		return 27 + int(r - 'A')
	}
	return 0
}

type Priorities []int

func (p Priorities)Sum() (s int) {
	if p != nil {
	for _, p := range p {
		s += p
	}
	}
	return
}
