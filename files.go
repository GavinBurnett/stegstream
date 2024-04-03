package main

import (
	"fmt"
	"io"
	"math"
	"math/rand"
	"os"
)

// FileExists: Check the given file exists
func FileExists(_file string) bool {

	var fileExists bool = false

	if len(_file) > 0 {

		// Try to get file info
		_, err := os.Stat(_file)

		// If any errors occur on getting file info, file does not exist
		if os.IsNotExist(err) || err != nil {
			if DEBUG == true {
				fmt.Println(UI_FileNotFound, _file)
			}
			fileExists = false
		} else {
			// File info found - file exists
			if DEBUG == true {
				fmt.Println(fmt.Sprintf(UI_FileFound, _file))
			}
			fileExists = true
		}
	} else {
		fmt.Print(fmt.Sprintf(UI_ParameterInvalid, GetFunctionName()))
		fmt.Println(fmt.Sprintf(UI_Parameter, "_file: "+_file))
	}

	return fileExists
}

// GetFileSize: Gets the size of the given file in bytes
func GetFileSize(_file string) int64 {

	var fileSize int64 = -1

	if len(_file) > 0 {

		// Get file info
		fileInfo, err := os.Stat(_file)
		if err == nil {

			// Get file size
			fileSize = fileInfo.Size()

			if fileSize == 0 {
				fmt.Println(UI_EmptyFile)
				fileSize = -1
			}

			if fileSize < 0 {
				fmt.Println(UI_InvalidFileSize, _file)
				fileSize = -1
			}

			if fileSize > math.MaxInt64 {
				fmt.Println(UI_FileTooBig)
				fileSize = -1
			}

		} else {
			fmt.Println(fmt.Sprintf(UI_NoFileSize, _file, err))
		}

	} else {
		fmt.Println(fmt.Sprintf(UI_ParameterInvalid, GetFunctionName()))
		fmt.Println(fmt.Sprintf(UI_Parameter, "_file:"+_file))
	}

	return fileSize
}

// DeleteFile: Delete the given file
func DeleteFile(_file string) bool {

	var deleteError bool = false
	var err error

	if len(_file) > 0 {
		err = os.Remove(_file)
		if err != nil {
			fmt.Println(UI_FileDeleteError, _file, err)
			deleteError = true
		} else {
			// File deleted
			if DEBUG == true {
				fmt.Println(UI_FileDeleted, _file)
			}
		}
	} else {
		fmt.Println(fmt.Sprintf(UI_ParameterInvalid, GetFunctionName()), "_file:", _file)
		deleteError = true
	}

	return deleteError
}

