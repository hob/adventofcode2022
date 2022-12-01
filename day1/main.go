package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Elf struct {
	number   int
	calories int
}

func main() {
	max := Elf{0, 0}
	current := Elf{0, 0}

	file, err := os.Open("day1/input1.txt")
	defer file.Close()

	check(err)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		currentText := scanner.Text()
		if currentText != "" {
			currentVal, err := strconv.Atoi(currentText)
			check(err)
			println(fmt.Sprintf("current value = %d", currentVal))
			current.calories = current.calories + currentVal
		} else {
			if current.calories > max.calories {
				max = current
			}
			current = Elf{current.number + 1, 0}
			println(fmt.Sprintf("created elf number %d", current.number))
		}
	}
	println(fmt.Sprintf("Elf number %d has %d calories", max.number, max.calories))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
