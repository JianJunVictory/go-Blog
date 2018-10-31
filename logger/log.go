// logger package ,to be continued
package logger

import (
	"os"
	"log"
	"time"
)
type MyLogger struct{
	FileName *os.File
	Logger *log.Logger
}
var LOG *MyLogger

func InitLog(){
	LOG =new(MyLogger)
	LOG.setlogFile()
	LOG.Logger=log.New(LOG.FileName,"",log.Lshortfile|log.LstdFlags)
}
func (ml *MyLogger)setlogFile(){
	log.Println(ml.FileName)
	now:=time.Now().Format("2006-01-02_15")
	filename := now+"_log.log"

	if ml.FileName == nil {
		ml.FileName,_=os.Create(filename)
		return
	}
	if fileExists(filename) {
		ml.FileName,_=os.Open(filename)
		return
	}
	ml.FileName.Close()
	ml.FileName,_=os.Create(filename)
}

func (ml *MyLogger)Info(v...interface{}){
	ml.Logger.SetPrefix("[INFO]")
	LOG.setlogFile()
	ml.Logger.Println(v)
}

func (ml *MyLogger)Warn(v...interface{}){
	ml.Logger.SetPrefix("[WARN]")
	LOG.setlogFile()
	ml.Logger.Println(v)
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
