package main

import (
	"advent22.spillane.farm/internal/util"
	"bufio"
	"fmt"
	"os"
	"strings"
)

var p1 = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
var p2 = "abcdefghijklmnopqrstuvwxyz"

func main() {
	findBadges()
}

func findBadges() {
	file, err := os.Open("cmd/day3/input.txt")
	defer file.Close()

	util.CheckErr(err)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	group := make([]string, 0)
	totalPriority := 0
	for scanner.Scan() {
		group = append(group, scanner.Text())
		if len(group) == 3 {
			s := findCharsInCommon(group)
			totalPriority += getItemTypePriority(s)
			//Reset for the next group
			group = make([]string, 0)
		}
	}
	println(fmt.Sprintf("Total Priority: %d", totalPriority))
}

func findMisplacedItems() {
	file, err := os.Open("cmd/day3/input.txt")
	defer file.Close()

	util.CheckErr(err)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	totalPriority := 0
	for scanner.Scan() {
		contents := scanner.Text()
		compartment1 := contents[0 : len(contents)/2]
		compartment2 := contents[len(contents)/2 : len(contents)]
		c := findCharsInCommon([]string{compartment1, compartment2})
		totalPriority += getItemTypePriority(c)
	}
	println(fmt.Sprintf("Total Priority: %d", totalPriority))
}

func getItemTypePriority(t string) int {
	if strings.Contains(p1, t) {
		return strings.Index(p1, t) + 1 + 26
	} else {
		return strings.Index(p2, t) + 1
	}
}

func findCharsInCommon(lists []string) string {
	for _, c := range lists[0] {
		theRest := lists[1:len(lists)]
		for i, s := range theRest {
			if !strings.Contains(s, string(c)) {
				break
			}
			if i == len(theRest)-1 {
				return string(c)
			}
		}
	}
	return ""
}
