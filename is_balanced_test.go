package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsBalance(t *testing.T) {
	tableTest := []struct {
		VALUE    string
		EXPECTED string
		GOT      string
	}{
		{VALUE: "{[()]}", EXPECTED: "YES", GOT: isBalanced("{[()]}")},
		{VALUE: "{[(])}", EXPECTED: "NO", GOT: isBalanced("{[(])}")},
		{VALUE: "{{[[(())]]}}", EXPECTED: "YES", GOT: isBalanced("{{[[(())]]}}")},
		{VALUE: "()", EXPECTED: "YES", GOT: isBalanced("()")},
		{VALUE: "{(([])[])[]]}", EXPECTED: "NO", GOT: isBalanced("{(([])[])[]]}")},
		{VALUE: "({[]})", EXPECTED: "YES", GOT: isBalanced("({[]})")},
		{VALUE: "({[})", EXPECTED: "NO", GOT: isBalanced("({[})")},
		{VALUE: "", EXPECTED: "YES", GOT: isBalanced("")},  // empty string
		{VALUE: "(", EXPECTED: "NO", GOT: isBalanced("(")}, // single opening
		{VALUE: ")", EXPECTED: "NO", GOT: isBalanced(")")}, // single closing
		{VALUE: "([]{})", EXPECTED: "YES", GOT: isBalanced("([]{})")},
	}

	for _, test := range tableTest {
		assert.Equal(t, test.EXPECTED, test.GOT)
	}

}
