package main

import (
	"fmt"
)

func transformAscii(asciiDec int32, offset int32, bottom int32, top int32) int32 {

	for offset > 25 {
		offset = (offset % 25) - (offset / 25)
	}

	if bottom <= asciiDec && asciiDec <= top {
		if top < asciiDec+offset {
			return bottom + ((asciiDec + offset) - top - 1)
		} else {
			return asciiDec + offset
		}
	}
	return asciiDec
}

func caesarCipher(s string, k int32) string {

	cipher := make([]byte, len(s))
	//65-90, 97-122

	for i, c := range s {
		asciiDec := int32(c)
		transDec := int32(c)
		if 65 <= asciiDec && asciiDec <= 90 {
			transDec = transformAscii(asciiDec, k, 65, 90)
		} else if 97 <= asciiDec && asciiDec <= 122 {
			transDec = transformAscii(asciiDec, k, 97, 122)
		}
		cipher[i] = byte(transDec)
	}

	return string(cipher)
}

func main() {
	s := "www.abc.xy"
	fmt.Println(s)
	fmt.Println(caesarCipher(s, 87))
}
