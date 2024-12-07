package intp

func lex(exp string) ([]string, error) {
	return nil, nil
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
