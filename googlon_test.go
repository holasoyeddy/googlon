package main

import (
	"testing"
)

func TestIsPropistion(t *testing.T) {

	// Arrange
	cases := []struct {
		value    string
		expected bool
	}{
		{value: "iygsex", expected: true},
		{value: "rhfl", expected: false},
		{value: "rsogxd", expected: true},
		{value: "qhmoixw", expected: false},
	}

	for tindex, tcase := range cases {

		// Act
		result := isPreposition(tcase.value)

		// Assert
		assertEqual(tindex+1, t, result, tcase.expected)
		t.Logf("Case %v: PASSED", tindex+1)
	}
}

func TestIsVerb(t *testing.T) {

	// Arrange
	cases := []struct {
		value    string
		expected bool
	}{
		{value: "fgixsr", expected: true},
		{value: "do", expected: false},
		{value: "cejfugn", expected: true},
		{value: "ermjdhsx", expected: false},
	}

	for tindex, tcase := range cases {

		// Act
		result := isVerb(tcase.value)

		// Assert
		assertEqual(tindex+1, t, result, tcase.expected)
		t.Logf("Case %v: PASSED", tindex+1)
	}
}

func TestIsSubjunctiveVerb(t *testing.T) {

	// Arrange
	cases := []struct {
		value    string
		expected bool
	}{
		{value: "ggixsr", expected: true},
		{value: "do", expected: false},
		{value: "ceasdcjfugn", expected: true},
		{value: "ermjdhsx", expected: false},
	}

	for tindex, tcase := range cases {

		// Act
		result := isVerb(tcase.value)

		// Assert
		assertEqual(tindex+1, t, result, tcase.expected)
		t.Logf("Case %v: PASSED", tindex+1)
	}
}

func TestGooglonSort(t *testing.T) {

	// Arrange
	cases := []struct {
		value    []string
		expected []string
	}{
		{
			value:    []string{"ghepqyd", "pdoymnwxei", "emjocsild", "shoce"},
			expected: []string{"shoce", "pdoymnwxei", "emjocsild", "ghepqyd"},
		},
		{
			value:    []string{"pdoymnwxei", "phfer", "pq", "podciy"},
			expected: []string{"podciy", "pq", "phfer", "pdoymnwxei"},
		},
		{
			value:    []string{"pppp", "ppp", "pp", "p"},
			expected: []string{"p", "pp", "ppp", "pppp"},
		},
	}

	for tindex, tcase := range cases {

		// Act
		googlonSort(tcase.value)

		for index := range tcase.value {
			// Assert
			assertEqual(tindex+1, t, tcase.value[index], tcase.expected[index])
		}

		t.Logf("Case %v: PASSED", tindex+1)
	}

}

// Simple wrapper function to assert equality.
func assertEqual(caseno int, t *testing.T, result interface{}, expected interface{}) {

	if result != expected {
		t.Fatalf("Case %v: FAILED - VALUE MISMATCH: Result - %v : Expected - %v", caseno, result, expected)
	}
}
func TestWordToNumber(t *testing.T) {

	cases := []struct {
		value    string
		expected int64
	}{
		{value: "gxjrc", expected: 605637},
		{value: "meofh", expected: 1833046},
		{value: "hej", expected: 5851},
		{value: "dcnql", expected: 2114076},
	}

	for tindex, tcase := range cases {

		result := wordToNumber(tcase.value)
		assertEqual(tindex+1, t, result, tcase.expected)
		t.Logf("Case %v: PASSED", tindex+1)
	}
}

func TestIsPrettyNumber(t *testing.T) {

	// Arrange
	cases := []struct {
		value    string
		expected bool
	}{
		{value: "gxjrc", expected: true},
		{value: "meofh", expected: false},
		{value: "hej", expected: false},
		{value: "dcnql", expected: true},
	}

	for tindex, tcase := range cases {

		// Act
		result := isPrettyNumber(tcase.value)

		// Assert
		assertEqual(tindex+1, t, result, tcase.expected)
		t.Logf("Case %v: PASSED", tindex+1)
	}
}
