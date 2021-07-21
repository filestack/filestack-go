package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/filestack/filestack-go/client"
	clientOptions "github.com/filestack/filestack-go/options/client"
	"github.com/filestack/filestack-go/resource"
	"github.com/filestack/filestack-go/security"
	"github.com/filestack/filestack-go/transformation/args"
)

const yourApiKey = ""
const yourSecret = ""
const yourExternalURL = ""

func validate() {
	if yourApiKey == "" {
		log.Fatal("Please set `yourApiKey` const.")
	}
	if yourSecret == "" {
		log.Fatal("Please set `yourSecret` const.")
	}
	if yourExternalURL == "" {
		log.Fatal("Please set `yourExternalURL` const.")
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
	tr := cli.MustNewTransformation(resource.NewExternalURL(yourExternalURL))
	tr.Vignette(args.NewVignetteArgs())

	ctx := context.Background()
	storeResponse, err := cli.Store(ctx, tr, args.NewStoreArgs())
	if err != nil {
		log.Fatalf("failed to store file: %v", err)
	}
	fileLink := cli.MustNewFileLink(storeResponse.Handle)

	err = fileLink.Delete(ctx)
	if err != nil {
		log.Fatalf("failed to delete the file: %v", err)
	}

	fmt.Println("delete was successful")
}
