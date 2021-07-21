package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/filestack/filestack-go/client"
	clientOptions "github.com/filestack/filestack-go/options/client"
	uploadOptions "github.com/filestack/filestack-go/options/upload"
	"github.com/filestack/filestack-go/security"
)

const yourApiKey = ""
const yourSecret = ""
const yourFilePath = ""

func validate() {
	if yourApiKey == "" {
		log.Fatal("Please set `yourApiKey` const.")
	}
	if yourSecret == "" {
		log.Fatal("Please set `yourSecret` const.")
	}
	if yourFilePath == "" {
		log.Fatal("Plase set `yourFilePath` const.")
	}
}

func main() {
	validate()

	// my security policy
	securityPolicy := security.NewSecurity(yourSecret, &security.Policy{
		Expiry: time.Now().Add(time.Duration(24 * time.Hour)).Unix(), // the only required parameter
	})

	// my customized http client
	httpClient := http.Client{Timeout: 60 * time.Second}

	cli, err := client.NewClient(
		yourApiKey,
		clientOptions.SecurityPolicy(securityPolicy),
		clientOptions.HTTPClient(&httpClient),
		clientOptions.MaxConcurrentUploads(10), // overwritten max concurrent uploads
	)
	if err != nil {
		log.Fatalf("failed to get client with security")
	}

	file, err := os.Open(yourFilePath)
	if err != nil {
		return
	}
	defer file.Close()

	fileLink, err := cli.Upload(
		context.Background(),
		file,
		uploadOptions.Intelligent(), // passing this option will enable fii (intelligent upload)
	)

	if err != nil {
		log.Fatalf("failed with error %s", err.Error())
	}

	fmt.Printf("The file has been successfuly uploaded: %s", fileLink.GetHandle().AsString())
}
