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
		c := findCharsInCommon(compartment1, compartment2)
		if strings.Contains(p1, c) {
			p := strings.Index(p1, c) + 1 + 26
			totalPriority += p
		} else {
			p := strings.Index(p2, c) + 1
			totalPriority += p
		}
	}
	println(fmt.Sprintf("Total Priority: %d", totalPriority))
}

func findCharsInCommon(compartment1 string, compartment2 string) string {
	for _, c := range compartment1 {
		s := string(c)
		if strings.Contains(compartment2, s) {
			return s
		}
	}
	return ""
}
