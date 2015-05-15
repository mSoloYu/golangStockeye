package main

import (
	"io/ioutil"
	"strings"

	"./utils"
)

const stockcodeFilename = "stock_choose.txt"

func parseFileToStringArray() (contentArray []string) {

	contents, err := ioutil.ReadFile(stockcodeFilename)
	utils.CheckError(err)

	contentArray = strings.Split(string(contents), "\n")
	return

}
