package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
)

// HiddenFileData: Data structure that describes the file hidden in the container file
type HiddenFileData struct {
	magicNumber int64
	fileName    [FILENAME_LENGTH]byte
	steps       int64
	spacing     int64
}

// Steg: Places the hide file data in to the container file data
func Steg(_containerFile string, _hideFile string) bool {

	var fileHidden bool = true
	var containerFileSize int64 = 0
	var hideFileSize int64 = 0
	const START_OFFSET int64 = 2000
	const BUFFER_SIZE int64 = 1000
	var spaceAvailable int64 = 0
	var spacing int64 = 0
	var countSpacing int64 = 0
	var containerFile *os.File
	var hiddenFile *os.File
	var buffer []byte
	var writeCounter int64 = 0
	var writeByte []byte
	var err error
	var containerFileWriteError bool = false
	var hiddenFileBytesRead int = 0
	var hiddenFileTotalBytesRead int = 0
	var containerFileBytesWritten int = 0
	var containerFileTotalBytesWritten int64 = 0
	var stepsCount int64 = 0

	// If file names for container and file to hide have been passed in
	if len(_containerFile) > 0 && len(_hideFile) > 0 {

		// If container file exists
		if FileExists(_containerFile) {

			// If file to hide exists
			if FileExists(_hideFile) {

				// Get container file size
				containerFileSize = GetFileSize(_containerFile)

				if DEBUG == true {
					fmt.Println(fmt.Sprintf(UI_FileSize, _containerFile, containerFileSize))
				}

				// If container file size is valid
				if containerFileSize != -1 && containerFileSize > 0 {

					// Get file to hide file size
					hideFileSize = GetFileSize(_hideFile)

					if DEBUG == true {
						fmt.Println(fmt.Sprintf(UI_FileSize, _hideFile, hideFileSize))
					}

					// If file to hide file size is valid
					if hideFileSize != -1 && hideFileSize > 0 {

						// Check if file to hide will fit into container file
						// File to hide size must be no bigger than 10% of the container file size minus the start offset
						spaceAvailable = containerFileSize - START_OFFSET

						if spaceAvailable > 0 {

							spaceAvailable = spaceAvailable * 10 / 100

							if spaceAvailable > 0 {

								if DEBUG == true {
									fmt.Println(fmt.Sprintf(UI_SpaceAvailable, spaceAvailable))
								}

								if spaceAvailable > hideFileSize {

									// File to hide will fit inside container file
									if DEBUG == true {
										fmt.Println(fmt.Sprintf(UI_FileToHideFits, _hideFile, _containerFile))
									}

									// Calculate spacing - locations in container file to write hidden file data to
									spacing = (containerFileSize - START_OFFSET) / hideFileSize

									if spacing > 0 {

										if DEBUG == true {
											fmt.Println(fmt.Sprintf(UI_UsingSpacing, spacing))
										}

										// Set initial location in container file to write hidden file data to
										countSpacing = START_OFFSET + 1

										// Open container file
										containerFile, err = os.OpenFile(_containerFile, os.O_WRONLY, 0)

										if err == nil {

											// Open file to hide
											hiddenFile, err = os.OpenFile(_hideFile, os.O_RDONLY, 0)

											if err == nil {

												// Work through file to hide writing its data into the container file
												for {

													// Set up buffer to read file to hide data into
													if hideFileSize < BUFFER_SIZE {
														buffer = make([]byte, hideFileSize)
													} else {
														buffer = make([]byte, BUFFER_SIZE)
													}

													if int64(len(buffer)) == BUFFER_SIZE || int64(len(buffer)) == hideFileSize {

														// Read the file to hide data into the buffer
														hiddenFileBytesRead, err = hiddenFile.Read(buffer)
														hiddenFileTotalBytesRead += hiddenFileBytesRead

														if err != nil {

															if err == io.EOF {
																// If end of file to hide found, stop writing its data to the container file
																if DEBUG == true {
																	fmt.Println(fmt.Sprintf(UI_EOF, _hideFile))
																}
																break
															} else {
																// Error reading file to hide - display an error and stop writing data
																fmt.Println(fmt.Sprintf(UI_FileReadError, _containerFile, err))
																fileHidden = false
																containerFileWriteError = true
																break
															}
														} else {

															// Loop around file to hide buffer writing buffer data to container file one byte at a time
															for writeCounter = 0; writeCounter != int64(len(buffer)); writeCounter++ {

																if UPDATE_UI == true {
																	// Update UI
																	fmt.Printf("\r" + fmt.Sprintf(UI_HiddenDataWriting, _hideFile, _containerFile) + ".  ")
																}

																// Get a byte from the buffer
																writeByte = append(writeByte, buffer[writeCounter])

																if DEBUG == true {
																	fmt.Println(fmt.Sprintf(UI_WritingToLocation, countSpacing))
																}

																// Write the byte to container file at spacing location
																containerFileBytesWritten, err = containerFile.WriteAt(writeByte, countSpacing)

																stepsCount++

																// Move on to the next spacing location
																countSpacing = countSpacing + spacing

																if UPDATE_UI == true {
																	// Update UI
																	fmt.Printf("\r" + fmt.Sprintf(UI_HiddenDataWriting, _hideFile, _containerFile) + ".. ")
																}

																if err != nil {

																	if err == io.EOF {
																		// If end of container file found, display an error and stop writing to it and display end of file
																		fmt.Println(fmt.Sprintf(UI_EOF, _containerFile))
																		fileHidden = false
																		containerFileWriteError = true
																		break
																	} else {
																		// Error writing to container file - display an error and stop writing data
																		fmt.Println(fmt.Sprintf(UI_FileWriteError, _containerFile, err))
																		fileHidden = false
																		containerFileWriteError = true
																		break
																	}
																} else {

																	// If the byte not written to the container file - display an error and stop writing data
																	if containerFileBytesWritten != len(writeByte) {
																		fmt.Println(fmt.Sprintf(UI_FileWriteError, _containerFile, ""))
																		fileHidden = false
																		containerFileWriteError = true
																		break
																	} else {

																		// Clear the byte out
																		writeByte = nil

																		// Keep track of amount of data written to the container file
																		containerFileTotalBytesWritten = containerFileTotalBytesWritten + int64(containerFileBytesWritten)

																		// If amount of data written to the container file is the same as the file to hide size, all file to hide data has been written to the container file - stop writing data
																		if containerFileTotalBytesWritten == hideFileSize {

																			if DEBUG == true {
																				fmt.Println(fmt.Sprintf(UI_HiddenDataWrittenOK, _hideFile, _containerFile))
																			}

																			break
																		}
																	}
																}

																if UPDATE_UI == true {
																	// Update UI
																	fmt.Printf("\r" + fmt.Sprintf(UI_HiddenDataWriting, _hideFile, _containerFile) + "...")
																}

															} // end byte at a time for loop

															if DEBUG == true {
																if containerFileWriteError == true {
																	// Container file written to with errors
																	fmt.Println(fmt.Sprintf(UI_FileWriteError, _containerFile, ""))
																}
															}
														}
													} else {
														// Failed to allocate buffer memory
														fmt.Println(UI_NoBufferMemory)
														fileHidden = false
														containerFileWriteError = true
														break
													}

													buffer = nil

												} // End read file to hide for loop

												// Write hidden file data to container file
												if containerFileWriteError == false {
													if WriteHiddenFileData(containerFile, _hideFile, stepsCount, spacing) == false {
														fmt.Println(fmt.Sprintf(UI_FileWriteError, containerFile.Name(), ""))
														fileHidden = false
													}
												}

												if UPDATE_UI == true {
													// Update UI
													fmt.Printf("\r")
												}

											} else {
												// Open file to hide error
												fmt.Println(UI_FileOpenError, err)
												fileHidden = false
											}

										} else {
											// Open container file error
											fmt.Println(UI_FileOpenError, err)
											fileHidden = false
										}

										// File to hide data written to contain file OK - close file handles down and clear buffer
										containerFile.Sync()
										containerFile.Close()

										hiddenFile.Close()

										buffer = nil

										if DEBUG == true {
											fmt.Println(fmt.Sprintf(UI_HiddenFileData, spacing, stepsCount, _hideFile))
										}

									} else {
										fmt.Println(fmt.Sprintf(UI_SpacingError, _hideFile, _containerFile))
										fileHidden = false
									}
								} else {
									fmt.Println(fmt.Sprintf(UI_SpacingError, _hideFile, _containerFile))
									fileHidden = false
								}
							} else {
								fmt.Println(fmt.Sprintf(UI_SpacingError, _hideFile, _containerFile))
								fileHidden = false
							}
						} else {
							fmt.Println(fmt.Sprintf(UI_SpacingError, _hideFile, _containerFile))
							fileHidden = false
						}
					} else {
						fmt.Println(UI_InvalidFileSize, _hideFile)
						fileHidden = false
					}
				} else {
					fmt.Println(UI_InvalidFileSize, _containerFile)
					fileHidden = false
				}
			} else {
				fmt.Println(UI_FileNotFound, _hideFile)
				fileHidden = false
			}
		} else {
			fmt.Println(UI_FileNotFound, _containerFile)
			fileHidden = false
		}
	} else {
		fmt.Print(fmt.Sprintf(UI_ParameterInvalid, GetFunctionName()))
		fmt.Println(fmt.Sprintf(UI_Parameter, "_containerFile: "+_containerFile+"_hideFile: "+_hideFile))
		fileHidden = false
	}

	return fileHidden
}

