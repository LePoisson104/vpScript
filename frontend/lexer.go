package frontend

import (
	"fmt"
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
	NotEqual = "NOT_EQUAL"
	OpenParen = "OPEN_PAREN"
	CloseParen = "CLOSE_PAREN"
	OpenBracket = "OPEN_BRACKET"
	CloseBracket = "CLOSE_BRACKET"
	Comma = "COMMA"
	GreaterThan = "GREATER_THAN"
	LessThan = "LESS_THAN"
	GreaterThanOREqualTo = "GREATER_THAN_OR_EQUAL_TO"
	LessThanOrEqualTO = "LESS_THAN_OR_EQUAL_TO"
	EOF = "END_OF_FILE"

	// Reserved keywords
	Print = "PRINT"
	Func = "FUNC"
	For = "FOR"
	While = "WHILE"
	If = "IF"
	Elif = "ELIF"
	Else = "ELSE"
	Return = "RETURN"
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

func isSkippable(src rune) bool{ 
	return src == ' ' || src == '\n' || src == '\t'
}

func isInt(char rune) bool {
	return unicode.IsDigit(char)
}

func Tokenize(sourceCode string) []Token {
	tokens := []Token{}
	position := 0

	for position <= len(sourceCode) - 1 {
		char := rune(sourceCode[position])
		
		if position == len(sourceCode) - 1 {
			tokens = append(tokens, newToken(EOF, "EOF"))
		} else if isSkippable(char) {
			position++
			continue
		} else if char == '(' {
			tokens = append(tokens, newToken(OpenParen, "("))
		} else if char == ')' {
			tokens = append(tokens, newToken(CloseParen, ")"))
		} else if char == '{' {
			tokens = append(tokens, newToken(OpenBracket, "{"))
		} else if char == '}' {
			tokens = append(tokens, newToken(CloseBracket, "}"))
		} else if char == '=' {
			tokens = append(tokens, newToken(Equals, "="))
		} else if char == '!' {
			str := ""
			for position < len(sourceCode) - 1 {
				char := rune(sourceCode[position])
				if char == '=' {
					str += string(char)
					break
				}
				str += string(char)
			}
			tokens = append(tokens, newToken(NotEqual, str))
		} else if char == '<' {
			char := rune(sourceCode[position + 1])
			if char == ' ' {
				tokens = append(tokens, newToken(LessThan, "<"))
			} else if char == '=' {
				tokens = append(tokens, newToken(LessThanOrEqualTO, "<="))
				position++
			}
		} else if char == '>' {
			char := rune(sourceCode[position + 1])
			if char == ' ' {
				tokens = append(tokens, newToken(GreaterThan, ">"))
			} else if char == '=' {
				tokens = append(tokens, newToken(GreaterThanOREqualTo, ">="))
				position++
			}
		} else if char == '+' || char == '-' || char == '*' || char == '/' {
			tokens = append(tokens, newToken(BinaryOperator, string(char)))
		} else if char == ',' {
			tokens = append(tokens, newToken(Comma, ","))
		} else if char == '"' {
			str := ""
			position++
			for position <= len(sourceCode) - 1 {
				char := rune(sourceCode[position])
				if char == '"'{
					break
				} 
				str += string(char)
				position++
			}
			tokens = append(tokens, newToken(String, str))
		} else if isAlpha(char) {
			str := ""
			for position <= len(sourceCode) - 1 {
				char := rune(sourceCode[position])
				if char == ' ' || !isAlpha(char) { 
					position--
					break
				} 
				str += string(char)
				position++
			}
			
			if str == "let" {
				tokens = append(tokens, newToken(Let, str))
			} else if str == "const" {
				tokens = append(tokens, newToken(Const, str))
			} else if str == "print" {
				tokens = append(tokens, newToken(Print, str))
			} else if str == "func" { 
				tokens = append(tokens, newToken(Func, str))
			} else if str == "for" { 
				tokens = append(tokens, newToken(For, str))
			} else if str == "while" { 
				tokens = append(tokens, newToken(While, str))
			} else if str == "if" { 
				tokens = append(tokens, newToken(If, str))
			} else if str == "elif" { 
				tokens = append(tokens, newToken(Elif, str))
			} else if str == "else" { 
				tokens = append(tokens, newToken(Else, str))
			} else if str == "return" {
				tokens = append(tokens, newToken(Return, str))
			} else {
				tokens = append(tokens, newToken(Identifier, str))
			}
		} else if isInt(char) {
			str := ""
			for position <= len(sourceCode) - 1 {
				char := rune(sourceCode[position])
				if !isInt(char) && char != '.' {
					position--
					break
				}
				str += string(char)
				position++
			}
			tokens = append(tokens, newToken(Number, str))
		} else if char == ';' {
			position++
			continue
		} else {
			fmt.Printf("Unrecognized character found in source code '%c'\n", char)
			break
		}
		position++
	}
	return tokens
}


