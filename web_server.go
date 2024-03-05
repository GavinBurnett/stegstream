package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"sync"
)

// Server active flag
var ServerUp bool = false

// File for web server to serve
var FileToServe string = ""

// Url of web server
var Url string = ""

// ServeFile: Serves file to web clients
func ServeFile(w http.ResponseWriter, r *http.Request) {

	if DEBUG == true {
		fmt.Println(fmt.Sprintf(UI_ServingFile, FileToServe))
	}
	http.ServeFile(w, r, FileToServe)
}

// StartWebServer: Start web server and listen for connections
func StartWebServer(_fileToServe string, _configData Config, _wait *sync.WaitGroup) {

	var hostname string = ""
	var serverHostname string = ""
	var fileType string = "/Audio"
	var err error

	if len(_fileToServe) > 0 {

		hostname, err = os.Hostname()

		if err != nil {
			fmt.Println(UI_NoHostname)

			// Signal calling method that web server start has failed
			_wait.Done()
		} else {

			// Set hostname name and port
			serverHostname = hostname + ":" + strconv.Itoa(_configData.Port)

			// Set file to serve and web server url
			FileToServe = _fileToServe
			Url = "http://" + serverHostname + fileType

			if DEBUG == true {
				fmt.Println(fmt.Sprintf(UI_WebServerStarting, Url, FileToServe))
			}

			// Set up web server
			mux := http.NewServeMux()

			mux.HandleFunc(fileType, ServeFile)

			ServerUp = true

			if DEBUG == true {
				fmt.Println(fmt.Sprintf(UI_WebServerStarted, Url))
			}

			// Signal calling method that web server start has completed
			_wait.Done()

			// Listen for client connections and serve file
			err = http.ListenAndServe(serverHostname, mux)

			if err != nil {
				fmt.Println(fmt.Sprintf(UI_WebServerError, err, ""))

				// Signal calling method that web server start has failed
				_wait.Done()

				ServerUp = false
			}
		}

	} else {
		fmt.Print(fmt.Sprintf(UI_ParameterInvalid, GetFunctionName()))
		fmt.Println(fmt.Sprintf(UI_Parameter, "_fileToServe: "+_fileToServe))

		ServerUp = false
	}
}