// WriteHiddenFileData: Writes the hidden file data to the container file
func WriteHiddenFileData(_containerFile *os.File, _filename string, _steps int64, _spacing int64) bool {

	var hiddenFileData HiddenFileData
	var filenameNoPath string
	var containerFileSize int64 = -1
	var seekOffset int64 = -1
	var dataWritten = false
	var writeContainerFileSize []byte
	const WRITE_CONTAINER_FILE_SIZE_BUFFER int = 8
	var bytesWritten int = -1
	var err error

	if _containerFile != nil && len(_filename) > 0 && _steps != -1 && _spacing != -1 {

		if len(_filename) > FILENAME_LENGTH {
			fmt.Println(fmt.Sprintf(UI_FileNameTooLong, _filename, FILENAME_LENGTH))
			dataWritten = false
		} else {

			// Set up hidden file data structure to write to container file
			hiddenFileData = HiddenFileData{}

			hiddenFileData.magicNumber = MAGIC_NUMBER
			filenameNoPath = filepath.Base(_filename)
			copy(hiddenFileData.fileName[:], filenameNoPath)
			hiddenFileData.steps = _steps
			hiddenFileData.spacing = _spacing

			if DEBUG == true {
				fmt.Println(fmt.Sprintf(UI_HiddenFileData, hiddenFileData.spacing, hiddenFileData.steps, hiddenFileData.fileName))
			}

			// Get size of container file
			containerFileSize = GetFileSize(_containerFile.Name())

			if DEBUG == true {
				fmt.Println(fmt.Sprintf(UI_FileSize, _containerFile.Name(), containerFileSize))
			}

			if containerFileSize == -1 {
				fmt.Println(fmt.Sprintf(UI_NoFileSize, _containerFile.Name(), ""))
				dataWritten = false
			} else {

				// Seek to end of container file
				seekOffset, err = _containerFile.Seek(containerFileSize, 0)

				if DEBUG == true {
					fmt.Println(fmt.Sprintf(UI_SeekOffset, seekOffset))
				}
				if err != nil {
					fmt.Println(fmt.Sprintf(UI_SeekFail, _containerFile.Name(), seekOffset, err))
					dataWritten = false
				} else {

					// Write hidden file data to end of container file
					err = binary.Write(_containerFile, binary.LittleEndian, hiddenFileData)
					if err != nil {
						fmt.Println(fmt.Sprintf(UI_FileWriteError, _containerFile.Name(), err))
						dataWritten = false
					} else {
						if DEBUG == true {
							fmt.Println(fmt.Sprintf(UI_HiddenDataWrittenOK, _filename, _containerFile.Name()))
						}
						dataWritten = true
					}
				}
			}

			if dataWritten == true {

				// Write location of hidden file data in container file to fixed offset in container file
				writeContainerFileSize = make([]byte, WRITE_CONTAINER_FILE_SIZE_BUFFER)

				if len(writeContainerFileSize) != WRITE_CONTAINER_FILE_SIZE_BUFFER {
					fmt.Println(UI_NoBufferMemory)
					dataWritten = false
				} else {

					binary.BigEndian.PutUint64(writeContainerFileSize, uint64(containerFileSize))
					bytesWritten, err = _containerFile.WriteAt(writeContainerFileSize, HIDDEN_FILE_DATA_OFFSET)

					if err != nil {
						fmt.Println(fmt.Sprintf(UI_FileWriteError, _containerFile.Name(), ""))
						dataWritten = false
					} else {

						if DEBUG == true {
							fmt.Println(fmt.Sprintf(UI_BytesWritten, bytesWritten))
						}

						if bytesWritten != WRITE_CONTAINER_FILE_SIZE_BUFFER {
							fmt.Println(fmt.Sprintf(UI_FileWriteError, _containerFile.Name(), ""))
							dataWritten = false
						} else {
							if DEBUG == true {
								fmt.Println(fmt.Sprintf(UI_HiddenDataWrittenOK, _filename, _containerFile.Name()))
							}
							dataWritten = true
						}
					}
				}
			}
		}
	} else {
		fmt.Print(fmt.Sprintf(UI_ParameterInvalid, GetFunctionName()))
		fmt.Println(fmt.Sprintf(UI_Parameter, "containerFile: "+_containerFile.Name()+"filename: "+_filename+"steps: "+strconv.FormatInt(_steps, 10)+"spacing:"+strconv.FormatInt(_spacing, 10)))
	}

	return dataWritten
}
