package logger

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"student_server/config"
	"student_server/lib/file"
	"time"
)

func InitLog() {

	driver := config.Logger.Driver
	var fileName string

	switch driver {
	case "daily":
		fileName = "log-" + time.Now().Format("2006-01-02") + ".log"
	case "single":
		fileName = "log.log"
	default:
		fileName = "log.log"
	}
	filePath := config.Logger.Path + "/" + fileName
	if !file.FileExists(filePath) {
		err := file.CreateFile(config.Logger.Path, fileName)
		if err != nil {
			fmt.Printf("create logFile error=%s", err.Error())
		}
	}
	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		fmt.Printf("open logFile error=%s", err.Error())
	}
	log.SetOutput(f)
}

func Info(msg string, data interface{}) {
	level := "[Info]"
	write(level, msg, data)
}

func Error(msg string, data interface{}) {
	level := "[Error]"
	write(level, msg, data)
}

func write(level string, msg string, data interface{}) {
	if data != nil {
		data, err := json.Marshal(data)
		if err != nil {
			fmt.Println("json err:", err)
		}
		log.Printf("%s %s %s", level, msg, data)
	} else {
		log.Printf("%s %s", level, msg)
	}
}
