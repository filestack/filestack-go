package main

import (
	"fmt"
	"log"

	"github.com/filestack/filestack-go/resource"
	"github.com/filestack/filestack-go/transformation"
	"github.com/filestack/filestack-go/transformation/args"
)

const yourHandle = ""

func validate() {
	if yourHandle == "" {
		log.Fatal("Please set `yourHandle` const.")
	}
}

func main() {
	validate()

	tr := transformation.MustNewTransformation(resource.NewHandle(yourHandle), "")
	tr.Resize(
		args.NewResizeArgs().
			SetWidth(240).
			SetHeight(320).
			SetFit("scale").
			SetAlign("center"),
	).Rotate(
		args.NewRotateArgs().
			SetDegrees(180),
	).Negative()

	fmt.Println(tr.BuildURL())
}
