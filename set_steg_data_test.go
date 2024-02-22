package main

import (
	"fmt"
	"testing"
)

func TestSteg(t *testing.T) {

	// The tests to run
	var tests = []struct {
		name           string
		containerFile  string
		hideFile       string
		expectedResult bool
	}{
		{"NoParameterData", "", "", false},
		{"ContainerFileDoesNotExist", "containerFakeFile", "", false},
		{"HideFileDoesNotExist", "", "hideFakeFile", false},
		{"BothFilesDoNotExist", "containerFakeFile", "hideFakeFile", false},
		{"ContainerFileEmpty", "EmptyFile", "100ByteHideFile", false},
		{"HideFileEmpty", "1000ByteContainerFile", "EmptyFile", false},
		{"HideFileTooBig", "1000ByteContainerFile", "1000ByteHideFile", false},
		{"HideFileWontFit", "1000ByteContainerFile", "10ByteHideFile", false},
		{"SmallHideFileFits", "10000ByteContainerFile", "10ByteHideFile", true},
		{"LargeHideFileFits", "10000ByteContainerFile", "100ByteHideFile", true},
	}

	// Set up test data
	CreateEmptyFile("EmptyFile")
	CreateFile("10ByteHideFile", 10)
	CreateFile("100ByteHideFile", 100)
	CreateFile("1000ByteHideFile", 1000)
	CreateFile("1000ByteContainerFile", 1000)
	CreateFile("10000ByteContainerFile", 10000)

	// Write name of function being tested to test results file
	LogResult("Steg")

	// Run the tests
	for _, currentTest := range tests {
		testname := fmt.Sprintf("%s", currentTest.name)
		t.Run(testname, func(t *testing.T) {

			result := Steg(currentTest.containerFile, currentTest.hideFile)

			if result != currentTest.expectedResult {
				LogResult(currentTest.name + " - " + fmt.Sprintf("Input: %s,%s Got: %t Expected: %t", currentTest.containerFile, currentTest.hideFile, result, currentTest.expectedResult) + " - FAIL")
			} else {
				LogResult(currentTest.name + " - " + fmt.Sprintf("Input: %s,%s Got: %t Expected: %t", currentTest.containerFile, currentTest.hideFile, result, currentTest.expectedResult) + " - PASS")
			}
		})
	}

	// Clean up test data
	DeleteFile("EmptyFile")
	DeleteFile("10ByteHideFile")
	DeleteFile("100ByteHideFile")
	DeleteFile("1000ByteHideFile")
	DeleteFile("1000ByteContainerFile")
	DeleteFile("10000ByteContainerFile")
}
