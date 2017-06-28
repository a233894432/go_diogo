package main

import "github.com/ivpusic/golog"
import "github.com/ivpusic/golog/appenders"

func main() {
	logger := golog.Default
	logger = golog.GetLogger("DIO")
	// make instance of file appender and enable it
	logger.Enable(appenders.File(golog.Conf{
		// file in which logs will be saved
		"path": "./log/log.txt",
	}))

	logger.Debug("some message")
}
