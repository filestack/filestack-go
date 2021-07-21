package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/filestack/filestack-go/client"
	"github.com/filestack/filestack-go/resource"
	"github.com/filestack/filestack-go/transformation/args"
)

const yourHandle = ""
const yourApiKey = ""
const pathToYourFile = ""

func validate() {
	if yourHandle == "" {
		log.Fatal("Please set `yourHandle` const.")
	}
	if yourApiKey == "" {
		log.Fatal("Please set `yourApiKey` const.")
	}
	if pathToYourFile == "" {
		log.Fatal("Please set `pathToYourFile` const.")
	}
}

func main() {
	validate()

	cli, err := client.NewClient(yourApiKey)
	if err != nil {
		log.Fatalf("failed to initialize the client: %v", err)
	}

	tr := cli.MustNewTransformation(resource.NewHandle(yourHandle))
	tr.BlackWhite(args.NewBlackWhiteArgs())
	bytes, err := cli.ExecuteTransformation(context.Background(), tr)
	if err != nil {
		log.Fatalf("failed to execute the transformation: %s", err.Error())
	}

	err = ioutil.WriteFile(pathToYourFile, bytes, 0644)
	if err != nil {
		log.Fatalf("failed to write file into: %s", pathToYourFile)
	}

	fmt.Println("file was saved")
}
