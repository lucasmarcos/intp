package intp

import (
	"errors"
)

func lex(exp string) ([]string, error) {
	tokens := make([]string, 0)

	for i := 0; i < len(exp); i++ {
		if exp[i] == ' ' {
		} else if exp[i] == '(' {
			tokens = append(tokens, "(")
		} else if exp[i] == ')' {
			tokens = append(tokens, ")")
		} else if exp[i] == '+' {
			tokens = append(tokens, "+")
		} else if exp[i] == '-' {
			tokens = append(tokens, "-")
		} else if exp[i] == '*' {
			tokens = append(tokens, "*")
		} else if exp[i] == '/' {
			tokens = append(tokens, "/")
		} else {
			if exp[i] >= '0' && exp[i] <= '9' {
				tokens = append(tokens, string(exp[i]))
			} else {
				return nil, errors.New("token invalido")
			}
		}
	}

	return tokens, nil
}

func parse(tokens []string) (string, error) {
	return "", nil
}

func run(exp string) (int, error) {
	return 0, nil
}

func Executar(expressao string) (int, bool) {
	if len(expressao) == 0 {
		return 0, true
	}

	tokens, err := lex(expressao)

	if err != nil {
		return 0, true
	}

	ast, err := parse(tokens)

	if err != nil {
		return 0, true
	}

	result, err := run(ast)

	if err != nil {
		return 0, true
	}

	return result, false
}
