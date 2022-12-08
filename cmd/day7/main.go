package main

import (
	"advent22.spillane.farm/internal/util"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	largeFileSizeThreshold = 100000
	spaceRequiredForUpdate = 30000000
	totalDiskSpace         = 70000000
)

type File struct {
	Size       int
	Name       string
	Parent     *File
	Children   []*File
	IsDir      bool
	nestedSize int
}

func (f *File) CalcSize() int {
	if f.nestedSize != 0 {
		return f.nestedSize
	}

	sum := 0
	for _, c := range f.Children {
		if c.IsDir {
			sum += c.CalcSize()
		} else {
			sum += c.Size
		}
	}

	f.nestedSize = sum
	return sum
}

func main() {
	file, err := os.Open("cmd/day7/input.txt")
	defer file.Close()

	util.CheckErr(err)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	root := &File{
		Size:     0,
		Parent:   nil,
		Children: nil,
		Name:     "/",
		IsDir:    true,
	}
	currentDir := root

	scanner.Scan() //First line is root which is already created so skip it
	i := 0
	for scanner.Scan() {
		i++
		println(i)
		line := scanner.Text()
		if line[0:1] == "$" { // command
			command := line[2:4]
			switch command {
			case "cd":
				dirName := line[5:]
				if dirName == ".." {
					currentDir = currentDir.Parent
				} else {
					for _, child := range currentDir.Children {
						if child.Name == dirName {
							currentDir = child
							break
						}
					}
				}
				break
			case "ls":
				//Don't do anything except move to the results of ls on the following lines
				break
			}
		} else if line[0:3] == "dir" {
			dir := &File{
				Name:     line[4:],
				Size:     0,
				Parent:   currentDir,
				Children: nil,
				IsDir:    true,
			}
			currentDir.Children = append(currentDir.Children, dir)
		} else { //It must be a file
			split := strings.Split(line, " ")
			size, _ := strconv.Atoi(split[0])
			file := &File{
				Size:     size,
				Name:     split[1],
				Parent:   currentDir,
				Children: nil,
				IsDir:    false,
			}
			currentDir.Children = append(currentDir.Children, file)
		}
	}
	//smallDirs := findSmallDirs(root)
	//println(fmt.Sprintf("total small dir size bytes: %d", smallDirs))
	freeSpace := totalDiskSpace - root.CalcSize()
	println(fmt.Sprintf("total free space: %d", freeSpace))
	spaceToFreeUp := spaceRequiredForUpdate - freeSpace
	println(fmt.Sprintf("space needed to free up: %d", spaceToFreeUp))
	closestMatch := root
	walk(root, func(dir *File) {
		size := dir.CalcSize()
		if size > spaceToFreeUp && size < closestMatch.CalcSize() {
			closestMatch = dir
		}
	})
	println(fmt.Sprintf("closest match is %d bytes", closestMatch.CalcSize()))
}

func findSmallDirs(dir *File) int {
	sum := 0
	if size := dir.CalcSize(); size <= largeFileSizeThreshold {
		sum += size
	}
	for _, file := range dir.Children {
		sum += findSmallDirs(file)
	}
	return sum
}

func walk(dir *File, callback func(dir *File)) {
	callback(dir)
	for _, file := range dir.Children {
		if file.IsDir {
			walk(file, callback)
		}
	}
}
