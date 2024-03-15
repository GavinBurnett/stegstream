package main

const (
	UI_InvalidArgs                   = `ERROR: Invalid arguments`
	UI_FileNotFound                  = `ERROR: File not found: `
	UI_NoParametersGiven             = `ERROR: No parameters specified`
	UI_ParameterInvalid              = `ERROR: Invalid parameter: %s`
	UI_EmptyFile                     = `ERROR: Empty file`
	UI_InvalidFileSize               = `ERROR: Invalid file size: `
	UI_FileTooBig                    = `ERROR: File too big`
	UI_NoFileSize                    = `ERROR: Can't get file size: %s %s`
	UI_FileReadError                 = `ERROR: File read error: %s %s`
	UI_FileWriteError                = `ERROR: File write error: %s %s`
	UI_WebServerError                = `ERROR: Web Server error: %s%s`
	UI_WebServerNotStarted           = `ERROR: Web Server failed to start`
	UI_NoBufferMemory                = `ERROR: Failed to allocate buffer memory`
	UI_FileOpenError                 = `ERROR: File open error: `
	UI_FileDeleteError               = `ERROR: File delete error: `
	UI_FileNameTooLong               = `ERROR: File name %s too long. Maximum %v characters`
	UI_SpacingError                  = `ERROR: File to hide: %s will not fit inside: %s`
	UI_SeekFail                      = `ERROR: File: %s seek to: %v failed. %s`
	UI_HiddenDataWrittenFail         = `ERROR: Failed to write hidden file: %s into container file: %s`
	UI_NoHostname                    = `ERROR: Failed to get hostname`
	UI_ConfigFileReadError           = `ERROR: Config file read error - line: `
	UI_ReadConfigFileError           = `ERROR: Reading config file: `
	UI_FileCreateError               = `ERROR: Failed to create file: `
	UI_RandomDataError               = `ERROR: Random data not generated`
	UI_RandomNumberError             = `ERROR: Random number not generated`
	UI_ParseError                    = `ERROR: Parse failed: `
	UI_ConfigInvalid                 = `ERROR: Config invalid`
	UI_StreamOnlyAndHideOnlySetError = `ERROR: Config file StreamOnly and HideOnly both enabled`
	UI_UsingDefault                  = `Using default setting: `
	UI_ConfigFileFound               = `Config file found - using settings`
	UI_ReadingConfigFile             = `Reading config file: `
	UI_Config                        = `Config: `
	UI_ConfigCorrection              = `Config entry: %s corrected from: %s to: %s`
	UI_ProcessedStringList           = `Processed string list: `
	UI_FileSize                      = `%s File size: %v`
	UI_SpaceAvailable                = `Space available: %v`
	UI_FileToHideFits                = `File to hide: %s will fit inside: %s`
	UI_UsingSpacing                  = `Using spacing: %v`
	UI_SeekOffset                    = `Seeked to offset: %v`
	UI_EOF                           = `EOF: %s`
	UI_BytesWritten                  = `Bytes written: %v`
	UI_WritingToLocation             = `Writing to container file location: %v`
	UI_WebServerStarting             = `Web Server Starting: %s File: %s`
	UI_WebServerStarted              = `Web Server Started: %s`
	UI_ServingFile                   = `Serving File: %s`
	UI_Arguments                     = `Arguments: `
	UI_FileFound                     = `File found: %s`
	UI_CreatingFile                  = `Creating file: `
	UI_Done                          = ` - Done.`
	UI_Parameter                     = `Parameters: %s`
	UI_HiddenDataWriting             = `Hiding file: %s in container file: %s `
	UI_HiddenDataWrittenOK           = `Hidden file: %s written into container file: %s successfully`
	UI_SkippingLine                  = ` Skipping Line: `
	UI_WaitingForShutdown            = `Waiting for shutdown`
	UI_ShutdownSignal                = `Shutdown signal detected`
	UI_ShuttingDown                  = `Shutting down`
	UI_CtrlCToExit                   = `Press CTRL+C or kill stegstream-server process to exit`
	UI_HiddenFileData                = `Hidden file data:
Spacing: %v
Steps: %v
File name: %s`
	UI_Help = `stegstream server v1.2 by gburnett@outlook.com

Arguments: 

./stegstream-server <container file> <file to hide>

Example:

./stegstream-server Waves.mp3 HideFile.txt`
)
