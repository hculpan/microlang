package lexer

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func isDigit(r rune) bool {
	return unicode.IsDigit(r)
}

func isLetter(r rune) bool {
	return unicode.IsLetter(r) || r == '_'
}

func isOperator(r rune) bool {
	return strings.ContainsRune("+-*/(){}=!;/.", r)
}

func Tokenize(input string) ([]Token, error) {
	var tokens []Token
	i := 0
	line := 1
	position := 1
	keywords := map[string]bool{
		"print":  true,
		"for":    true,
		"int":    true,
		"float":  true,
		"string": true,
		"if":     true,
	}

	for i < len(input) {
		r := rune(input[i])
		if r == '\n' {
			i++
			line++
			position = 1
		} else if unicode.IsSpace(r) {
			i++
			position++
		} else if isDigit(r) {
			start := i
			i++
			for i < len(input) && isDigit(rune(input[i])) {
				i++
			}
			if i < len(input) && rune(input[i]) == '.' {
				i++
				for i < len(input) && isDigit(rune(input[i])) {
					i++
				}
				value, _ := strconv.ParseFloat(input[start:i], 64)
				tokens = append(tokens, FloatToken{Value: value, Line: line, Position: position})
			} else {
				value, _ := strconv.ParseInt(input[start:i], 10, 64)
				tokens = append(tokens, IntegerToken{Value: value, Line: line, Position: position})
			}
			position += i - start
		} else if isLetter(r) || r == '"' {
			start := i
			if r == '"' {
				i++
				for i < len(input) && rune(input[i]) != '"' {
					i++
				}
				if i >= len(input) {
					return nil, fmt.Errorf("unterminated string literal at line %d, position %d", line, position)
				}
				i++
				tokens = append(tokens, StringToken{Value: input[start:i], Line: line, Position: position})
			} else {
				i++
				for i < len(input) && isLetter(rune(input[i])) {
					i++
				}
				word := input[start:i]
				if keywords[word] {
					tokens = append(tokens, KeywordToken{Value: word, Line: line, Position: position})
				} else {
					tokens = append(tokens, IdentifierToken{Value: word, Line: line, Position: position})
				}
			}
			position += i - start
		} else if isOperator(r) {
			start := i
			i++
			if i < len(input) {
				nextR := rune(input[i])
				if (r == '+' || r == '-' || r == '=') && nextR == '=' {
					i++
				} else if r == '!' && nextR == '=' {
					i++
				} else if r == '/' && nextR == '/' {
					i++
					start = i
					for i < len(input) && rune(input[i]) != '\n' {
						i++
					}
					tokens = append(tokens, CommentToken{Value: input[start:i], Line: line, Position: position})
					continue
				} else if r == '/' && nextR == '*' {
					i++
					start = i
					for i+1 < len(input) && (rune(input[i]) != '*' || rune(input[i+1]) != '/') {
						if rune(input[i]) == '\n' {
							line++
							position = 1
						} else {
							position++
						}
						i++
					}
					if i+1 >= len(input) {
						return nil, fmt.Errorf("unterminated multiline comment at line %d, position %d", line, position)
					}
					i += 2 // Move past "*/"
					tokens = append(tokens, CommentToken{Value: input[start : i-2], Line: line, Position: position})
					position += 2
					continue
				}
			}
			tokens = append(tokens, OperatorToken{Value: input[start:i], Line: line, Position: position})
			position += i - start
		} else {
			return nil, fmt.Errorf("unexpected character at line %d, position %d: %c", line, position, r)
		}
	}
	return tokens, nil
}
