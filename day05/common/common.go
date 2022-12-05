package common

import (
	"aoc2022/util"
	"strings"
)

type Stack []string

func ParseStacks(lines []string) []Stack {
	lastLineIdx := len(lines) - 1
	lastLine := lines[lastLineIdx]
	numStacks := len(strings.Split(lastLine, " "))
	stacks := make([]Stack, numStacks)

	for i := range lines[:lastLineIdx] {
		line := lines[lastLineIdx-1-i]

		remainder := line
		stackIdx := 0
		for len(remainder) > 0 {
			element := remainder[:3]
			remainder = remainder[3:]
			remainder = strings.TrimPrefix(remainder, " ")
			elementName := element[1]
			if elementName != ' ' {
				stacks[stackIdx] = append(stacks[stackIdx], string(elementName))
			}
			stackIdx++
		}
	}
	return stacks
}

type Command struct {
	amount int
	source int
	target int
}

func ParseCommands(lines []string) []Command {
	var commands []Command
	for _, line := range lines {
		lineParts := strings.Split(line, " ")
		commands = append(commands, Command{util.StringToNum(lineParts[1]), util.StringToNum(lineParts[3]) - 1, util.StringToNum(lineParts[5]) - 1})
	}
	return commands
}

func ApplyCommands(stacks []Stack, commands []Command, stackMoveStrategy func(Stack, Stack) Stack) {
	for _, cmd := range commands {
		srcStack := stacks[cmd.source]
		stacks[cmd.source] = srcStack[:len(srcStack)-cmd.amount]
		mvStack := srcStack[len(srcStack)-cmd.amount:]
		targetStack := stacks[cmd.target]
		targetStack = stackMoveStrategy(targetStack, mvStack)
		stacks[cmd.target] = targetStack
	}
}

func ParseInput(lines []string) ([]Stack, []Command) {
	breakIdx := -1
	for idx, line := range lines {
		if line == "" {
			breakIdx = idx
			break
		}
	}

	stacksLines := lines[:breakIdx]
	stacks := ParseStacks(stacksLines)

	commandLines := lines[breakIdx+1:]
	commands := ParseCommands(commandLines)
	return stacks, commands
}

func GetStacksTops(stacks []Stack) string {
	stacksTops := ""
	for _, stack := range stacks {
		if len(stack) > 0 {
			stacksTops += stack[len(stack)-1]
		}
	}
	return stacksTops
}
