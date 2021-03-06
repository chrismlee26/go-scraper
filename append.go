package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func AppendFile() {
	file, err := os.OpenFile("test.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	defer file.Close()

	len, err := file.WriteString("Text to be appended to the file\n")
	if err != nil {
		log.Fatalf("Failed writing to file: %s", err)
	}
	fmt.Printf("\nLength %d bytes", len)
	fmt.Printf("\nFile contents: %s", file.Name())
}

func ReadFile() {
	data, err := ioutil.ReadFile("test.json")
	if err != nil {
		log.Panicf("failed opening file: %s", err)
	}
	fmt.Printf("\nLength: %d bytes:", len(data))
	fmt.Printf("\nData: %s", data)
	fmt.Printf("\nError: %v", err)
}

func main() {
	fmt.Printf("####### Append File #######\n")
	AppendFile()
	fmt.Printf("\n######## Read File #######\n")
	ReadFile()
}
