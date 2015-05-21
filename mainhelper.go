package main

import (
	"io/ioutil"
	"strings"

	"./utils"
)

const fundcodeAllFilename = "fund_all.txt"
const fundcodeFilename = "fund_choose.txt"
const stockcodeAllFilename = "stock_a.txt"
const stockcodeFilename = "stock_choose.txt"

const typestock = 1
const typefund = 2

func parseFileToStringArray(parseType int) (contentArray []string) {

	var filename string
	if parseType == 1 {
		filename = stockcodeFilename
	} else {
		filename = fundcodeFilename
	}
	contents, err := ioutil.ReadFile(filename)
	utils.CheckError(err)

	contentArray = strings.Split(string(contents), "\n")
	contentArray = contentArray[:len(contentArray)-1]
	return

}

func parseFileToStringArrayAll(parseType int) (contentArray []string) {

	var filename string
	if parseType == 1 {
		filename = stockcodeAllFilename
	} else {
		filename = fundcodeAllFilename
	}
	contents, err := ioutil.ReadFile(filename)
	utils.CheckError(err)

	contentArray = strings.Split(string(contents), "\n")
	contentArray = contentArray[:len(contentArray)-1]
	return

}
