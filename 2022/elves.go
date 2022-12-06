package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Elf struct {
	calories []int64
}

func (e *Elf) String() string {
	return fmt.Sprintf("Elf{Calories %v Total %d}", e.calories, e.TotalCalories())
}

func (e *Elf) AddCalories(is ...int64) {
	e.calories = append(e.calories, is...)
}

func (e *Elf) TotalCalories() (t int64) {
	if e.calories == nil {
		return
	}
	for _, c := range e.calories {
		t += c
	}
	return
}

type ElvesByCalories []*Elf

func (el ElvesByCalories) Less(i, j int) bool {
	return el[i].TotalCalories() < el[j].TotalCalories()
}
func (el ElvesByCalories) Swap(i, j int) {
	el[i], el[j] = el[j], el[i]
}
func (el ElvesByCalories) Len() int {
	return len(el)
}

func (el ElvesByCalories) TotalCalories() (t int64) {
	for _, e := range el {
		t += e.TotalCalories()
	}
	return
}

func ReadCalorieList(in io.Reader) []*Elf {
	fileScanner := bufio.NewScanner(in)
	var e *Elf
	var elves []*Elf
	for fileScanner.Scan() {
		var cal int64
		line := strings.TrimSpace(fileScanner.Text())
		if line == "" {
			// new elf
			if e != nil {
				elves = append(elves, e)
				e = nil
			}
			continue
		}
		if c, err := strconv.ParseInt(line, 10, 64); err != nil {
			panic(fmt.Errorf("Expected blank lines or parseable integers; line '%s' was neither", line))
		} else {
			cal = c
		}
		if e == nil {
			e = new(Elf)
		}
		e.AddCalories(cal)
	}
	if e != nil {
		elves = append(elves, e)
	}
	return elves
}
