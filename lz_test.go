package lz

import (
	"fmt"
	"testing"
)

func Test(t *testing.T){
	string := "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"
	//for {
		jsonString, err := DecompressFromBase64(string)
		fmt.Println(jsonString, err)
	//}
}