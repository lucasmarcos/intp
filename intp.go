package intp

import (
	"errors"
	"strconv"
)

type Folha struct {
	Op    string
	Valor int
	Esq   *Folha
	Dir   *Folha
}

func lex(exp string) ([]string, error) {
	tokens := make([]string, 0)
	num := ""

	for i := 0; i < len(exp); i++ {
		ch := exp[i]

		if ch == ' ' {
			continue
		}

		if ch == '(' || ch == ')' || ch == '+' || ch == '-' || ch == '*' || ch == '/' {
			if num != "" {
				tokens = append(tokens, num)
				num = ""
			}
			tokens = append(tokens, string(ch))
		} else if ch >= '0' && ch <= '9' {
			num += string(ch)
		} else {
			return nil, errors.New("token invalido")
		}
	}

	if num != "" {
		tokens = append(tokens, num)
	}

	return tokens, nil
}

func parse(tokens []string) (*Folha, error) {
	pos := 0

	proximo := func() string {
		if pos < len(tokens) {
			token := tokens[pos]
			pos++
			return token
		}
		return ""
	}

	return parseExp(tokens, &pos, proximo)
}

func parseExp(tokens []string, pos *int, proximo func() string) (*Folha, error) {
	esq, err := parseTermo(tokens, pos, proximo)

	if err != nil {
		return nil, err
	}

	for {
		token := proximo()
		if token == "+" || token == "-" {
			dir, err := parseTermo(tokens, pos, proximo)

			if err != nil {
				return nil, err
			}

			esq = &Folha{
				Op:  token,
				Esq: esq,
				Dir: dir,
			}
		} else {
			(*pos)--
			break
		}
	}

	return esq, nil
}

func parseTermo(tokens []string, pos *int, proximo func() string) (*Folha, error) {
	esq, err := parseFator(tokens, pos, proximo)

	if err != nil {
		return nil, err
	}

	for {
		token := proximo()

		if token == "*" || token == "/" {
			dir, err := parseFator(tokens, pos, proximo)

			if err != nil {
				return nil, err
			}

			esq = &Folha{
				Op:  token,
				Esq: esq,
				Dir: dir,
			}
		} else {
			(*pos)--
			break
		}
	}

	return esq, nil
}

func parseFator(tokens []string, pos *int, proximo func() string) (*Folha, error) {
	token := proximo()

	var esq *Folha

	if token >= "0" && token <= "9" {
		esq, _ = parseNumber(token)
	} else if token == "(" {
		esq, _ = parseExp(tokens, pos, proximo)

		if proximo() != ")" {
			return nil, errors.New("parenteses nao fechados")
		}
	} else {
		return nil, errors.New("token invalido " + token)
	}

	return esq, nil
}

func parseNumber(token string) (*Folha, error) {
	val, err := strconv.Atoi(token)

	if err != nil {
		return nil, errors.New("numero invalido " + token)
	}

	return &Folha{Valor: val}, nil
}

func run(ast *Folha) (int, error) {
	if ast.Op == "" {
		return ast.Valor, nil
	}

	esq, err := run(ast.Esq)

	if err != nil {
		return 0, err
	}

	dir, err := run(ast.Dir)

	if err != nil {
		return 0, err
	}

	switch ast.Op {
	case "+":
		return esq + dir, nil
	case "-":
		return esq - dir, nil
	case "*":
		return esq * dir, nil
	case "/":
		if dir == 0 {
			return 0, errors.New("divisao por zero")
		}
		return esq / dir, nil
	default:
		return 0, errors.New("operador invalido " + ast.Op)
	}
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
