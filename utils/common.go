package utils

import (
	"io/ioutil"
	"log"
	"strings"
)

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
