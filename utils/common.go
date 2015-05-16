package utils

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func ParseCmdFlagToDateArray() (hasCmdFlag bool, tranDateArray []string) {

	hasCmdFlag = false
	tranDateArray = make([]string, 0)

	sizeArgs := len(os.Args)
	if sizeArgs == 1 {
		return
	}

	hasCmdFlag = true
	tranDateArray = GetDateArray(ParseDateWithIntFmt(os.Args[1]))
	return

}

func ParseFileToStringArray(filename string) (contentArray []string) {

	contents, err := ioutil.ReadFile(filename)
	CheckError(err)

	contentArray = strings.Split(string(contents), "\n")
	return

}

func CheckError(err error) {
	if err != nil {
		log.Println("EEEE - ", err)
	}
}
