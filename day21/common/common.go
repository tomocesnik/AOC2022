package common

import (
	"aoc2022/util"
	"strings"
)

type ConstantNode struct {
	Value int
}

type FunctionNode struct {
	Arg1     string
	Arg2     string
	Operator string
}

func (n ConstantNode) CalcValue(expressions map[string]Expression) int {
	return n.Value
}

func (n FunctionNode) CalcValue(expressions map[string]Expression) int {
	arg1Node := expressions[n.Arg1]
	arg2Node := expressions[n.Arg2]
	switch n.Operator {
	case "+":
		return arg1Node.CalcValue(expressions) + arg2Node.CalcValue(expressions)
	case "-":
		return arg1Node.CalcValue(expressions) - arg2Node.CalcValue(expressions)
	case "*":
		return arg1Node.CalcValue(expressions) * arg2Node.CalcValue(expressions)
	case "/":
		return arg1Node.CalcValue(expressions) / arg2Node.CalcValue(expressions)
	}
	return 0
}

type Expression interface {
	CalcValue(expressions map[string]Expression) int
}

func ParseExpressions(lines []string) map[string]Expression {
	var expressionsMap map[string]Expression = make(map[string]Expression)
	for _, line := range lines {
		name, expr := ParseExpression(line)
		expressionsMap[name] = expr
	}
	return expressionsMap
}

func ParseExpression(line string) (string, Expression) {
	parts := strings.Split(line, ": ")
	varName := parts[0]
	expression := parts[1]
	eParts := strings.Split(expression, " ")
	if len(eParts) >= 3 {
		return varName, FunctionNode{Arg1: eParts[0], Arg2: eParts[2], Operator: eParts[1]}
	}
	return varName, ConstantNode{Value: util.StringToNum(eParts[0])}
}
