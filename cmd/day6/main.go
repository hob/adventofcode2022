package main

import (
	"advent22.spillane.farm/internal/util"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("cmd/day6/input.txt")
	defer file.Close()

	util.CheckErr(err)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	line := scanner.Text()
	i := 0
	for !allUniqueChars(line[i : i+14]) {
		i++
	}
	println(fmt.Sprintf("Uniquness occurred at character %d with characters %s", i+14, line[i:i+14]))
}

func allUniqueChars(s string) bool {
	for i := 0; i < len(s); i++ {
		if strings.Count(s, string(s[i])) > 1 {
			return false
		}
	}
	return true
}
