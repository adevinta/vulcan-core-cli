package main

import (
	_ "github.com/adevinta/vulcan-core-cli/vulcan-core/design"

	"github.com/goadesign/goa/design"
	"github.com/goadesign/goa/goagen/codegen"
	genclient "github.com/goadesign/goa/goagen/gen_client"
	genswagger "github.com/goadesign/goa/goagen/gen_swagger"
)

func main() {
	codegen.ParseDSL()
	codegen.Run(
		genswagger.NewGenerator(
			genswagger.API(design.Design),
			genswagger.OutDir("_resources"),
		),
		genclient.NewGenerator(
			genclient.API(design.Design),
		),
	)
}
