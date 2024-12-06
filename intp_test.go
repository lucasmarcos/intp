package intp_test

import "testing"
import "intp"

func TestesBasicos(t *testing.T) {
	input := ""

	output := intp.Executar(input)

	if output != false {
		t.Errorf("%q %v %v", input, output, false)
	}
}
