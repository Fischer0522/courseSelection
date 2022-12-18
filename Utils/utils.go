package Utils

import (
	"bufio"
	"fmt"
	"jwb/model"
	"log"
	"os"
	"sync"
)

var mu sync.Mutex

func GenerateCookie(cookie model.Cookie) string{
	result := "JSESSIONID="+cookie.JSESS+"; BIGipServerp_new_hr_-_authserver.cumt.edu.cn="+cookie.BIGipServer+"; TWFID="+cookie.TWFID
	return result
}
func WriteLog(content string,fileName string) error{
	mu.Lock()
	outputFile, outputError := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if outputError != nil {
		fmt.Printf("An error occurred with file opening or creation\n")
		return outputError
	}
	defer mu.Unlock()
	defer outputFile.Close()

	outputWriter := bufio.NewWriter(outputFile)

	_,err := outputWriter.WriteString(content+"\n")

	if err != nil {
		log.Fatal("write log file fail")
		return err
	}

	err2 := outputWriter.Flush()
	if err2 != nil {
		log.Fatal("flushFail")
		return err2
	}
	return nil
}