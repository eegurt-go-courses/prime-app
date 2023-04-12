package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_isPrime(t *testing.T) {
	primeTests := []struct {
		name     string
		testNum  int
		expected bool
		msg      string
	}{
		{"prime", 7, true, "7 is a prime number!"},
		{"not prime", 8, false, "8 is not a prime number because it is divisible by 2!"},
		{"zero", 0, false, "0 is not prime, by definition!"},
		{"one", 1, false, "1 is not prime, by definition!"},
		{"negative number", -11, false, "Negative numbers are not prime, by definition!"},
	}

	for _, e := range primeTests {
		result, msg := isPrime(e.testNum)
		if e.expected && !result {
			t.Errorf("%s: expected true but got false", e.name)
		}

		if !e.expected && result {
			t.Errorf("%s: expected false but got true", e.name)
		}

		if e.msg != msg {
			t.Errorf("%s: expected %s but got %s", e.name, e.msg, msg)
		}
	}
}

func getStdout(f func()) string {
	old := os.Stdout

	r, w, _ := os.Pipe()

	os.Stdout = w
	f()
	w.Close()

	os.Stdout = old

	bytes, _ := io.ReadAll(r)
	r.Close()
	out := string(bytes)

	return out
}

func Test_prompt(t *testing.T) {
	out := getStdout(prompt)

	expected := "-> "

	ok := strings.EqualFold(out, expected)
	if !ok {
		t.Errorf("Expected \"%s\", got \"%s\"", expected, out)
	}
}

func Test_intro(t *testing.T) {
	out := getStdout(intro)

	var sb strings.Builder
	sb.WriteString("Is it Prime?\n")
	sb.WriteString("------------\n")
	sb.WriteString("Enter a whole number, and we'll tell you if it is a prime number or not. Enter q to quit.\n")
	sb.WriteString("-> ")
	expected := sb.String()

	ok := strings.EqualFold(out, expected)
	if !ok {
		t.Errorf("Expected \"%s\", got \"%s\"", expected, out)
	}
}
