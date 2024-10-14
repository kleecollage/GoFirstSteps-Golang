package greetings_test

import (
	"regexp"
	"testing"

	"github.com/kleecollage/GoFirstSteps-Golang/greetings"
)

func TestHelloName(t *testing.T) {
	name := "Michael"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := greetings.Hello("Michael")

	if !want.MatchString(msg) || err != nil {
		t.Fail()
		t.Fatalf(`Hello("Michael") = %q, %v, quiere coincidencia para %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := greetings.Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, quiere "", error`, msg, err)
	}
}
