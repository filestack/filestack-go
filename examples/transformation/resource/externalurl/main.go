package main

import (
	"fmt"
	"log"

	"github.com/filestack/filestack-go/resource"
	"github.com/filestack/filestack-go/transformation"
	"github.com/filestack/filestack-go/transformation/args"
)

const yourURL = ""
const yourApiKey = ""

func validate() {
	if yourURL == "" {
		log.Fatal("Please set `yourURL` const.")
	}
	if yourApiKey == "" {
		log.Fatal("Please set `yourApiKey` const.")
	}
}

func main() {
	validate()

	tr := transformation.MustNewTransformation(resource.NewExternalURL(yourURL), yourApiKey)
	url := tr.BlackWhite(args.NewBlackWhiteArgs()).BuildURL()
	fmt.Println(url)
}
