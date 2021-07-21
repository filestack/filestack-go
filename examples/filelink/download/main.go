package main

import (
	"fmt"
	"log"
	"os"

	"github.com/filestack/filestack-go/client"
)

const yourApiKey = ""
const yourHandle = ""
const yourLocalPath = ""

func validate() {
	if yourApiKey == "" {
		log.Fatal("Please set `yourApiKey`.")
	}
	if yourHandle == "" {
		log.Fatal("Please set `yourHandle`.")
	}
	if yourLocalPath == "" {
		log.Fatal("Please set `yourLocalPath`.")
	}
}

func main() {
	validate()

	cli, err := client.NewClient(yourApiKey)
	if err != nil {
		log.Fatalf("failed to initialize client: %v", err)
	}
	fileLink := cli.MustNewFileLink(yourHandle)

	filePath := yourLocalPath
	file, fileErr := os.Create(filePath)
	if fileErr != nil {
		log.Fatalf("could not create a local file in: %v", filePath)
	}
	defer file.Close()

	bytes, err := fileLink.Download(file)
	if err != nil {
		log.Fatalf("failed to download file: %v", filePath)
	}

	fmt.Printf("successfully downloaded file (size=%v)", bytes)
}
