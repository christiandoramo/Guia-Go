package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) { // nome deve começar com Test
	name := "Christian"
	want := regexp.MustCompile((`\b` + name + `\b`))
	msg, err := Hello("Christian")
	if !want.MatchString(msg) || err != nil { // se o nome não faz parte então há erro
		t.Fatalf(`Hello("Christian") = %q, %v, esperava dar match para %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) { // *testing.T é o tipo teste recebido ao rodar go test
	msg, err := Hello("")
	if msg != "" || err == nil { // caso o retorno seja ok mesmo com "" então da erro
		t.Fatalf(`Hello("") = %q, %v, esperava "", erro`, msg, err)
	}

}
