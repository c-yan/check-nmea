package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func calcCheckSum(s string) int {
	result := 0
	for _, b := range []byte(s) {
		result ^= int(b)
	}
	return result
}

func isBrokenLine(s string) bool {
	if s[0] != '$' {
		return true
	}
	if len(s)-3 < 0 || s[len(s)-3] != '*' {
		return true
	}
	return fmt.Sprintf("%02X", calcCheckSum(s[1:len(s)-3])) != s[len(s)-2:]
}

func checkFileContent(fileName string) int {
	result := 0

	r, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("%s を開くのに失敗しました.\n%v\n", fileName, err)
	}
	defer r.Close()

	reader := bufio.NewScanner(r)
	n := 0
	for reader.Scan() {
		l := reader.Text()
		if n == 0 && l[0] == '/' {
			continue
		}
		if isBrokenLine(l) {
			result++
		}
		n++
	}

	return result
}

func main() {
	log.SetFlags(log.Lshortfile)
	for i := 1; i < len(os.Args); i++ {
		errors := checkFileContent(os.Args[i])
		if errors != 0 {
			fmt.Printf("%s has %d errors.\n", os.Args[i], errors)
		}
	}
}
