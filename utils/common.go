package utils

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func GetIntVal(str string) int {
	intVal, _ := strconv.Atoi(str)
	return intVal
}

func GetInt64Val(str string) int64 {
	intVal, _ := strconv.ParseInt(str, 10, 64)
	return intVal
}

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

func ParseFloatByReplace(targetStr, replaceStr string) float32 {
	floatVal, _ := strconv.ParseFloat(strings.Replace(strings.TrimSpace(targetStr), replaceStr, "", -1), 32)
	return float32(floatVal)
}

func ParseIntByReplace(targetStr, replaceStr string) int {
	intVal, _ := strconv.Atoi(strings.Replace(strings.TrimSpace(targetStr), replaceStr, "", -1))
	return intVal
}

func CheckError(err error) {
	if err != nil {
		log.Println("EEEE - ", err)
	}
}
