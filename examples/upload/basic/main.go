package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/filestack/filestack-go/client"
)

const yourApiKey = ""
const yourFilePath = ""

func validate() {
	if yourApiKey == "" {
		log.Fatal("Please set `yourApiKey` const.")
	}
	if yourFilePath == "" {
		log.Fatal("Please set `yourFilePath` const.")
	}
}

func main() {
	validate()

	cli, err := client.NewClient(yourApiKey)
	if err != nil {
		log.Fatalf("failed to initialize the client: %v", err)
	}

	file, err := os.Open(yourFilePath)
	if err != nil {
		return
	}
	defer file.Close()

	fileLink, err := cli.Upload(context.Background(), file)
	if err != nil {
		log.Fatalf("failed to upload file: %v", err)
	}

	fmt.Println(fileLink.GetHandle().AsString())
}
