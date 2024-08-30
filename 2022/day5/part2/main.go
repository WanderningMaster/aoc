package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/WanderningMaster/aoc/internal/utils"
)

type TokenType = string

const (
	OpenBracket  TokenType = "OPEN_BRACKET"
	CloseBracket TokenType = "CLOSE_BRACKET"
	Space        TokenType = "SPACE"
	Empty        TokenType = "EMPTY"
	NewLine      TokenType = "NEW_LINE"
	Ident        TokenType = "IDENT"
	Illegal      TokenType = "ILLEGAL"
	Eof          TokenType = "EOF"
)

type Token struct {
	type_   TokenType
	literal string
}

type Tokenizer struct {
	pos   int
	ch    rune
	input string
}

func (t *Tokenizer) readChar() {
	if t.pos >= len(t.input) {
		t.ch = '\x00'
	} else {
		t.ch = rune(t.input[t.pos])
	}
	t.pos += 1
}

func NewTokenizer(input string) Tokenizer {
	var tokenizer = Tokenizer{
		pos:   0,
		input: input,
	}
	tokenizer.readChar()
	return tokenizer
}

func (t *Tokenizer) getNextToken() Token {
	var tok Token
	var tokNil bool = true

	switch t.ch {
	case '[':
		tok = CreateToken(OpenBracket, string(t.ch))
		tokNil = false
	case ']':
		tok = CreateToken(CloseBracket, string(t.ch))
		tokNil = false
	case ' ':
		if t.input[t.pos] == ' ' {
			tok = CreateToken(Empty, "empty")
			t.pos += 3
			tokNil = false
		} else {
			tok = CreateToken(Space, string(t.ch))
			tokNil = false
		}
	case '\n':
		tok = CreateToken(NewLine, string(t.ch))
		tokNil = false
	case '\x00':
		tok = CreateToken(Eof, "eof")
		tokNil = false
	}

	if isLiteral(t.ch) {
		tok = CreateToken(Ident, string(t.ch))
	} else if tokNil {
		tok = CreateToken(Illegal, string(t.ch))
	}
	t.readChar()

	return tok
}

const a = int('a')
const z = int('z')

const A = int('A')
const Z = int('Z')

func CreateToken(type_ TokenType, literal string) Token {
	return Token{type_: type_, literal: literal}
}

func isLiteral(ch rune) bool {
	assCode := int(ch)
	return a <= assCode && z >= assCode ||
		A <= assCode && Z >= assCode
}

func tokenize(input string) []Token {
	lexer := NewTokenizer(input)

	var tokens []Token
	for {
		token := lexer.getNextToken()
		if token.type_ == Eof {
			break
		}
		tokens = append(tokens, token)
	}

	return tokens
}

func parseStacks(input []string) []*utils.Stack {
	in := strings.Join(input, "\n")
	fmt.Println(in)
	in += "\x00"

	tokens := tokenize(in)
	var stacks []*utils.Stack
	stack_id := 0
	for _, token := range tokens {
		switch token.type_ {
		case Ident:
			{
				if ok := stack_id < len(stacks); !ok {
					stacks = append(stacks, utils.NewStack())
				}
				stacks[stack_id].Push(token.literal)
			}
		case Space:
			stack_id += 1
		case Empty:
			stack_id += 1
		case NewLine:
			stack_id = 0
		}
	}

	return stacks
}

func separateStacksFromCommands(input string) ([]string, []string) {
	scanner := bufio.NewScanner(strings.NewReader(input))

	stacks := []string{}
	commands := []string{}

	commandSection := false
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			commandSection = true
			continue
		}

		if commandSection {
			commands = append(commands, line)
		} else {
			stacks = append(stacks, line)
		}
	}

	slices.Reverse(stacks)
	return stacks[1:], commands
}

func extractCommand(command string) (int, int, int) {
	var cnt, from, to int
	idents := strings.Split(command, " ")
	cnt, _ = strconv.Atoi(idents[1])
	from, _ = strconv.Atoi(idents[3])
	to, _ = strconv.Atoi(idents[5])
	return cnt, from, to
}

func move(cnt int, from int, to int, stacks []*utils.Stack) {
	composedCrates := []string{}
	for cnt > 0 {
		crateToMove, _ := stacks[from-1].Pop()
		composedCrates = append(composedCrates, crateToMove.(string))
		cnt -= 1
	}
	slices.Reverse(composedCrates)
	for _, crate := range composedCrates {
		stacks[to-1].Push(crate)
	}
}

func moveCrates(commands []string, stacks []*utils.Stack) {
	for _, command := range commands {
		cnt, from, to := extractCommand(command)
		move(cnt, from, to, stacks)
	}
}

func printResult(stacks []*utils.Stack) {
	result := ""
	for _, stack := range stacks {
		data, _ := stack.Pop()
		result += data.(string)
	}
	fmt.Printf("Result: %v\n", result)
}

func main() {
	dirname, err := utils.Dirname()
	if err != nil {
		log.Fatalln(err)
	}

	input, err := os.ReadFile(dirname + "/2022/day5/in.txt")

	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	stacks, commands := separateStacksFromCommands(string(input))
	parsedStacks := parseStacks(stacks)
	moveCrates(commands, parsedStacks)
	printResult(parsedStacks)
}
