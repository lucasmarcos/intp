package intp_test

import "testing"
import "intp"

func TestBasicoSoma(t *testing.T) {
	input := "1+2"
	expected := 3

	output, err := intp.Executar(input)

	if err != false || output != expected {
		t.Errorf("%s, %d, %d", input, expected, output)
	}
}

func TestBasicoSub(t *testing.T) {
	input := "5-3"
	expected := 2

	output, err := intp.Executar(input)

	if err != false || output != expected {
		t.Errorf("%s, %d, %d", input, expected, output)
	}
}

func TestBasicoMult(t *testing.T) {
	input := "4*3"
	expected := 12

	output, err := intp.Executar(input)

	if err != false || output != expected {
		t.Errorf("%s, %d, %d", input, expected, output)
	}
}

func TestBasicoDiv(t *testing.T) {
	input := "8/2"
	expected := 4

	output, err := intp.Executar(input)

	if err != false || output != expected {
		t.Errorf("%s, %d, %d", input, expected, output)
	}
}

func TestNegativoSoma(t *testing.T) {
	input := "-1+2"
	expected := 1

	output, err := intp.Executar(input)

	if err != false || output != expected {
		t.Errorf("%s, %d, %d", input, expected, output)
	}
}

func TestNegativoMult(t *testing.T) {
	input := "-3*4"
	expected := -12

	output, err := intp.Executar(input)

	if err != false || output != expected {
		t.Errorf("%s, %d, %d", input, expected, output)
	}
}

func TestNegativoDiv(t *testing.T) {
	input := "6/-2"
	expected := -3

	output, err := intp.Executar(input)

	if err != false || output != expected {
		t.Errorf("%s, %d, %d", input, expected, output)
	}
}

func TestPrecedenciaSem(t *testing.T) {
	input := "2+3*4"
	expected := 14

	output, err := intp.Executar(input)

	if err != false || output != expected {
		t.Errorf("%s, %d, %d", input, expected, output)
	}
}

func TestPrecedenciaCom(t *testing.T) {
	input := "(2+3)*4"
	expected := 20

	output, err := intp.Executar(input)

	if err != false || output != expected {
		t.Errorf("%s, %d, %d", input, expected, output)
	}
}

func TestPrecedenciaAninhado(t *testing.T) {
	input := "2*(3+(4-1))"
	expected := 12

	output, err := intp.Executar(input)

	if err != false || output != expected {
		t.Errorf("%s, %d, %d", input, expected, output)
	}
}

func TestSintaxeEspaco(t *testing.T) {
	input := "1 + 2"
	expected := 3

	output, err := intp.Executar(input)

	if err != false || output != expected {
		t.Errorf("%s, %d, %d", input, expected, output)
	}
}

func TestSintaxeIncompleto(t *testing.T) {
	input := "2 +"

	_, err := intp.Executar(input)

	if err != true {
		t.Errorf("%s, %t", input, err)
	}
}

func TestSintaxeOperadores1(t *testing.T) {
	input := "2**2"

	_, err := intp.Executar(input)

	if err != true {
		t.Errorf("%s, %t", input, err)
	}
}

func TestSintaxeOperadores2(t *testing.T) {
	input := "2*/2"

	_, err := intp.Executar(input)

	if err != true {
		t.Errorf("%s, %t", input, err)
	}
}

func TestSintaxeDesbalanco(t *testing.T) {
	input := "(2+3"

	_, err := intp.Executar(input)

	if err != true {
		t.Errorf("%s, %t", input, err)
	}
}

func TestSintaxeMisto(t *testing.T) {
	input := "((1 + 2) * (3 + 4)) - 5"
	expected := 16

	output, err := intp.Executar(input)

	if err != false || output != expected {
		t.Errorf("%s, %d, %d", input, expected, output)
	}
}

func TestSemanticaDivZero(t *testing.T) {
	input := "5/0"

	_, err := intp.Executar(input)

	if err != true {
		t.Errorf("%s, %t", input, err)
	}
}

func TestSemanticaVazio(t *testing.T) {
	input := ""

	_, err := intp.Executar(input)

	if err != true {
		t.Errorf("%s, %t", input, err)
	}
}
