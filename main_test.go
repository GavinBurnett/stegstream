package main

import (
	"fmt"
	"testing"
	"time"
)

func TestAutoShutdown(t *testing.T) {

	// The tests to run
	var tests = []struct {
		name           string
		input          time.Time
		expectedResult bool
	}{
		{"ZeroTime", time.Time{}, false},
		{"Past", time.Date(2000, 04, 11, 12, 00, 00, 00, time.Local), true},
		{"Future", time.Date(2044, 04, 11, 12, 00, 00, 00, time.Local), false},
		{"CurrentTime", time.Now(), true},
	}

	// Write name of function being tested to test results file
	LogResult("AutoShutdown")

	// Run the tests
	for _, currentTest := range tests {
		testname := fmt.Sprintf("%s", currentTest.name)
		t.Run(testname, func(t *testing.T) {

			result := AutoShutdown(currentTest.input)

			if result != currentTest.expectedResult {
				LogResult(currentTest.name + " - " + fmt.Sprintf("Input: %s Got: %t Expected: %t", currentTest.input, result, currentTest.expectedResult) + " - FAIL")
			} else {
				LogResult(currentTest.name + " - " + fmt.Sprintf("Input: %s Got: %t Expected: %t", currentTest.input, result, currentTest.expectedResult) + " - PASS")
			}
		})
	}
}

func TestIsStringHelpArgument(t *testing.T) {

	// The tests to run
	var tests = []struct {
		name           string
		input          string
		expectedResult bool
	}{
		{"NoParameterData", "", false},
		{"NotAHelpArgument", "fg", false},
		{"?", "?", true},
		{"/?", "/?", true},
		{"-?", "-?", true},
		{"--?", "--?", true},
		{"h", "h", true},
		{"/h", "/h", true},
		{"-h", "-h", true},
		{"--h", "--h", true},
		{"help", "help", true},
		{"/help", "/help", true},
		{"-help", "-help", true},
		{"--help", "--help", true},
	}

	// Write name of function being tested to test results file
	LogResult("IsStringHelpArgument")

	// Run the tests
	for _, currentTest := range tests {
		testname := fmt.Sprintf("%s", currentTest.name)
		t.Run(testname, func(t *testing.T) {

			result := IsStringHelpArgument(currentTest.input)

			if result != currentTest.expectedResult {
				LogResult(currentTest.name + " - " + fmt.Sprintf("Input: %s Got: %t Expected: %t", currentTest.input, result, currentTest.expectedResult) + " - FAIL")
			} else {
				LogResult(currentTest.name + " - " + fmt.Sprintf("Input: %s Got: %t Expected: %t", currentTest.input, result, currentTest.expectedResult) + " - PASS")
			}
		})
	}
}
