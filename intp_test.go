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
}

func TestPrecedenciaCom(t *testing.T) {
}

func TestPrecedenciaAninhado(t *testing.T) {
}

func TestSintaxeEspaco(t *testing.T) {
}

func TestSintaxeIncompleto(t *testing.T) {
}

func TestSintaxeOperadores(t *testing.T) {
}

func TestSintaxeDesbalanco(t *testing.T) {
}

func TestSintaxeMisto(t *testing.T) {
}

func TestSemanticaDivZero(t *testing.T) {
}

func TestSemanticaVazio(t *testing.T) {
}
