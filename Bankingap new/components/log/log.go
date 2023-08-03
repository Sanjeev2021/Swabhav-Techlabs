package log

import (
	"fmt"
	"os"
	"time"

)

type Logger interface {
	Print(value ...string)
	New()
}
type Log struct {
	SaveInFile bool
}

func GetLogger(saveinfile bool) *Log {
	return &Log{
		SaveInFile: saveinfile,
	}
}

func (l *Log) Print(value ...interface{}) {
	if l.SaveInFile {
		filename := GetFileName()
		data := GetDataToAppend(value...)
		err := appendToFile(filename, data)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	timestamp := time.Now().Format("2006-01-02 15:04:05")
	valuesWithTimestamp := append([]interface{}{timestamp + " -"}, value...)
	fmt.Println(valuesWithTimestamp...)
}

func GetFileName() string {
	return "C:/Users/Sanjeev Yadav/OneDrive/Desktop/Bankingap new/log.txt"
}

func appendToFile(filename string, data []byte) error {
	// Opening the file
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return err
	}
	return nil
}

func GetDataToAppend(value ...interface{}) []byte {
	var data []byte
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	data = append(data, []byte("\n")...)
	data = append(data, []byte(timestamp)...)
	data = append(data, []byte(" - ")...)
	for _, v := range value {
		data = append(data, []byte(fmt.Sprint(v))...)
	}

	return data
}
