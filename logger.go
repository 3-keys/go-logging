package logging

import (
	"fmt"
	"github.com/fatih/color"
	"os"
	"path"
	"time"
)

type Logger struct {
	infoFilePath, errorFilePath string
}

func (logger *Logger) getInfoFile() (*os.File, error) {
	return os.OpenFile(logger.infoFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
}
func (logger *Logger) getErrorFile() (*os.File, error) {
	return os.OpenFile(logger.errorFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
}

func New(infoFilePath string, errorFilePath string) *Logger {

	err := os.MkdirAll(path.Dir(infoFilePath), 0755)
	if err != nil {
		fmt.Print(err)
	}

	err = os.MkdirAll(path.Dir(errorFilePath), 0755)
	if err != nil {
		fmt.Println(err)
	}

	logger := &Logger{infoFilePath, errorFilePath}

	return logger
}

func (logger *Logger) Debug(v ...interface{}) {
	log := fmt.Sprintf("%s %s %s", time.Now().Format("2006-01-02 15:04:05"), "[DEBUG]", fmt.Sprintln(v...))
	fmt.Print(color.YellowString(log))
}

func (logger *Logger) Verbose(v ...interface{}) {
	log := fmt.Sprintf("%s %s %s", time.Now().Format("2006-01-02 15:04:05"), "[VERBOSE]", fmt.Sprintln(v...))
	fmt.Print(color.MagentaString(log))
}

func (logger *Logger) Info(v ...interface{}) {
	log := fmt.Sprintf("%s %s %s", time.Now().Format("2006-01-02 15:04:05"), "[INFO]", fmt.Sprintln(v...))
	fmt.Print(color.BlueString(log))

	infoFile, err := logger.getInfoFile()

	if err != nil {
		fmt.Println(err)
	} else {
		infoFile.WriteString(log)
	}

	defer infoFile.Close()

}

func (logger *Logger) Error(v ...interface{}) {
	log := fmt.Sprintf("%s %s %s", time.Now().Format("2006-01-02 15:04:05"), "[ERROR]", fmt.Sprintln(v...))
	fmt.Print(color.RedString(log))

	errorFile, err := logger.getErrorFile()

	if err != nil {
		fmt.Println(err)
	} else {
		errorFile.WriteString(log)
	}

	defer errorFile.Close()
}

func (logger *Logger) Fatal(v ...interface{}) {

	logger.Error(v)
	os.Exit(1)

}
