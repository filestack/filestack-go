package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/filestack/filestack-go/client"
	clientOptions "github.com/filestack/filestack-go/options/client"
	uploadOptions "github.com/filestack/filestack-go/options/upload"
	"github.com/filestack/filestack-go/security"
	"github.com/filestack/filestack-go/store"
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
		log.Fatal("Please set `yourFilePath` const.")
	}
}

func main() {
	validate()

	security := security.NewSecurity(yourSecret, &security.Policy{
		Expiry: time.Now().Add(time.Duration(24 * time.Hour)).Unix(), // the only required parameter
	})

	cli, err := client.NewClient(yourApiKey, clientOptions.SecurityPolicy(security))
	if err != nil {
		log.Fatalf("failed to initialize the client: %v", err)
	}

	file, err := os.Open(yourFilePath)
	if err != nil {
		return
	}
	defer file.Close()

	fileLink, err := cli.Upload(
		context.Background(),
		file,
		uploadOptions.StoreParams(&store.Params{
			FileName: "myvideo.mp4",
			MimeType: "video/mp4",
		}),
		uploadOptions.Intelligent(),
	)
	if err != nil {
		log.Fatalf("failed to upload file: %v", err)
	}

	fmt.Println(fileLink.GetHandle().AsString())
}
