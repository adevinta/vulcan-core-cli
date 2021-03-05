package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("Checks", func() {
	BasePath("/checks")

	Action("show", func() {
		Routing(GET("/:id"))
		Params(func() {
			Param("id", UUID)
		})
		Description("Get a Check by its ID")
		Response(OK, CheckMediaData)
		Response(NotFound)
		Response(InternalServerError)
		Response(BadRequest)
	})

})

var CheckMediaData = MediaType("application/vnd.checkdata+json", func() {
	Attributes(func() {
		Attribute("id", UUID)
		Attribute("status", String)
		Attribute("target", String)
		Attribute("checktype_name", String)
		Attribute("image", String)
		Attribute("options", String)
		Attribute("report", String)
		Attribute("raw", String)
		Attribute("tag", String)
		Attribute("assettype", String)
		Required("id", "checktype_name", "target", "status")
	})

	View("default", func() {
		Attribute("id")
		Attribute("status")
		Attribute("target")
		Attribute("checktype_name")
		Attribute("image")
		Attribute("options")
		Attribute("report")
		Attribute("raw")
		Attribute("tag")
		Attribute("assettype")
	})
})
