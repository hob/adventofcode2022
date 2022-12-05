package main

import (
	"advent22.spillane.farm/internal/util"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("cmd/day5/input.txt")
	defer file.Close()

	util.CheckErr(err)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	//Init stacks
	stacks := make([][]string, 0)
	for i := 0; i < 9; i++ {
		stacks = append(stacks, make([]string, 0))
	}
	i := 0
	for scanner.Scan() {
		if i < 8 {
			stacks = addToStacks(scanner.Text(), stacks)
		}
		if i > 9 {
			line := scanner.Text()
			split := strings.Split(line, " ")
			move, _ := strconv.Atoi(split[1])
			from, _ := strconv.Atoi(split[3])
			to, _ := strconv.Atoi(split[5])
			fromStack := stacks[from-1]
			items := fromStack[:move]
			//Remove from the "from" stack
			stacks[from-1] = fromStack[move:]
			//Add to the to stack
			copy := cp(items)
			stacks[to-1] = append(copy, stacks[to-1]...)
		}
		i++
	}
	for i := 0; i < 9; i++ {
		print(stacks[i][0])
	}
}

func cp(s []string) []string {
	reversed := make([]string, 0)
	for i := 0; i < len(s); i++ {
		reversed = append(reversed, s[i])
	}

	return reversed
}

func addToStacks(line string, stacks [][]string) [][]string {
	start := 0
	for i := 0; i < 9; i++ {
		stack := stacks[i]
		item := line[start+1 : start+2]
		if strings.Trim(item, " ") != "" {
			stacks[i] = append(stack, item)
		}
		start += 4
	}
	return stacks
}
