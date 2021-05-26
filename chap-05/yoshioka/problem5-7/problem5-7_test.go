package main

import "testing"

func TestLCS1(t *testing.T) {
	s1 := "abcde"
	s2 := "acbef"

	r := LCS(s1, s2)
	if r != 3 {
		t.Error("error in 1st test")
	}

}

func TestLCS2(t *testing.T) {
	s1 := "pirikapirirara"
	s2 := "poporinapeperuto"

	r := LCS(s1, s2)
	if r != 6 {
		t.Error("error in 2nd test")
	}

}
