package main

import (
	"advent22.spillane.farm/internal/util"
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strings"
)

const (
	lossValue = 0
	tieValue  = 3
	winValue  = 6
)

type Roshamboable interface {
	// Winner Returns compares to input & returns the winner.  If both are equal, returns nil
	Winner(comp Roshamboable) Roshamboable
	WinsAgainst() Roshamboable
	LosesTo() Roshamboable
	GetValue() int
}

type Rock struct{}
type Paper struct{}
type Scissors struct{}

func (r Rock) Winner(comp Roshamboable) Roshamboable {
	if reflect.TypeOf(comp) == reflect.TypeOf(Scissors{}) {
		return r
	}
	if reflect.TypeOf(comp) == reflect.TypeOf(Paper{}) {
		return comp
	}
	return nil
}

func (p Paper) Winner(comp Roshamboable) Roshamboable {
	if reflect.TypeOf(comp) == reflect.TypeOf(Rock{}) {
		return p
	}
	if reflect.TypeOf(comp) == reflect.TypeOf(Scissors{}) {
		return comp
	}
	return nil
}

func (s Scissors) Winner(comp Roshamboable) Roshamboable {
	if reflect.TypeOf(comp) == reflect.TypeOf(Paper{}) {
		return s
	}
	if reflect.TypeOf(comp) == reflect.TypeOf(Rock{}) {
		return comp
	}
	return nil
}

func (r Rock) GetValue() int {
	return 1
}

func (p Paper) GetValue() int {
	return 2
}

func (s Scissors) GetValue() int {
	return 3
}

func (r Rock) WinsAgainst() Roshamboable {
	return Scissors{}
}

func (r Rock) LosesTo() Roshamboable {
	return Paper{}
}

func (p Paper) WinsAgainst() Roshamboable {
	return Rock{}
}

func (p Paper) LosesTo() Roshamboable {
	return Scissors{}
}

func (s Scissors) WinsAgainst() Roshamboable {
	return Paper{}
}

func (s Scissors) LosesTo() Roshamboable {
	return Rock{}
}

var plays = map[string]Roshamboable{
	"A": Rock{},
	"B": Paper{},
	"C": Scissors{},
}
var results = map[string]int{
	"X": -1, //lose
	"Y": 0,  //tie
	"Z": 1,  //win
}

func main() {
	file, err := os.Open("cmd/day2/input.txt")
	defer file.Close()

	util.CheckErr(err)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	score := 0
	for scanner.Scan() {
		currentText := scanner.Text()
		split := strings.Split(currentText, " ")
		play := plays[split[0]]
		result := results[split[1]]
		switch result {
		case -1:
			score += play.WinsAgainst().GetValue()
			score += lossValue
			break
		case 0:
			score += play.GetValue()
			score += tieValue
			break
		case 1:
			score += play.LosesTo().GetValue()
			score += winValue
			break
		}
	}
	println(fmt.Sprintf("Final score: %d", score))
}
