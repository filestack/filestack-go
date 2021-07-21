package main

import (
	"fmt"
	"log"

	"github.com/filestack/filestack-go/resource"
	"github.com/filestack/filestack-go/transformation"
	"github.com/filestack/filestack-go/transformation/args"
)

const storageAlias = ""
const pathToYourFile = ""
const yourApiKey = ""

func validate() {
	if storageAlias == "" {
		log.Fatal("Please set `storageAlias` const.")
	}
	if pathToYourFile == "" {
		log.Fatal("Please set `pathToYourFile` const.")
	}
	if yourApiKey == "" {
		log.Fatal("Please set `yourApiKey` const.")
	}
}

func main() {
	validate()

	tr := transformation.MustNewTransformation(resource.NewStorageAlias(storageAlias, pathToYourFile), yourApiKey)
	url := tr.BlackWhite(args.NewBlackWhiteArgs()).BuildURL()
	fmt.Println(url)
}
