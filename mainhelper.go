package main

import (
	"io/ioutil"
	"strings"

	"./utils"
)

const stockcodeAllFilename = "stock_a.txt"
const stockcodeFilename = "stock_choose.txt"

func parseFileToStringArray() (contentArray []string) {

	contents, err := ioutil.ReadFile(stockcodeFilename)
	utils.CheckError(err)

	contentArray = strings.Split(string(contents), "\n")
	contentArray = contentArray[:len(contentArray)-1]
	return

}

func parseFileToStringArrayAll() (contentArray []string) {

	contents, err := ioutil.ReadFile(stockcodeAllFilename)
	utils.CheckError(err)

	contentArray = strings.Split(string(contents), "\n")
	contentArray = contentArray[:len(contentArray)-1]
	return

}
