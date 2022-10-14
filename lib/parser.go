package lib

import (
	"fmt"
	"strconv"
	"strings"
)

// Parse returns list of tokens and syntax errors if found
func Parse(content string) ([]Token, error) {
	tokens := make([]Token, 0)

	chars := strings.Split(content, "")

	currBlock := ""
	currType := NONE

	// start := 0 (used for debugging/error messages later)
	line := 1

	for _, char := range chars {
		sct := singleCharToken(char)

		if char == "\n" {
			// new line
			if currType == STRINGVALUE {
				// string not ended
				return tokens, fmt.Errorf("Reached end of line while parsing string on line %d", line)
			}

			if currType != NONE {
				tokens = append(tokens, NewToken(currType, currBlock))
				currType = NONE
				currBlock = ""

				tokens = append(tokens, NewToken(NEWLINE, ""))
			}

			line++
		} else if sct.Type != NONE && currType != STRINGVALUE {
			// the char is a single char token
			if currType == UNKNOWNVALUE {
				tokens = unknown(currBlock, tokens)
			} else if currType != NONE {
				tokens = append(tokens, NewToken(currType, currBlock))
			}
			tokens = append(tokens, sct)
			currBlock = ""
			currType = NONE
		} else if currType == STRINGVALUE {
			// in string
			if char == "\"" {
				// end string
				tokens = append(tokens, NewToken(STRINGVALUE, currBlock))
				currBlock = ""
				currType = NONE
			} else {
				// part of string
				currBlock += char
			}
		} else if currType == NUMVALUE {
			if !isNum(char) {
				// not digit in num
				return tokens, fmt.Errorf("Unexpected character found while parsing number on line %d", line)
			}
			currBlock += char
		} else {
			// not in string
			if char == "\"" {
				// beginning of string
				currType = STRINGVALUE
			} else if char == " " {
				// space
				if currType != NONE {
					// could be bool or identifier
					tokens = unknown(currBlock, tokens)
					currBlock = ""
					currType = NONE
				}
			} else if currType == NONE {
				if isNum(char) {
					currType = NUMVALUE
				} else if isAlpha(char) {
					currType = UNKNOWNVALUE
				} else {
					return tokens, fmt.Errorf("Unexpected character %s on line %d", char, line)
				}
				currBlock += char
			} else if currType == UNKNOWNVALUE {
				if isNum(char) || isAlpha(char) || char == "_" {
					currBlock += char
				}
			} else {
				return tokens, fmt.Errorf("Unexpected character %s on line %d", char, line)
			}
		}
	}

	return tokens, nil
}

func isNum(str string) bool {
	_, err := strconv.ParseFloat(str, 64)
	return err == nil
}

func isAlpha(str string) bool {
	// could be optimized with binary search
	return strings.Contains("qwertyuiopasdfghjklzxcvbnm", str)
}

func unknown(currBlock string, tokens []Token) []Token {
	_, err := strconv.ParseBool(currBlock)
	typeTok := typeToken(currBlock)
	if err != nil {
		// not bool
		if typeTok.Type != NONE {
			return append(tokens, typeTok)
		}
		return append(tokens, NewToken(IDENTIFIER, currBlock))
	}
	return append(tokens, NewToken(BOOLEANVALUE, currBlock))
}

func typeToken(str string) Token {
	switch str {
	case "int":
		return NewToken(INTTYPE, "")
	case "string":
		return NewToken(STRINGTYPE, "")
	case "float":
		return NewToken(FLOATTYPE, "")
	case "bool":
		return NewToken(BOOLEANTYPE, "")
	}
	return NewToken(NONE, "")
}

func singleCharToken(str string) Token {
	switch str {
	case "(":
		return NewToken(LEFTPAREN, "")
	case ")":
		return NewToken(RIGHTPAREN, "")
	case "[":
		return NewToken(LEFTBRACKET, "")
	case "]":
		return NewToken(RIGHTBRACKET, "")
	case "{":
		return NewToken(LEFTCURLY, "")
	case "}":
		return NewToken(RIGHTCURLY, "")
	case ",":
		return NewToken(COMMA, "")
	case "+":
		return NewToken(ADD, "")
	case "-":
		return NewToken(SUBTRACT, "")
	case "*":
		return NewToken(MULTIPLY, "")
	case "/":
		return NewToken(DIVIDE, "")
	case "%":
		return NewToken(MODULUS, "")
	case "=":
		return NewToken(ASSIGN, "")
	case "<":
		return NewToken(LESS, "")
	case ">":
		return NewToken(GREATER, "")
	case "!":
		return NewToken(NOT, "")
	case "&":
		return NewToken(AND, "")
	case "|":
		return NewToken(OR, "")
	}
	return NewToken(NONE, "")
}
