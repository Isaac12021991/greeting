package greeting

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {

	name := "isac"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("isaac")

	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Heloo("isaac) = %q, %v, quiere con para %#q, nill`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")

	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, quiere`, msg, err)
	}
}
