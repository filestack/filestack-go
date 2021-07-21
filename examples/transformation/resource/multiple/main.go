package main

import (
	"fmt"
	"log"

	"github.com/filestack/filestack-go/resource"
	"github.com/filestack/filestack-go/transformation"
)

const yourApiKey = ""
const yourHandle1 = ""
const yourHandle2 = ""
const yourURL = ""

func validate() {
	if yourApiKey == "" {
		log.Fatal("Please set `yourApiKey` const.")
	}
	if yourHandle1 == "" {
		log.Fatal("Please set `yourHandle1` const.")
	}
	if yourHandle2 == "" {
		log.Fatal("Please set `yourHandle2` const.")
	}
	if yourURL == "" {
		log.Fatal("Please set `yourURL` const.")
	}
}

func main() {
	validate()

	tr := transformation.MustNewTransformationMultiResource(
		[]resource.Resource{
			resource.NewHandle(yourHandle1),
			resource.NewHandle(yourHandle2),
			resource.NewExternalURL(yourURL),
		},
		yourApiKey,
	)
	zipTransformation := tr.Zip()
	url := zipTransformation.BuildURL()
	fmt.Println(url)
}
