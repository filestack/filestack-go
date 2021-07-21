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
)

const yourApiKey = ""
const yourSecret = ""
const yourWorkflowID = ""
const yourExternalURL = ""

func validate() {
	if yourApiKey == "" {
		log.Fatal("Please set `yourApiKey` const.")
	}
	if yourSecret == "" {
		log.Fatal("Please set `yourSecret` const.")
	}
	if yourWorkflowID == "" {
		log.Fatal("Please set `yourWorkflowID` const.")
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
		log.Fatal("failed to initialize the client")
	}

	externalURL := resource.NewExternalURL(yourExternalURL)
	ctx := context.Background()
	runResponse, err := cli.RunWorkflow(ctx, yourWorkflowID, externalURL)
	if err != nil {
		log.Fatalf("failed to run workflow: %v", err)
	}

	for {
		statusResponse, err := cli.CheckWorkflowStatus(ctx, runResponse.JobID)
		if err != nil {
			log.Fatalf("failed to check workflow status: %v", err)
		}
		fmt.Println(statusResponse.Status)
		if statusResponse.Status == "Finished" {
			fmt.Println(statusResponse.Results)
			break
		}
		time.Sleep(time.Second * 5)
	}
}
