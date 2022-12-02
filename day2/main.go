package main

import (
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

var plays = map[string]Roshamboable{
	"A": Rock{},
	"B": Paper{},
	"C": Scissors{},
}
var responses = map[string]Roshamboable{
	"X": Rock{},
	"Y": Paper{},
	"Z": Scissors{},
}

func main() {
	file, err := os.Open("day2/input.txt")
	defer file.Close()

	check(err)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	score := 0
	for scanner.Scan() {
		currentText := scanner.Text()
		split := strings.Split(currentText, " ")
		play := plays[split[0]]
		response := responses[split[1]]
		score += response.GetValue()

		result := play.Winner(response)
		if result == response {
			score += winValue
		} else if result == play {
			score += lossValue
		} else {
			score += tieValue
		}
	}
	println(fmt.Sprintf("Final score: %d", score))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
