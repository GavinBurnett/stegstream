package main

import (
	"fmt"
	"testing"
)

func TestReadConfigFile(t *testing.T) {

	// Set up test data
	var configData Config = Config{DEFAULT_PORT, DEFAULT_STREAM_ONLY, DEFAULT_HIDE_ONLY, DEFAULT_WIPE_AUDIO, DEFAULT_WIPE_HIDDEN}

	// The tests to run
	var tests = []struct {
		name           string
		input          string
		expectedResult Config
	}{
		{"NoParameterData", "", configData},
		{"FileDoesNotExist", "fakeFile", configData},
		{"FileExists", CONFIG_FILE, configData},
	}

	// Write name of function being tested to test results file
	LogResult("ReadConfigFile")

	// Run the tests
	for _, currentTest := range tests {
		testname := fmt.Sprintf("%s", currentTest.name)
		t.Run(testname, func(t *testing.T) {

			result := ReadConfigFile(currentTest.input)

			if result != currentTest.expectedResult {
				LogResult(currentTest.name + " - " + fmt.Sprintf("Input: %s Got: %v Expected: %v", currentTest.input, result, currentTest.expectedResult) + " - FAIL")
			} else {
				LogResult(currentTest.name + " - " + fmt.Sprintf("Input: %s Got: %v Expected: %v", currentTest.input, result, currentTest.expectedResult) + " - PASS")
			}
		})
	}
}

func TestCheckConfigFile(t *testing.T) {

	// Set up test data
	var defaultData Config = Config{DEFAULT_PORT, DEFAULT_STREAM_ONLY, DEFAULT_HIDE_ONLY, DEFAULT_WIPE_AUDIO, DEFAULT_WIPE_HIDDEN}
	var portChanged Config = Config{8181, DEFAULT_STREAM_ONLY, DEFAULT_HIDE_ONLY, DEFAULT_WIPE_AUDIO, DEFAULT_WIPE_HIDDEN}
	var invalidPort Config = Config{-1, DEFAULT_STREAM_ONLY, DEFAULT_HIDE_ONLY, DEFAULT_WIPE_AUDIO, DEFAULT_WIPE_HIDDEN}
	var streamOnlyOn Config = Config{DEFAULT_PORT, true, DEFAULT_HIDE_ONLY, DEFAULT_WIPE_AUDIO, DEFAULT_WIPE_HIDDEN}
	var hideOnlyOn Config = Config{DEFAULT_PORT, DEFAULT_STREAM_ONLY, true, DEFAULT_WIPE_AUDIO, DEFAULT_WIPE_HIDDEN}
	var streamOnlyHideOnlyBothOn Config = Config{DEFAULT_PORT, true, true, DEFAULT_WIPE_AUDIO, DEFAULT_WIPE_HIDDEN}
	var wipeAudioOn Config = Config{DEFAULT_PORT, DEFAULT_STREAM_ONLY, DEFAULT_HIDE_ONLY, true, DEFAULT_WIPE_HIDDEN}
	var wipeAudioHideOnlyOn Config = Config{DEFAULT_PORT, DEFAULT_STREAM_ONLY, true, true, DEFAULT_WIPE_HIDDEN}
	var wipeHiddenOn Config = Config{DEFAULT_PORT, DEFAULT_STREAM_ONLY, DEFAULT_HIDE_ONLY, DEFAULT_WIPE_AUDIO, true}
	var wipeHiddenStreamOnlyBothOn Config = Config{DEFAULT_PORT, true, DEFAULT_HIDE_ONLY, DEFAULT_WIPE_AUDIO, true}

	// The tests to run
	var tests = []struct {
		name            string
		input           Config
		expectedConfig  Config
		expectedBoolean bool
	}{
		{"DefaultData", defaultData, defaultData, true},
		{"PortChanged", portChanged, portChanged, true},
		{"InvalidPort", invalidPort, defaultData, true},
		{"streamOnlyOn", streamOnlyOn, streamOnlyOn, true},
		{"hideOnlyOn", hideOnlyOn, hideOnlyOn, true},
		{"streamOnlyHideOnlyBothOn", streamOnlyHideOnlyBothOn, streamOnlyHideOnlyBothOn, false},
		{"wipeAudioOn", wipeAudioOn, wipeAudioOn, true},
		{"wipeAudioHideOnlyOn", wipeAudioHideOnlyOn, wipeAudioHideOnlyOn, false},
		{"wipeHiddenOn", wipeHiddenOn, wipeHiddenOn, true},
		{"wipeHiddenStreamOnlyBothOn", wipeHiddenStreamOnlyBothOn, wipeHiddenStreamOnlyBothOn, false},
	}

	// Write name of function being tested to test results file
	LogResult("CheckConfigFile")

	// Run the tests
	for _, currentTest := range tests {
		testname := fmt.Sprintf("%s", currentTest.name)
		t.Run(testname, func(t *testing.T) {

			resultConfig, resultBoolean := CheckConfigFile(currentTest.input)

			if resultConfig != currentTest.expectedConfig && resultBoolean != currentTest.expectedBoolean {
				LogResult(currentTest.name + " - " + fmt.Sprintf("Input: %v Got: %v %t Expected: %v %t", currentTest.input, resultConfig, resultBoolean, currentTest.expectedConfig, currentTest.expectedBoolean) + " - FAIL")
			} else {
				LogResult(currentTest.name + " - " + fmt.Sprintf("Input: %v Got: %v %t Expected: %v %t", currentTest.input, resultConfig, resultBoolean, currentTest.expectedConfig, currentTest.expectedBoolean) + " - PASS")
			}
		})
	}
}

