package main

import (
	"fmt"
	"strings"
	"unicode"
)

const (
	// Literal Types
	Number = "NUMBER"
	String = "STRING"
	Identifier = "IDENTIFIER"

	// keywords
	Let = "LET"
	Const = "CONST"

	// Grouping * Opearators
	BinaryOperator = "BINARY_OPERATOR"
	DoubleQuote = "DoubleQuote"
	Equals = "EQUALS"
	OpenParen = "OPEN_PAREN"
	CloseParen = "CLOSE_PAREN"
	OpenBracket = "OPEN_BRACKET"
	CloseBracket = "CLOSE_BRACKET"
)

type Token struct {
	Type string;
	Value string
}

func newToken(tokenType string, value string) Token {
	return Token{
		Type: tokenType,
		Value: value,
	}
}

func isAlpha(char rune) bool{
	return unicode.ToUpper(char) != unicode.ToLower(char)
}

func isSkippable(src string) bool{
	return src == " " || src == "\n" || src == "\t"
}

func isInt(char rune) bool {
	return unicode.IsDigit(char)
}

func tokenize(sourceCode string) []Token {
	tokens := []Token{}
	src := strings.Split(sourceCode, " ")

	for _, char := range src {
		if isSkippable(char) {
			fmt.Println("skip")
			continue
		} else if char == "(" {
			tokens = append(tokens, newToken(OpenParen, "("))
		} else if char == ")" {
			tokens = append(tokens, newToken(CloseParen, ")"))
		} else if char == "{" {
			tokens = append(tokens, newToken(OpenBracket, "{"))
		} else if char == "}" {
			tokens = append(tokens, newToken(CloseBracket, "}"))
		} else if char == `"` {
			tokens = append(tokens, newToken(DoubleQuote, `"`))
		} else if char == "+" || char == "-" || char == "*" || char == "/" {
			tokens = append(tokens, newToken(BinaryOperator, char))
		} else if isAlpha([]rune(char)[0]) {
			str := ""
			i := 0
			for i < len(char){
    			str += string(char[i])
				i++
			}
			if str == "let" {
				tokens = append(tokens, newToken(Let, str))
			} else if str == "const" {
				tokens = append(tokens, newToken(Const, str))
			} else {
				tokens = append(tokens, newToken(Identifier, str))
			}
		} else if isInt([]rune(char)[0]){
			numStr := ""
			i := 0
			for i < len(char){
				numStr += string(char[i])
				i++
			}
			tokens = append(tokens, newToken(Number, numStr))
		} else {
			fmt.Printf("Unrecognized character found in source code '%s'\n", char)
			break
		}
	}
	return tokens
}

func main() {
	fmt.Println(tokenize(`let string = "hello world"`))
}
