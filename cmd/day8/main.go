package main

import (
	"advent22.spillane.farm/internal/util"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("cmd/day8/input.txt")
	defer file.Close()
	util.CheckErr(err)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	forest := buildForest(scanner)
	totalVisible := 0
	maxScenicScore := 0
	for y, row := range forest {
		for x, tree := range row {
			if isVisible(tree, x, y, forest) {
				totalVisible++
			}
			score := calculateScenicScore(tree, x, y, forest)
			if score > maxScenicScore {
				maxScenicScore = score
			}
		}
	}
	println(fmt.Sprintf("Total visible trees: %d", totalVisible))
	println(fmt.Sprintf("Max scenic score: %d", maxScenicScore))
}

func calculateScenicScore(tree int, x int, y int, forest [][]int) int {
	//If they're on the perimeter then viewing distance is always 0 since anything * 0 == 0
	if x == 0 || x == len(forest)-1 || y == 0 || y == len(forest[0])-1 {
		return 0
	}
	northScore := 0
	for i := y - 1; i >= 0; i-- {
		north := forest[i][x]
		northScore++
		if north >= tree {
			break
		}
	}
	southScore := 0
	for i := y + 1; i < len(forest); i++ {
		south := forest[i][x]
		southScore++
		if south >= tree {
			break
		}
	}
	westScore := 0
	for i := x - 1; i >= 0; i-- {
		west := forest[y][i]
		westScore++
		if west >= tree {
			break
		}
	}
	eastScore := 0
	for i := x + 1; i < len(forest[y]); i++ {
		east := forest[y][i]
		eastScore++
		if east >= tree {
			break
		}
	}
	return northScore * southScore * westScore * eastScore
}

func isVisible(tree int, x int, y int, forest [][]int) bool {
	//If they're on the perimeter then return true
	if x == 0 || x == len(forest)-1 || y == 0 || y == len(forest[0])-1 {
		return true
	}
	for i := y - 1; i >= 0; i-- {
		north := forest[i][x]
		if north >= tree {
			break
		}
		if i == 0 {
			return true
		}
	}
	for i := y + 1; i < len(forest); i++ {
		south := forest[i][x]
		if south >= tree {
			break
		}
		if i == len(forest)-1 {
			return true
		}
	}
	for i := x - 1; i >= 0; i-- {
		west := forest[y][i]
		if west >= tree {
			break
		}
		if i == 0 {
			return true
		}
	}
	for i := x + 1; i < len(forest[y]); i++ {
		east := forest[y][i]
		if east >= tree {
			break
		}
		if i == len(forest[y])-1 {
			return true
		}
	}
	return false
}

func buildForest(scanner *bufio.Scanner) [][]int {
	forest := make([][]int, 0)
	for scanner.Scan() {
		text := scanner.Text()
		row := make([]int, 0)
		for _, c := range text {
			i, _ := strconv.Atoi(string(c))
			row = append(row, i)
		}
		forest = append(forest, row)
	}
	return forest
}
