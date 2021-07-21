package main

import (
	"fmt"
	"log"

	"github.com/filestack/filestack-go/resource"
	"github.com/filestack/filestack-go/transformation"
	"github.com/filestack/filestack-go/transformation/args"
)

const yourHandle = ""
const yourApiKey = ""

func validate() {
	if yourHandle == "" {
		log.Fatal("Please set `yourHandle` const.")
	}
	if yourApiKey == "" {
		log.Fatal("Please set `yourApiKey` const.")
	}
}

func main() {
	validate()

	tr := transformation.MustNewTransformation(resource.NewHandle(yourHandle), yourApiKey)
	url := tr.BlackWhite(args.NewBlackWhiteArgs()).BuildURL()
	fmt.Println(url)
}
