package utils

import (
	"log"
	"testing"
)

func TestParseToStringArray(t *testing.T) {

	contentArr := ParseFileToStringArray("tmp.txt")
	for _, content := range contentArr {
		log.Println(content)
	}
	t.Fail()

}
