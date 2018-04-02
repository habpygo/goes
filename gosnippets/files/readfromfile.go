package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadLine(file string) {
	f, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	r := bufio.NewReaderSize(f, 4*1024)
	line, isPrefix, err := r.ReadLine()
	for err == nil && !isPrefix {
		s := string(line) // convert slice of bytes to the string
		fmt.Println(s)
		line, isPrefix, err = r.ReadLine()
	}
	if isPrefix {
		fmt.Println("buffersize too small")
		return
	}
	if err != io.EOF {
		fmt.Println(err)
		return
	}
}

func main() {
	fmt.Println("Import file in Go")

	ReadLine("bankledger.json")

}
