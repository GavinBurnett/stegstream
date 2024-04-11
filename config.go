package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

// Config file data
type Config struct {
	Port         int       // Port number
	StreamOnly   bool      // Stream with no hidden data
	HideOnly     bool      // Hide the file with no streaming server
	WipeAudio    bool      // Wipe audio file when server shuts down
	WipeHidden   bool      // Wipe hidden file when server shuts down
	AutoShutdown time.Time // Date and time to automatically shut down server
}

// Config file strings
const LINE_COMMENT = "#"
const LINE_CONFIG_ENTRY = "="

const PORT_CONFIG = "Port"
const STREAM_ONLY_CONFIG = "StreamOnly"
const HIDE_ONLY_CONFIG = "HideOnly"
const WIPE_AUDIO_CONFIG = "WipeAudio"
const WIPE_HIDDEN_CONFIG = "WipeHidden"
const AUTO_SHUTDOWN_CONFIG = "AutoShutdown"

// ReadConfigFile: Read config file data
func ReadConfigFile(_configFile string) Config {

	var configData Config = Config{DEFAULT_PORT, DEFAULT_STREAM_ONLY, DEFAULT_HIDE_ONLY, DEFAULT_WIPE_AUDIO, DEFAULT_WIPE_HIDDEN, DEFAULT_AUTO_SHUTDOWN}
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

								// Get stream only
								if strings.Contains(configFileReader.Text(), STREAM_ONLY_CONFIG) {
									configData.StreamOnly = ParseStringToBool(strings.Split(configFileReader.Text(), LINE_CONFIG_ENTRY)[1])
								}

								// Get hide file only
								if strings.Contains(configFileReader.Text(), HIDE_ONLY_CONFIG) {
									configData.HideOnly = ParseStringToBool(strings.Split(configFileReader.Text(), LINE_CONFIG_ENTRY)[1])
								}

								// Get wipe audio file
								if strings.Contains(configFileReader.Text(), WIPE_AUDIO_CONFIG) {
									configData.WipeAudio = ParseStringToBool(strings.Split(configFileReader.Text(), LINE_CONFIG_ENTRY)[1])
								}

								// Get wipe hidden file
								if strings.Contains(configFileReader.Text(), WIPE_HIDDEN_CONFIG) {
									configData.WipeHidden = ParseStringToBool(strings.Split(configFileReader.Text(), LINE_CONFIG_ENTRY)[1])
								}

								// Get automatically shut down server date and time
								if strings.Contains(configFileReader.Text(), AUTO_SHUTDOWN_CONFIG) {
									configData.AutoShutdown = ParseStringToDateTime(strings.Split(configFileReader.Text(), LINE_CONFIG_ENTRY)[1])
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

	// Check for auto shutdown date that is in the past
	if _config.AutoShutdown != (time.Time{}) {
		if _config.AutoShutdown.Before(time.Now().Local()) {
			fmt.Println(UI_AutoShutdownTimeInPast)
			_config.AutoShutdown = time.Time{}
		}
	}

	// Unit tests

	// Checks for entries that contradict each other - display any errors

	// Check stream only and hide only are not both true
	if _config.StreamOnly == true && _config.HideOnly == true {
		fmt.Println(UI_StreamOnlyAndHideOnlySetError)
		configDataValid = false
	}

	// Check hide only and wipe audio are not both true
	if _config.HideOnly == true && _config.WipeAudio == true {
		fmt.Println(UI_HideOnlyAndWipeAudioSetError)
		configDataValid = false
	}

	// Check stream only and wipe hidden are not both true
	if _config.StreamOnly == true && _config.WipeHidden == true {
		fmt.Println(UI_StreamOnlyAndWipeHiddenSetError)
		configDataValid = false
	}

	// Check auto shutdown and hide only are not both true
	if _config.AutoShutdown != (time.Time{}) && _config.HideOnly == true {
		fmt.Println(UI_AutoShutdownAndHideOnlySetError)
		configDataValid = false
	}

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
					fmt.Println(UI_UsingDefault, "-1")
					outputInt = -1
				}
			} else {
				fmt.Println(UI_ParseError, _inputString)
				fmt.Println(UI_UsingDefault, "-1")
				outputInt = -1
			}

		} else {
			fmt.Println(UI_ParseError, _inputString)
			fmt.Println(UI_UsingDefault, "-1")
			outputInt = -1
		}

	} else {
		fmt.Println(UI_ParseError, GetFunctionName(), _inputString)
		fmt.Println(UI_UsingDefault, "-1")
		outputInt = -1
	}

	return outputInt
}

// ParseStringToBool: Converts given string into a bool
func ParseStringToBool(_inputString string) bool {

	var parsedBool bool = false
	var parseErr error

	if len(_inputString) > 0 {

		parsedBool, parseErr = strconv.ParseBool(_inputString)

		if parseErr != nil {
			fmt.Println(UI_ParseError, _inputString)
			fmt.Println(UI_UsingDefault, "false")
		}

	} else {
		fmt.Println(UI_ParseError, GetFunctionName(), _inputString)
		fmt.Println(UI_UsingDefault, "false")
	}

	return parsedBool
}

// ParseStringToDateTime: Converts given string into a date time
func ParseStringToDateTime(_inputString string) time.Time {

	const format string = "02/01/2006 15:04"

	var loc *time.Location
	var dateTime time.Time
	var parseErr error

	if len(_inputString) > 0 {

		loc, parseErr = time.LoadLocation("Local")

		if parseErr != nil {
			fmt.Println(UI_LocaleNotFound)
			fmt.Println(UI_UsingDefault, dateTime)
		} else {

			dateTime, parseErr = time.ParseInLocation(format, _inputString, loc)

			if parseErr != nil {
				fmt.Println(UI_ParseError, _inputString)
				fmt.Println(UI_UsingDefault, dateTime)
			}
		}
	} else {
		fmt.Println(UI_ParseError, GetFunctionName(), _inputString)
		fmt.Println(UI_UsingDefault, dateTime)
	}

	return dateTime
}
