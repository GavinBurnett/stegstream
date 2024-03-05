package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// Config file data
type Config struct {
	Port int // Port number
}

// Config file strings
const LINE_COMMENT = "#"
const LINE_CONFIG_ENTRY = "="

const PORT_CONFIG = "Port"

// ReadConfigFile: Read config file data
func ReadConfigFile(_configFile string) Config {

	var configData Config = Config{DEFAULT_PORT}
	var configFile *os.File
	var configFileReader *bufio.Scanner
	var err error

	if DEBUG == true {
		fmt.Println(UI_ReadingConfigFile, _configFile)
	}

	if len(_configFile) > 0 {

		if FileExists(_configFile) {

			configFile, err = os.Open(_configFile)

			if err == nil {

				if DEBUG == true {
					fmt.Println(fmt.Sprintf(UI_FileFound, _configFile))
				}

				// Read every line in the config file and store the data in Config struct
				configFileReader = bufio.NewScanner(configFile)

				for configFileReader.Scan() {

					if len(configFileReader.Text()) > 0 {

						// Process current line in config file

						// if DEBUG == true {
						// 	fmt.Println(configFileReader.Text())
						// }

						if strings.HasPrefix(configFileReader.Text(), LINE_COMMENT) {

							// Do not process any commented out lines
							if DEBUG == true {
								fmt.Println(_configFile + UI_SkippingLine + configFileReader.Text())
							}

						} else {

							// Process current line
							if (len(configFileReader.Text())) > 0 {

								// Get Port number
								if strings.Contains(configFileReader.Text(), PORT_CONFIG) {
									configData.Port = ParseStringToInt(strings.Split(configFileReader.Text(), LINE_CONFIG_ENTRY)[1])
								}

							}
						}
					} else {
						// Line is empty
					}

				} // end for

				if DEBUG == true {
					fmt.Println(UI_Config, configData)
				}

			} else {
				fmt.Println(UI_ReadConfigFileError, GetFunctionName(), _configFile, err)
			}

			configFile.Close()

		} else {
			fmt.Println(UI_ReadConfigFileError, GetFunctionName(), _configFile, UI_FileNotFound)
		}

	} else {
		fmt.Println(UI_ReadConfigFileError, GetFunctionName(), _configFile, UI_FileNotFound)
	}

	return configData
}

// CheckConfigFile: Check config file data for invalid entries and make corrections
func CheckConfigFile(_config Config) (Config, bool) {

	var configDataValid bool = true

	// Check for invalid entries and make corrections

	// Check and correct port number
	if _config.Port <= 1024 || _config.Port >= 65535 {
		fmt.Println(fmt.Sprintf(UI_ConfigCorrection, "Port", strconv.Itoa(_config.Port), strconv.Itoa(DEFAULT_PORT)))
		_config.Port = DEFAULT_PORT
	}

	// Checks for entries that contradict each other - display any errors

	return _config, configDataValid
}

// ParseStringToInt: Converts given string into an int
func ParseStringToInt(_inputString string) int {

	var parsedInt int64 = -1
	var outputInt int = -1
	var parseErr error

	if len(_inputString) > 0 {

		parsedInt, parseErr = strconv.ParseInt(_inputString, 10, 64)

		if parseErr == nil {

			if parsedInt > 0 && parsedInt < math.MaxInt64 {

				outputInt = int(parsedInt)

				if outputInt < 0 || outputInt > math.MaxInt {
					fmt.Println(UI_ParseError, _inputString)
					outputInt = -1
				}
			} else {
				fmt.Println(UI_ParseError, _inputString)
				outputInt = -1
			}

		} else {
			fmt.Println(UI_ParseError, _inputString)
			outputInt = -1
		}

	} else {
		fmt.Println(UI_ParseError, GetFunctionName(), _inputString)
		outputInt = -1
	}

	return outputInt
}
