package main

import (
	"fmt"
	"testing"
)

func TestFileExists(t *testing.T) {

	// The tests to run
	var tests = []struct {
		name           string
		input          string
		expectedResult bool
	}{
		{"NoParameterData", "", false},
		{"FileDoesNotExist", "fakeFile", false},
		{"FileExists", "TestFile", true},
	}

	// Set up test data
	CreateFile("TestFile", 100)

	// Write name of function being tested to test results file
	LogResult("FileExists")

	// Run the tests
	for _, currentTest := range tests {
		testname := fmt.Sprintf("%s", currentTest.name)
		t.Run(testname, func(t *testing.T) {

			result := FileExists(currentTest.input)

			if result != currentTest.expectedResult {
				LogResult(currentTest.name + " - " + fmt.Sprintf("Input: %s Got: %t Expected: %t", currentTest.input, result, currentTest.expectedResult) + " - FAIL")
			} else {
				LogResult(currentTest.name + " - " + fmt.Sprintf("Input: %s Got: %t Expected: %t", currentTest.input, result, currentTest.expectedResult) + " - PASS")
			}
		})
	}

	// Clean up test data
	DeleteFile("TestFile")
}

func TestGetFileSize(t *testing.T) {

	// The tests to run
	var tests = []struct {
		name           string
		input          string
		expectedResult int64
	}{
		{"NoParameterData", "", -1},
		{"FileDoesNotExist", "fakeFile", -1},
		{"1ByteFile", "1ByteFile", 1},
		{"100ByteFile", "100ByteFile", 100},
	}

	// Set up test data
	CreateFile("1ByteFile", 1)
	CreateFile("100ByteFile", 100)

	// Write name of function being tested to test results file
	LogResult("GetFileSize")

	// Run the tests
	for _, currentTest := range tests {
		testname := fmt.Sprintf("%s", currentTest.name)
		t.Run(testname, func(t *testing.T) {

			result := GetFileSize(currentTest.input)

			if result != currentTest.expectedResult {
				LogResult(currentTest.name + " - " + fmt.Sprintf("Input: %s Got: %d Expected: %d", currentTest.input, result, currentTest.expectedResult) + " - FAIL")
			} else {
				LogResult(currentTest.name + " - " + fmt.Sprintf("Input: %s Got: %d Expected: %d", currentTest.input, result, currentTest.expectedResult) + " - PASS")
			}
		})
	}

	// Clean up test data
	DeleteFile("1ByteFile")
	DeleteFile("100ByteFile")
}

func TestDeleteFile(t *testing.T) {

	// The tests to run
	var tests = []struct {
		name           string
		input          string
		expectedResult bool
	}{
		{"NoParameterData", "", true},
		{"FileDoesNotExist", "fakeFile", true},
		{"FileExists", "TestFile", false},
	}

	// Set up test data
	CreateFile("TestFile", 10)

	// Write name of function being tested to test results file
	LogResult("DeleteFile")

	// Run the tests
	for _, currentTest := range tests {
		testname := fmt.Sprintf("%s", currentTest.name)
		t.Run(testname, func(t *testing.T) {

			result := DeleteFile(currentTest.input)

			if result != currentTest.expectedResult {
				LogResult(currentTest.name + " - " + fmt.Sprintf("Input: %s Got: %t Expected: %t", currentTest.input, result, currentTest.expectedResult) + " - FAIL")
			} else {
				LogResult(currentTest.name + " - " + fmt.Sprintf("Input: %s Got: %t Expected: %t", currentTest.input, result, currentTest.expectedResult) + " - PASS")
			}
		})
	}
}

func TestWipeFile(t *testing.T) {

	// The tests to run
	var tests = []struct {
		name           string
		input          string
		expectedResult bool
	}{
		{"NoParameterData", "", false},
		{"FileDoesNotExist", "fakeFile", false},
		{"FileExists10", "TestFile10", true},
		{"FileExists100", "TestFile100", true},
		{"FileExists1000", "TestFile1000", true},
		{"FileExists1234", "TestFile1234", true},
		{"FileExists5678", "TestFile5678", true},
		{"FileExists98765", "TestFile98765", true},
	}

	// Set up test data
	CreateFile("TestFile10", 10)
	CreateFile("TestFile100", 100)
	CreateFile("TestFile1000", 1000)
	CreateFile("TestFile1234", 1234)
	CreateFile("TestFile5678", 5678)
	CreateFile("TestFile98765", 98765)

	// Write name of function being tested to test results file
	LogResult("WipeFile")

	// Run the tests
	for _, currentTest := range tests {
		testname := fmt.Sprintf("%s", currentTest.name)
		t.Run(testname, func(t *testing.T) {

			result := WipeFile(currentTest.input)

			if result != currentTest.expectedResult {
				LogResult(currentTest.name + " - " + fmt.Sprintf("Input: %s Got: %t Expected: %t", currentTest.input, result, currentTest.expectedResult) + " - FAIL")
			} else {
				LogResult(currentTest.name + " - " + fmt.Sprintf("Input: %s Got: %t Expected: %t", currentTest.input, result, currentTest.expectedResult) + " - PASS")
			}
		})
	}
}
