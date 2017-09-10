package logging

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/fatih/color"
)

type Logger struct {
	infoFile, errorFile *os.File
}

func New(infoFilePath string, errorFilePath string) *Logger {

	infoFile, err := os.OpenFile(infoFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
	if err != nil {
		log.Println(err)
		return nil
	}

	errorFile, err := os.OpenFile(errorFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
	if err != nil {
		log.Println(err)
		return nil
	}

	logger := &Logger{infoFile, errorFile}

	return logger
}

func (self *Logger) Debug(v ...interface{}) {
	log := fmt.Sprintf("%s %s %s", time.Now().Format("2006-01-02 15:04:05"), "[DEBUG]", fmt.Sprint(v...))
	fmt.Println(color.YellowString(log))
}

func (self *Logger) Info(v ...interface{}) {
	log := fmt.Sprintf("%s %s %s", time.Now().Format("2006-01-02 15:04:05"), "[INFO]", fmt.Sprint(v...))
	fmt.Println(color.CyanString(log))
	if self.infoFile != nil {
		self.infoFile.WriteString(log + "\n")
	}
}

func (self *Logger) Error(v ...interface{}) {
	log := fmt.Sprintf("%s %s %s", time.Now().Format("2006-01-02 15:04:05"), "[ERROR]", fmt.Sprint(v...))
	fmt.Println(color.RedString(log))
	if self.errorFile != nil {
		self.errorFile.WriteString(log + "\n")
	}
	if self.infoFile != nil {
		self.infoFile.WriteString(log + "\n")
	}
}
