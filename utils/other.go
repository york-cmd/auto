package utils

import (
	"fmt"
	"log"
	"os"
)

func Pwd() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("%v", err)
	}
	return dir
}
func CheckError(name string, err error) {
	if err != nil {
		log.Printf("%v error : %v", name, err)
		mes := fmt.Sprintf("%v error : %v", name, err)
		SendMessageToServerChan("程序终止, 请查看报错情况 !", mes)
		os.Exit(1)
	}
}
