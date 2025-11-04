package frontend

import (
	"fmt"
	"unicode"
)

type TokenType string

const (
	// Literal Types
	Number TokenType = "Number"
	String TokenType = "String"
	Identifier TokenType = "Identifier"

	// keywords
	Let TokenType = "Let"
	Const TokenType = "Const"

	// Grouping * Opearators
	BinaryOperator TokenType = "BinaryOperator"
	DoubleQuote TokenType = "DoubleQuote"
	Equals TokenType = "Equals"
	NotEqual TokenType = "NotEqual"
	OpenParen TokenType = "OpenParen"
	CloseParen TokenType = "CloseParen"
	OpenBracket TokenType = "OpenBracket"
	CloseBracket TokenType = "CloseBracket"
	Comma TokenType = "Comma"
	GreaterThan TokenType = "GreaterThan"
	LessThan TokenType = "LessThan"
	GreaterThanOREqualTo TokenType = "GreaterThanOREqualTo"
	LessThanOrEqualTO TokenType = "LessThanOrEqualTO"
	EOF TokenType = "EOF"

	// Reserved keywords
	Print TokenType = "Print"
	Func TokenType = "Func"
	For TokenType = "For"
	While TokenType = "While"
	If TokenType = "If"
	Elif TokenType = "Elif"
	Else TokenType = "Else"
	Return TokenType = "Return"
)

type Token struct {
	Type TokenType;
	Value string
}

func newToken(tokenType TokenType, value string) Token {
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
		
		if isSkippable(char) {
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
	tokens = append(tokens, newToken(EOF, "EOF"))

	return tokens
}


