package utils

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"time"
)

func NewLogger() *log.Logger {
	date := time.Now().UTC().Format("01-02-2006")
	filename := fmt.Sprintf("%s_log.txt", date)

	filepath := path.Join("log", filename)

	fileWriter, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		log.Fatalf("could not open the file: %v", err)
	}

	writers := io.MultiWriter(os.Stdin, fileWriter)

	return log.New(writers, "", log.Ldate|log.Ltime)
}
