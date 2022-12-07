package common

import (
	"aoc2022/util"
	"strings"
)

const RootName = "/"

type FileNode struct {
	name     string
	size     int
	parent   *FileNode
	children []*FileNode
}

func ParseLinesToFiletree(lines []string) *FileNode {
	const prefixCmd = "$ "

	rootDir := FileNode{RootName, 0, nil, []*FileNode{}}
	var currentDir *FileNode = nil

	for _, line := range lines {
		if strings.HasPrefix(line, prefixCmd) {
			currentDir = ParseCmd(line[len(prefixCmd):], &rootDir, currentDir)
		} else {
			ParseLsOutput(line, currentDir)
		}
	}
	return &rootDir
}

func ParseCmd(cmd string, rootDir *FileNode, currentDir *FileNode) *FileNode {
	const cmdCd = "cd "
	const parentDir = ".."

	if strings.HasPrefix(cmd, cmdCd) {
		targetDir := cmd[len(cmdCd):]
		if targetDir == RootName {
			return rootDir
		} else if targetDir == parentDir {
			return currentDir.parent
		}
		for _, child := range currentDir.children {
			if child.name == targetDir {
				return child
			}
		}
	}
	return currentDir
}

func ParseLsOutput(lsOutput string, currentDir *FileNode) {
	const prefixDir = "dir "

	if strings.HasPrefix(lsOutput, prefixDir) {
		dirName := lsOutput[len(prefixDir):]
		currentDir.children = append(currentDir.children, &FileNode{dirName, 0, currentDir, []*FileNode{}})
	} else {
		lsParts := strings.Split(lsOutput, " ")
		currentDir.children = append(currentDir.children, &FileNode{lsParts[1], util.StringToNum(lsParts[0]), currentDir, []*FileNode{}})
	}
}

func FindDirsSizes(fileNode *FileNode, dirsSizes *map[*FileNode]int) int {
	if len(fileNode.children) <= 0 {
		return 0
	}

	sum := 0
	for _, child := range fileNode.children {
		sum += child.size
		sum += FindDirsSizes(child, dirsSizes)
	}
	(*dirsSizes)[fileNode] = sum
	return sum
}
