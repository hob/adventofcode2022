package main

import (
	"advent22.spillane.farm/internal/util"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("cmd/day4/input.txt")
	defer file.Close()

	util.CheckErr(err)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	numFullyContained := 0
	for scanner.Scan() {
		assignments := strings.Split(scanner.Text(), ",")
		a1 := strings.Split(assignments[0], "-")
		a2 := strings.Split(assignments[1], "-")
		a1Start, _ := strconv.Atoi(a1[0])
		a1End, _ := strconv.Atoi(a1[1])
		a2Start, _ := strconv.Atoi(a2[0])
		a2End, _ := strconv.Atoi(a2[1])
		if (a1Start <= a2Start && a1End >= a2End) || (a2Start <= a1Start && a2End >= a1End) {
			numFullyContained++
		}
	}
	println(fmt.Sprintf("Number of assignments with full overlap: %d", numFullyContained))
}
