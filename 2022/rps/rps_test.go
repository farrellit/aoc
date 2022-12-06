package rps

import (
	"bytes"
	"fmt"
	"testing"
)

type ChoiceTest struct {
	c   Choice
	str string
}

func TestChoice(t *testing.T) {
	for _, test := range []ChoiceTest{
		ChoiceTest{Rock, "Rock"},
		ChoiceTest{Paper, "Paper"},
		ChoiceTest{Scissors, "Scissors"},
		ChoiceTest{0, "Invalid"},
		ChoiceTest{4, "Invalid"},
	} {
		t.Run(fmt.Sprintf("%d", test.c), func(t *testing.T) {
			t.Run("StringEq"+test.str, func(t *testing.T) {
				if s := test.c.String(); s != test.str {
					t.Errorf("Excpected %s got %s", test.str, s)
				}
			})
		})
	}
}

type RoundTest struct {
	round  Round
	scores [2]int
}

func (rt *RoundTest) Name() string {
	return fmt.Sprint(rt.round, "Score", rt.round.Scores(), "Eq", rt.scores)
}
func (rt *RoundTest) Test(t *testing.T) {
	score0, score1 := rt.round.Score()
	if score0 != rt.scores[0] || score1 != rt.scores[1] {
		t.Errorf("Hand %v: expected %v, got %v",
			rt.round, rt.scores, []int{score0, score1},
		)
	}
}

func TestRound(t *testing.T) {
	for _, test := range []RoundTest{
		RoundTest{Round{[2]Choice{Rock, Paper}}, [2]int{1, 8}},
		RoundTest{Round{[2]Choice{Rock, Scissors}}, [2]int{7, 3}},
		RoundTest{Round{[2]Choice{Rock, Rock}}, [2]int{4, 4}},
		RoundTest{Round{[2]Choice{Paper, Rock}}, [2]int{8, 1}},
		RoundTest{Round{[2]Choice{Paper, Scissors}}, [2]int{2, 9}},
		RoundTest{Round{[2]Choice{Paper, Paper}}, [2]int{5, 5}},
		RoundTest{Round{[2]Choice{Scissors, Rock}}, [2]int{3, 7}},
		RoundTest{Round{[2]Choice{Scissors, Paper}}, [2]int{9, 2}},
		RoundTest{Round{[2]Choice{Scissors, Scissors}}, [2]int{6, 6}},
	} {
		t.Run(test.Name(), test.Test)
	}
}

type ParseChoiceTest struct {
	s string
	c Choice
}

func TestParseChoice(t *testing.T) {
	for _, test := range []ParseChoiceTest{
		ParseChoiceTest{"a", Rock},
		ParseChoiceTest{"x", Rock},
		ParseChoiceTest{"b", Paper},
		ParseChoiceTest{"y", Paper},
		ParseChoiceTest{"c", Scissors},
		ParseChoiceTest{"z", Scissors},
		ParseChoiceTest{"A", Rock},
		ParseChoiceTest{"X", Rock},
		ParseChoiceTest{"B", Paper},
		ParseChoiceTest{"Y", Paper},
		ParseChoiceTest{"C", Scissors},
		ParseChoiceTest{"Z", Scissors},
	} {
		t.Run(fmt.Sprintf("%s %d", test.s, test.c), func(t *testing.T) {
			if c, err := ParseChoice(test.s); err != nil {
				t.Error(err.Error())
			} else if c != test.c {
				t.Errorf("Expected %d got %d", test.c, c)
			}
		})
	}
	t.Run("InvalidChoice", func(t *testing.T) {
		r, e := ParseChoice("d")
		if e == nil {
			t.Errorf("Expected error parsing `d`, got nil")
		}
		if r != Invalid {
			t.Errorf("Expected Invalid (%d) choice parsing `d`, got %d", Invalid, r)
		}
	})
}

func TestGame(t *testing.T) {
	var valid_input string = `A Y
B X
C Z`
	t.Run("ReadRounds", func(t *testing.T) {
		t.Run("Valid", func(t *testing.T) {
			game := new(Game)
			if err := game.ReadRounds(bytes.NewBufferString(valid_input)); err != nil {
				t.Errorf("Failed to ReadRounds with input %s: %s", valid_input, err.Error())
			}
		})
		t.Run("Invalid", func(t *testing.T) {
			t.Run("BadField", func(t *testing.T) {
				game := new(Game)
				if err := game.ReadRounds(bytes.NewBufferString("A T")); err == nil {
					t.Errorf("Expected error from invalid choice field, got nil")
				}
			})
			t.Run("MissingField", func(t *testing.T) {
				for _, str := range []string{" ", "A", "      Z     "} {
					t.Run(str, func(t *testing.T) {
						game := new(Game)
						if err := game.ReadRounds(bytes.NewBufferString(str)); err == nil {
							t.Errorf("Expected error from input '%s' missing fields, got nil and game is %v", str, game)
						} else {
							t.Logf("Got expected error from input '%s', %s", str, err.Error())
						}
					})
				}
			})
		})
	})
	t.Run("Scores", func(t *testing.T) {
		t.Run("Empty", func(t *testing.T) {
			var game Game
			expected := [2]int{0, 0}
			if scores := game.Scores(); scores != expected {
				t.Errorf("Game.Scores: expected %v, got %v", expected, scores)
			}
		})
		t.Run("Valid", func(t *testing.T) {
			var game Game
			if err := game.ReadRounds(bytes.NewBufferString(valid_input)); err != nil {
				fmt.Errorf("Failed to read valid input")
			} else {
				expected := [2]int{1 + 8 + 6, 15}
				if scores := game.Scores(); scores != expected {
					t.Errorf("Game.Scores: Expected %v, got %v", expected, scores)
				}
			}
		})
	})
}
