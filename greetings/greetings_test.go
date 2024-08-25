package greetings

import (
	"regexp"
	"testing"
)

/*
	testHelloName caalls greetings.Hello with a name, cheking
	for  a valid return value
*/

func TestHelloName(t *testing.T) {
	name := "Rerin"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Rerin")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Rerin") = %q, %v, want match for %#q, nill`, msg, err, want)
	}
}

/*
	TestHelloEmpty caalls greetings.Hello with an empty string
	checking for an error 
*/

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
