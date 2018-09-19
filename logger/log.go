// logger package ,to be continued
package logger

import (
	"io"
	"log"
	"os"
)

// Logger log
type Logger struct {
	logger *log.Logger
}

// Info level
func (log *Logger) Info(info interface{}) {
	log.logger.SetPrefix("[INFO]")
	log.logger.Println(info)
}

// Warn level
func (log *Logger) Warn(info interface{}) {
	log.logger.SetPrefix("[Warn]")
	log.logger.Println(info)
}

// SetLogFile set log file
func (log *Logger) SetLogFile(out io.Writer) {
	log.logger.SetOutput(out)
}

func fileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func createFile(name string) error {
	fo, err := os.Create(name)
	if err != nil {
		return err
	}
	defer func() {
		fo.Close()
	}()
	return nil
}
