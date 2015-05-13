package main

import (
	"io/ioutil"
	"strings"

	"./utils"
)

func parseFileToStringArray(filename string) (contentArray []string) {

	contents, err := ioutil.ReadFile(filename)
	utils.CheckError(err)

	contentArray = strings.Split(string(contents), "\n")
	return

}