func TestParseStringToInt(t *testing.T) {

	// The tests to run
	var tests = []struct {
		name           string
		input          string
		expectedResult int
	}{
		{"NoParameterData", "", -1},
		{"InvalidData1", "invalidData", -1},
		{"InvalidData2", "34e", -1},
		{"InvalidData3", "e56", -1},
		{"ZeroNumber", "0", -1},
		{"MinusNumber", "-4", -1},
		{"10", "10", 10},
		{"8080", "8080", 8080},
	}

	// Write name of function being tested to test results file
	LogResult("ParseStringToInt")

	// Run the tests
	for _, currentTest := range tests {
		testname := fmt.Sprintf("%s", currentTest.name)
		t.Run(testname, func(t *testing.T) {

			result := ParseStringToInt(currentTest.input)

			if result != currentTest.expectedResult {
				LogResult(currentTest.name + " - " + fmt.Sprintf("Input: %s Got: %d Expected: %d", currentTest.input, result, currentTest.expectedResult) + " - FAIL")
			} else {
				LogResult(currentTest.name + " - " + fmt.Sprintf("Input: %s Got: %d Expected: %d", currentTest.input, result, currentTest.expectedResult) + " - PASS")
			}
		})
	}
}

func TestParseStringToBool(t *testing.T) {

	// The tests to run
	var tests = []struct {
		name           string
		input          string
		expectedResult bool
	}{
		{"NoParameterData", "", false},
		{"InvalidData1", "invalidData", false},
		{"InvalidData2", "34e", false},
		{"MinusNumber", "-4", false},
		{"ZeroNumber", "0", false},
		{"OneNumber", "1", true},
		{"LowerCaseTrue", "true", true},
		{"UpperCaseTrue", "TRUE", true},
		{"LowerCaseFalse", "false", false},
		{"UpperCaseFalse", "FALSE", false},
	}

	// Write name of function being tested to test results file
	LogResult("ParseStringToBool")

	// Run the tests
	for _, currentTest := range tests {
		testname := fmt.Sprintf("%s", currentTest.name)
		t.Run(testname, func(t *testing.T) {

			result := ParseStringToBool(currentTest.input)

			if result != currentTest.expectedResult {
				LogResult(currentTest.name + " - " + fmt.Sprintf("Input: %s Got: %t Expected: %t", currentTest.input, result, currentTest.expectedResult) + " - FAIL")
			} else {
				LogResult(currentTest.name + " - " + fmt.Sprintf("Input: %s Got: %t Expected: %t", currentTest.input, result, currentTest.expectedResult) + " - PASS")
			}
		})
	}
}
