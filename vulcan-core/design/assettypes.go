/*
Copyright 2021 Adevinta
*/

package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("Assettypes", func() {
	BasePath("/assettypes")

	Action("index", func() {
		Routing(GET(""))
		Description("Get all assettypes and per each one the checktypes that are accepting that concrete assettype")
		Response(OK, AssettypeMediaCollection)
		Response(InternalServerError)
		Response(BadRequest)
	})
})

var AssettypeType = Type("AssettypeType", func() {
	Attribute("assettype", String, func() {
		Example("DomainName")
	})
	Attribute("name", ArrayOf(String))
})

var AssettypeMediaCollection = CollectionOf(AssettypeMedia)
var AssettypeMedia = MediaType("application/vnd.assettype+json", func() {
	Reference(AssettypeType)
	Attributes(func() {
		Attribute("assettype")
		Attribute("name")
	})
	View("default", func() {
		Attribute("assettype")
		Attribute("name")
	})
})
