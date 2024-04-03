package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"time"
)

const TestResultsFile string = "ServerTestResults.txt"

// LogResult: Logs a given test result to the test results file
func LogResult(_result string) {

	var file *os.File
	var err error

	if len(_result) > 0 {

		// Create and/or open the test results file
		file, err = os.OpenFile(TestResultsFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println(UI_FileOpenError, TestResultsFile, err)
		}

		if err == nil {
			// Write the test result to the file
			_, err = file.WriteString(_result + "\n")
			if err != nil {
				fmt.Println(fmt.Sprintf(UI_FileWriteError, TestResultsFile, err))
			}
		}

		file.Sync()
		file.Close()

	} else {
		fmt.Println(fmt.Sprintf(UI_ParameterInvalid, GetFunctionName()), "_result", _result)
	}
}

// CreateFile: Creates a file with the given file name and file size, filled with random data
func CreateFile(_fileName string, _fileSize int64) bool {

	var fileCreated bool = true
	var randomData []byte
	var bytesWritten int64
	var writeBuffer *bufio.Writer

	if len(_fileName) > 0 && _fileSize > 0 {

		fmt.Print(UI_CreatingFile + _fileName)

		// Create the file
		file, err := os.Create(_fileName)
		if err == nil {

			writeBuffer = bufio.NewWriter(file)
			writeBuffer = bufio.NewWriterSize(writeBuffer, int(_fileSize))

			for bytesWritten = 0; bytesWritten < _fileSize; bytesWritten += _fileSize {

				randomData = GetRandomData(_fileSize)

				if randomData != nil && int64(len(randomData)) == _fileSize {

					bytesWritten, err := writeBuffer.Write(randomData)

					if err != nil || bytesWritten != len(randomData) {
						fileCreated = false
						break
					}

					randomData = nil
				} else {
					fmt.Println(UI_RandomDataError)
				}
			} // end for

			writeBuffer.Flush()
			writeBuffer = nil

			if fileCreated == false {
				fmt.Println(UI_FileCreateError, _fileName, _fileSize, err.Error())

			} else {
				// File created
				fmt.Println(UI_Done)
			}

		} else {
			if DEBUG == true {
				fmt.Println(UI_FileCreateError, _fileName, _fileSize)
			}
			fileCreated = false
		}

		file.Sync()
		file.Close()

	} else {
		fileCreated = false
		fmt.Println(UI_FileCreateError, _fileName, _fileSize, GetFunctionName())
	}

	return fileCreated
}

// CreateEmptyFile: Creates an empty file with the given file name
func CreateEmptyFile(_fileName string) bool {

	var fileCreated bool = true

	if len(_fileName) > 0 {

		fmt.Print(UI_CreatingFile + _fileName)

		// Create the file
		file, err := os.Create(_fileName)
		if err == nil {

			if fileCreated == false {
				fmt.Println(UI_FileCreateError, _fileName, err.Error())

			} else {
				// File created
				fmt.Println(UI_Done)
			}

		} else {
			if DEBUG == true {
				fmt.Println(UI_FileCreateError, _fileName)
			}
			fileCreated = false
		}

		file.Sync()
		file.Close()

	} else {
		fileCreated = false
		fmt.Println(UI_FileCreateError, _fileName, GetFunctionName())
	}

	return fileCreated
}

// GetRandomNumber: Gets a random number between given low and high values
func GetRandomNumber(_low int64, _high int64) int64 {

	var randomNumber int64
	randomNumber = -1

	if _low >= 0 && _low <= math.MaxInt64 && _high >= 0 && _high <= math.MaxInt64 {

		if _low < _high {
			rand.Seed(time.Now().UnixNano())
			randomNumber = rand.Int63n(_high-_low+1) + _low
		}

	} else {
		fmt.Println(fmt.Sprintf(UI_ParameterInvalid, GetFunctionName()), "_low:", _low, "_high:", _high)
	}

	return randomNumber
}

// GetRandomData: Gets a byte array of given size filled with random data
func GetRandomData(_size int64) []byte {

	var randomData []byte
	randomData = nil

	if _size > 0 && _size <= math.MaxInt64 {

		randomData = make([]byte, _size)

		if int64(len(randomData)) != _size {
			fmt.Println(UI_RandomDataError)
		} else {
			_, err := rand.Read(randomData)

			if err != nil {
				fmt.Println(UI_RandomDataError, _size, err.Error())
			}

			if int64(len(randomData)) != _size {
				fmt.Println(UI_RandomDataError)
			}
		}
	} else {
		fmt.Println(fmt.Sprintf(UI_ParameterInvalid, GetFunctionName()), "_size:", _size)
	}

	return randomData
}
