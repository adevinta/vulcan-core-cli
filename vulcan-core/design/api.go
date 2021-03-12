/*
Copyright 2021 Adevinta
*/

package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("vulcan-core", func() {
	Title("Vulcan Core")
	Description("Vulnerability Scanning Framework API")
	Version("1.0")
	BasePath("/v1")
	Scheme("http")
	Host("localhost:3000")
	Consumes("application/json")
	ResponseTemplate(InternalServerError, func() {
		Description("Server error")
		Status(500)
		Media(ErrorMedia)
	})
	ResponseTemplate(Created, func() {
		Description("Resource created")
		Status(201)
		Media("application/vnd.check+json")
	})
})

var _ = Resource("swagger", func() {
	Origin("*", func() {
		Methods("GET") // Allow all origins to retrieve the Swagger JSON (CORS)
	})
	Files("/v1/swagger.json", "./_resources/swagger/swagger.json")
})
