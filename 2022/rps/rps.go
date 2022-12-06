package rps

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type ResultScore int

const (
	WinScore  ResultScore = 6
	DrawScore             = 3
	LoseScore             = 0
)

type Choice uint8

func (c Choice) String() string {
	switch c {
	case Rock:
		return "Rock"
	case Paper:
		return "Paper"
	case Scissors:
		return "Scissors"
	}
	return "Invalid"
}

const (
	Invalid  Choice = 0
	Rock            = 1
	Paper           = 2
	Scissors        = 3
)

type Round struct {
	choices [2]Choice
}

func (r Round) Scores() (scores [2]int) {
	scores[0], scores[1] = r.Score()
	return
}

func (r Round) Score() (score0, score1 int) {
	score0, score1 = int(r.choices[0]), int(r.choices[1])
	// draw
	if r.choices[0] == r.choices[1] {
		score0 += int(DrawScore)
		score1 += int(DrawScore)
	} else {
		switch r.choices[0] {
		case Rock:
			switch r.choices[1] {
			case Paper:
				score1 += int(WinScore)
				score0 += int(LoseScore)
			case Scissors:
				score0 += int(WinScore)
				score1 += int(LoseScore)
			}
		case Paper:
			switch r.choices[1] {
			case Rock:
				score0 += int(WinScore)
				score1 += int(LoseScore)
			case Scissors:
				score1 += int(WinScore)
				score0 += int(LoseScore)
			}
		case Scissors:
			switch r.choices[1] {
			case Rock:
				score1 += int(WinScore)
				score0 += int(LoseScore)
			case Paper:
				score0 += int(WinScore)
				score1 += int(LoseScore)
			}
		}
	}
	return
}

type Game struct {
	rounds []Round
}

func ParseChoice(s string) (Choice, error) {
	switch strings.ToUpper(s) {
	case "A":
		fallthrough
	case "X":
		return Rock, nil
	case "B":
		fallthrough
	case "Y":
		return Paper, nil
	case "C":
		fallthrough
	case "Z":
		return Scissors, nil
	}
	return Invalid, fmt.Errorf("'%s' is not a valid choice", s)
}

func (g *Game) ReadRounds(input io.Reader) error {
	scan := bufio.NewScanner(input)
	var lineno = 0
	for scan.Scan() {
		lineno++
		line := strings.TrimSpace(scan.Text())
		fields := strings.Fields(line)
		if len(fields) < 2 {
			return fmt.Errorf("Invalid line #%d: must have at least 2 whitespace separated fields: %s", lineno, line)
		}
		var round Round
		for i := 0; i < 2; i++ {
			if c, err := ParseChoice(fields[i]); err != nil {
				return fmt.Errorf("Error parsing line #%d '%s' field #%d '%s': %w",
					lineno, line, i+1, fields[i], err)
			} else {
				round.choices[i] = c
			}
		}
		g.rounds = append(g.rounds, round)
	}
	return nil
}

func (g *Game) Scores() (scores [2]int) {
	if g.rounds == nil {
		return
	}
	for _, r := range g.rounds {
		s := r.Scores()
		scores[0] += s[0]
		scores[1] += s[1]
	}
	return
}
