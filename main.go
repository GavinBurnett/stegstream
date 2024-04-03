// stegstream-server project main.go
package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// Main: program entry point
func main() {

	var configData Config = Config{DEFAULT_PORT, DEFAULT_STREAM_ONLY, DEFAULT_HIDE_ONLY, DEFAULT_WIPE_AUDIO, DEFAULT_WIPE_HIDDEN}
	var configDataValid bool = false
	var stegOK bool = false
	var waitWebServer sync.WaitGroup
	exitCode := 0

	// Hardcode command line arguments for testing
	// testArgs := []string{"", "", ""}
	// testArgs[0] = "./stegstream-server"
	// testArgs[1] = "/home/gkb/go/stegstream_files/Waves.mp3"
	// testArgs[2] = "/home/gkb/go/stegstream_files/Burning Chrome.txt"
	// os.Args = testArgs

	if os.Args != nil {

		if DEBUG == true {
			fmt.Println(len(os.Args), UI_Arguments, os.Args)
		}

		if len(os.Args) == 1 {
			// No user arguments given - display help
			fmt.Println(UI_Help)
		}
		if len(os.Args) == 2 {
			if IsStringHelpArgument(os.Args[1]) {
				// User has given help argument - display help
				fmt.Println(UI_Help)
			} else {
				// User has given only one argument that is not a help argument - display error
				exitCode = 1
				fmt.Println(UI_InvalidArgs)
			}
		}
		if len(os.Args) == 3 {

			// Check if config file is present
			if FileExists(CONFIG_FILE) {
				// Config file found - use the settings in it
				fmt.Println(UI_ConfigFileFound)
				configData = ReadConfigFile(CONFIG_FILE)
			}
			// Config file not found - use default settings

			// Check config is valid and correct any errors if possible
			configData, configDataValid = CheckConfigFile(configData)

			if configDataValid == false {
				// Config data not valid
				exitCode = 1
				fmt.Println(UI_ConfigInvalid)
			} else {

				if DEBUG == true {
					fmt.Println(UI_Config, configData)
				}

				// If stream only config setting is false
				if configData.StreamOnly == false {

					// Hide file to hide inside container file
					stegOK = Steg(os.Args[1], os.Args[2])

					if stegOK == true {
						fmt.Println(fmt.Sprintf(UI_HiddenDataWrittenOK, os.Args[2], os.Args[1]))
					}
				}

				// If file to hide has been hidden ok, or stream only setting is true
				if stegOK == true || configData.StreamOnly == true {

					// Start web server if hide file only is false
					if configData.HideOnly == false {

						// If container file exists
						if FileExists(os.Args[1]) {

							// Start web server
							waitWebServer.Add(1)
							go StartWebServer(os.Args[1], configData, &waitWebServer)
							waitWebServer.Wait()

							if ServerUp == true {

								// Web server started - listen for shutdown signal
								fmt.Println(fmt.Sprintf(UI_WebServerStarted, Url))
								if configData.WipeAudio == true {
									fmt.Println(fmt.Sprintf(UI_WipeAudioWarning, os.Args[1]))
								}
								if configData.WipeHidden == true {
									fmt.Println(fmt.Sprintf(UI_WipeHiddenWarning, os.Args[2]))
								}
								fmt.Println(UI_CtrlCToExit)

								WaitForShutdown()

								// Wipe audio and hidden files if set in config
								if configData.WipeAudio == true {
									if WipeFile(os.Args[1]) == true {
										fmt.Println(fmt.Sprintf(UI_WipedFile, os.Args[1]))
									} else {
										fmt.Println(fmt.Sprintf(UI_WipeFileError, os.Args[1]))
									}
								}

								if configData.WipeHidden == true {
									if WipeFile(os.Args[2]) == true {
										fmt.Println(fmt.Sprintf(UI_WipedFile, os.Args[2]))
									} else {
										fmt.Println(fmt.Sprintf(UI_WipeFileError, os.Args[2]))
									}
								}

							} else {
								exitCode = 1
								fmt.Println(UI_WebServerNotStarted)
							}

						} else {
							exitCode = 1
							fmt.Println(UI_FileNotFound, os.Args[1])
						}
					}

				} else {
					exitCode = 1
					fmt.Println(fmt.Sprintf(UI_HiddenDataWrittenFail, os.Args[2], os.Args[1]))
				}
			}
		}
		if len(os.Args) > 3 {
			// Too many arguments - display error
			exitCode = 1
			fmt.Println(UI_InvalidArgs)
		}

	} else {
		// No arguments
		exitCode = 1
		fmt.Println(UI_NoParametersGiven)
	}

	os.Exit(exitCode)
}

// WaitForShutdown: Waits for CTRL+C or external process kill command
func WaitForShutdown() {

	var shutdown bool = false
	var killChannel chan (os.Signal)

	if DEBUG == true {
		fmt.Println(UI_WaitingForShutdown)
	}

	// Create channel to listen for shutdown signal
	killChannel = make(chan os.Signal)
	signal.Notify(killChannel, os.Interrupt, syscall.SIGTERM)

	// Set stub function to listen on channel and set shutdown flag
	go func() {
		<-killChannel
		shutdown = true
		fmt.Println(UI_ShuttingDown)
	}()

	// Loop until shutdown flag set
	for {
		time.Sleep(10 * time.Second)
		if shutdown == true {
			if DEBUG == true {
				fmt.Println(UI_ShutdownSignal)
			}
			break
		}
	} // end for loop
}

// IsStringHelpArgument: Returns true if given string is a help argument, false if it is not
func IsStringHelpArgument(_theString string) bool {

	isHelpArgument := false

	if len(_theString) > 0 {

		switch _theString {
		case "?":
			isHelpArgument = true
		case "/?":
			isHelpArgument = true
		case "-?":
			isHelpArgument = true
		case "--?":
			isHelpArgument = true
		case "h":
			isHelpArgument = true
		case "/h":
			isHelpArgument = true
		case "-h":
			isHelpArgument = true
		case "--h":
			isHelpArgument = true
		case "help":
			isHelpArgument = true
		case "/help":
			isHelpArgument = true
		case "-help":
			isHelpArgument = true
		case "--help":
			isHelpArgument = true
		}

	} else {
		fmt.Print(fmt.Sprintf(UI_ParameterInvalid, GetFunctionName()))
		fmt.Println(fmt.Sprintf(UI_Parameter, "_theString:"+_theString))
	}

	return isHelpArgument
}
