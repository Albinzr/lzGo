package lz

import (
	"bytes"
	"compress/lzw"
	"fmt"
	"testing"
)



func Test(t *testing.T){



	var data = []byte(getData())
	var buf bytes.Buffer

	com := lzw.NewWriter(&buf, lzw.LSB, 8)
	com.Write(data)

	fmt.Println("w->>>>*************************",buf.Len())

	com.Close()

	var output = make([]byte, len(data))

	dec := lzw.NewReader(&buf, lzw.LSB, 8)

	r, err := dec.Read(output)

	if err != nil {
		fmt.Println("read error:", err)
	}

	fmt.Println("read", r, "bytes")

	fmt.Printf("output: %#v\n", string(output[:r]))
}

func getData() string {
	return ""
}


