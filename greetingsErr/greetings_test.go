package greetingsErr

import (
    "testing"
    "regexp"
)

// TestHelloName calls greetings.Hello with a name, checking for a valid return value.
func TestHelloName (t *testing.T) {
    name := "Nipun"
    want := regexp.MustCompile(`\b` + name + `\b`)
    msg, err := Hello("Nipun")
    if !want.MatchString(msg) || err != nil {
        t.Errorf(`Hello("Nipun") = %q %v, want match for %#q, nil`, msg, err, want)
    }
}

// TestHelloEmpty calls greetings.Hello with an empty string, checking for an error.
func TestHelloEmpty(t *testing.T) {
    msg, err := Hello("")
    if msg != "" || err == nil {
        t.Errorf(`Hello("") = %q, %v, want "", error`, msg, err)
    }
}

// The go test command executes test functions (whose names begin with Test) in test files (whose names end with 
// _test.go). You can add the -v flag to get verbose output that lists all of the tests and their results.