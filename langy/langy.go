package main

import (
	"fmt"
)

type token struct {
	kind  string
	value string
}

var current = 0

func tokenize(input string) []token {

	var tokens []token

	for current < len([]rune(input)) {

		char := string([]rune(input)[current])

		if char == "(" {
			tokens = append(tokens, token{kind: "paren", value: "("})
			current++
			continue
		}
		if char == ")" {
			tokens = append(tokens, token{kind: "paren", value: ")"})
			current++
			continue
		}

		if char == " " {
			current++
			continue
		}

		if []rune(input)[current] >= 'a' && []rune(input)[current] <= 'z' {
			tokens = append(tokens, token{kind: "name", value: getName(input)})
		}

		if []rune(input)[current] >= '0' && []rune(input)[current] <= '9' {
			tokens = append(tokens, token{kind: "number", value: getNumber(input)})
		}

	}

	return tokens
}

func getName(input string) string {
	temp := ""
	for []rune(input)[current] >= 'a' && []rune(input)[current] <= 'z' {

		temp += string([]rune(input)[current])
		current++
	}
	return temp
}

func getNumber(input string) string {
	temp := ""
	for []rune(input)[current] >= '0' && []rune(input)[current] <= '9' {

		temp += string([]rune(input)[current])
		current++
	}
	return temp
}

type node struct {
	kind   string
	name   string
	value  string
	params []node
	body   []node
}

var pc int

func parse(input []token) node {

	pc = 0

	var ast = node{
		kind: "Program",
		body: []node{},
	}

	for pc < len(input) {
		ast.body = append(ast.body, walk(input))
	}

	return ast
}

func walk(input []token) node {

	if input[pc].kind == "number" {
		n := node{
			kind:  "NumberLiteral",
			value: input[pc].value,
		}
		pc++
		return n
	}

	if input[pc].kind == "paren" && input[pc].value == "(" {
		pc++

		n := node{
			kind:   "CallExpression",
			name:   input[pc].value,
			params: []node{},
		}

		pc++

		for input[pc].kind != "paren" || (input[pc].kind == "paren" && input[pc].value != ")") {
			n.params = append(n.params, walk(input))
		}
		pc++ //skip closing

		return n

	}

	// if input[pc].kind == "name" {
	// 	fmt.Println("pc: ", pc)
	// 	n := node{
	// 		kind: "CallExpression",
	// 		name: input[pc].value,
	// 	}
	// 	pc++
	// 	n.params = append(n.params, walk(input))
	// 	return n
	// }

	// pc++
	return node{}
}

func main() {
	program := "(add 10 (subtract 7 6))"

	tokens := tokenize(program)
	fmt.Println("tokens")
	fmt.Println(tokens)

	ast := parse(tokens)
	fmt.Println("ast")
	fmt.Printf("%+v\n", ast)
	// fmt.Println(ast)
	// fmt.Println(json.MarshalIndent(ast, "", "    "))
}
