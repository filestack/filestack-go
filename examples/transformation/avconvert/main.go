package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/filestack/filestack-go/client"
	clientOptions "github.com/filestack/filestack-go/options/client"
	transformationOptions "github.com/filestack/filestack-go/options/transformation"
	"github.com/filestack/filestack-go/resource"
	"github.com/filestack/filestack-go/security"
	"github.com/filestack/filestack-go/transformation/args"
)

const yourApiKey = ""
const yourSecret = ""
const yourHandle = ""

func validate() {
	if yourApiKey == "" {
		log.Fatal("Please set `yourApiKey` const.")
	}
	if yourSecret == "" {
		log.Fatal("Please set `yourSecret` const.")
	}
	if yourHandle == "" {
		log.Fatal("Please set `yourHandle` const.")
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

	av := cli.MustNewAudioVisual(
		resource.NewHandle(yourHandle),
		args.AVConvertOptions{Width: 320, Height: 240},
		transformationOptions.SecurityPolicy(security),
	)

	completes := false
	timeLimitSec := 60
	ctx := context.Background()
	for i := 1; i < timeLimitSec; i++ {
		status, err := av.Status(ctx)
		if err != nil {
			log.Fatalf("failed with error: %s", err.Error())
		}
		fmt.Printf("current status: %s\n", status)
		if status == "completed" {
			completes = true
			break
		}
		time.Sleep(time.Second)
	}
	if !completes {
		log.Fatalf("conversion was not completed in %v seconds", timeLimitSec)
	}

	fileLink, err := av.ToFileLink(ctx)
	if err != nil {
		log.Fatalf("failed with error: %s", err.Error())
	}
	log.Fatalf("fileLink: %s", fileLink.AsString())
}
