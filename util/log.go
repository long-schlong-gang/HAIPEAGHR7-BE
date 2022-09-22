/*
 *	Logger setup
 */

package util

import (
	"os"

	"github.com/kpango/glg"
)

var fileLog *os.File

func InitLogger() {

	// Add file logger
	if GlobalEnvs.LogFile != "" {
		// Check file can be written to
		fileLog := glg.FileWriter(GlobalEnvs.LogFile, LogFilePermissions)
		glg.Get().
			AddWriter(fileLog).
			SetMode(glg.BOTH)
		glg.Infof("Started logging to '%s'", GlobalEnvs.LogFile)
	}

	// Extra logger config here if necessary

	glg.Info("Initialized logger")
}

func TerminateLogger() {
	err := fileLog.Close()
	if err != nil {
		glg.Errorf("Failed to close log file '%s': %v", GlobalEnvs.LogFile, err)
	}
	glg.Info("Terminated logger")
}