// WipeFile: Wipes the given file
func WipeFile(_file string) bool {

	var file *os.File
	var fileSize int64 = 0
	var currentWipePass int = 0
	const WIPE_PASSES int = 5
	const BUFFER_SIZE int64 = 1000
	var randomDataCount int = 0
	var buffer []byte
	var bytesWritten int = 0
	var totalBytesWritten int64 = 0
	var wipeError bool = false
	var fileWiped bool = false
	var err error

	if len(_file) > 0 {

		if DEBUG == true {
			fmt.Println(fmt.Sprintf(UI_WipingFile, _file))
		}

		// Open file
		file, err = os.OpenFile(_file, os.O_WRONLY, 0)

		if err == nil {

			// Get file size
			fileSize = GetFileSize(_file)

			if DEBUG == true {
				fmt.Println(fmt.Sprintf(UI_FileSize, _file, fileSize))
			}

			if fileSize != -1 {

				// Wipe the file multiple times
				for currentWipePass = 0; currentWipePass != WIPE_PASSES; currentWipePass++ {

					// Ovewrite file with data in buffer
					for totalBytesWritten = 0; totalBytesWritten != fileSize; totalBytesWritten++ {

						if UPDATE_UI == true {
							// Update UI
							fmt.Printf("\r" + fmt.Sprintf(UI_WipingFile, _file) + ".  ")
						}

						// Create write buffer
						if (totalBytesWritten + BUFFER_SIZE) > fileSize {
							// If data left to write out is smaller than the buffer, set the buffer size to the size of data left
							buffer = make([]byte, (fileSize - totalBytesWritten))

							if int64(len(buffer)) != (fileSize - totalBytesWritten) {
								// Failed to allocate buffer memory
								fmt.Println(UI_NoBufferMemory)
								wipeError = true
								break
							}
						} else {
							buffer = make([]byte, BUFFER_SIZE)

							if int64(len(buffer)) != BUFFER_SIZE {
								// Failed to allocate buffer memory
								fmt.Println(UI_NoBufferMemory)
								wipeError = true
								break
							}
						}

						if DEBUG == true {
							fmt.Println(fmt.Sprintf(UI_BufferSize, len(buffer)))
						}

						// Fill buffer with random data
						randomDataCount, err = rand.Read(buffer)

						if UPDATE_UI == true {
							// Update UI
							fmt.Printf("\r" + fmt.Sprintf(UI_WipingFile, _file) + ".. ")
						}

						if err == nil {

							if randomDataCount == len(buffer) {

								// Overwrite file with buffer contents
								bytesWritten, err = file.WriteAt(buffer, totalBytesWritten)

								if err != nil {

									if err == io.EOF {
										// If end of file found, display an error and stop writing to it and display end of file
										fmt.Println(fmt.Sprintf(UI_EOF, _file))
										wipeError = true
										break
									} else {
										// Error writing to file - display an error and stop writing data
										fmt.Println(fmt.Sprintf(UI_FileWriteError, _file, err))
										wipeError = true
										break
									}
								} else {
									// If the buffer has not been written to the file - display an error and stop writing data
									if bytesWritten != len(buffer) {
										fmt.Println(fmt.Sprintf(UI_FileWriteError, _file, ""))
										wipeError = true
										break
									} else {
										// Keep track of amount of data written to the file
										totalBytesWritten = totalBytesWritten + int64(bytesWritten)

										if DEBUG == true {
											fmt.Println(fmt.Sprintf(UI_BytesWritten, totalBytesWritten))
										}

										// If amount of data written to the file is the same as the file size, all data has been written to the file - stop writing data
										if totalBytesWritten == fileSize {

											if DEBUG == true {
												fmt.Println(UI_DataWriteComplete)
												fmt.Println(fmt.Sprintf(UI_FileSize, _file, fileSize))
											}

											break
										}
									}
								}

								buffer = nil

							} else {
								fmt.Println(UI_RandomDataError)
								wipeError = true
								break
							}

						} else {
							fmt.Println(UI_RandomDataError)
							wipeError = true
							break
						}

						if UPDATE_UI == true {
							// Update UI
							fmt.Printf("\r" + fmt.Sprintf(UI_WipingFile, _file) + "...")
						}

					} // end buffer loop

					// If error has occured stop wiping file
					if wipeError == true {
						break
					}

				} // end wipe pass loop

			} else {
				fmt.Println(UI_InvalidFileSize, _file)
			}

		} else {
			fmt.Println(UI_FileOpenError, err)
		}

		// Close file handle down and clear buffer
		file.Sync()
		file.Close()

		// If all wipe passes have completed without error, delete file
		if wipeError == false {
			if DeleteFile(_file) == false {
				fileWiped = true
				if DEBUG == true {
					fmt.Println(fmt.Sprintf(UI_WipedFile, _file))
				}
			} else {
				if DEBUG == true {
					fmt.Println(fmt.Sprintf(UI_WipeFileError, _file))
				}
			}
		} else {
			if DEBUG == true {
				fmt.Println(fmt.Sprintf(UI_WipeFileError, _file))
			}
		}

	} else {
		fmt.Print(fmt.Sprintf(UI_ParameterInvalid, GetFunctionName()))
		fmt.Println(fmt.Sprintf(UI_Parameter, "_file: "+_file))
	}

	return fileWiped
}
