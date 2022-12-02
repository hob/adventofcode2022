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
	first := Elf{0, 0}
	second := Elf{0, 0}
	third := Elf{0, 0}
	current := Elf{0, 0}

	file, err := os.Open("day1/input.txt")
	defer file.Close()

	error.Check(err)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		currentText := scanner.Text()
		if currentText != "" {
			currentVal, err := strconv.Atoi(currentText)
			check(err)
			current.calories = current.calories + currentVal
		} else {
			println(fmt.Sprintf("Elf number %d has %d calories", current.number, current.calories))
			if current.calories > first.calories {
				if first.calories > 0 {
					second = first
				}
				if second.calories != 0 {
					third = second
				}
				first = current
			} else if current.calories > second.calories {
				if second.calories > 0 {
					third = second
				}
				second = current
			} else if current.calories > third.calories {
				third = current
			}
			current = Elf{current.number + 1, 0}
			println(fmt.Sprintf("created elf number %d", current.number))
		}
	}
	println(fmt.Sprintf("1st place (Elf number %d) has %d calories", first.number, first.calories))
	println(fmt.Sprintf("2nd place (Elf number %d) has %d calories", second.number, second.calories))
	println(fmt.Sprintf("3rd place (Elf number %d) has %d calories", third.number, third.calories))
	println(first.calories + second.calories + third.calories)
}
