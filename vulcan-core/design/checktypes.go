package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("Checktypes", func() {
	BasePath("/checktypes")

	Action("index", func() {
		Routing(GET(""))
		Params(func() {
			Param("enabled", String, func() {
				Default("true")
			})
		})
		Description("Get all checktypes")
		Response(OK, ChecktypeMediaCollection)
		Response(InternalServerError)
		Response(BadRequest)
	})

	Action("show", func() {
		Routing(GET("/:id"))
		Params(func() {
			Param("id", UUID)
		})
		Description("Get a Checktype by its ID")
		Response(OK, ChecktypeMedia)
		Response(NotFound)
		Response(InternalServerError)
		Response(BadRequest)
	})
})

var ChecktypeType = Type("ChecktypeType", func() {
	Attribute("id", UUID)
	Attribute("name", String, func() {
		MinLength(1)
		Example("nmap")
		Pattern("^[[:word:]]+")
	})
	Attribute("description", String, func() {
		MaxLength(255)
		Pattern("^[[:print:]]+")
	})
	Attribute("options", String, "Default configuration options for the Checktype. It should be in JSON format")
	Attribute("enabled", Boolean)
	Attribute("timeout", Integer, "Specifies the maximum amount of time that the check should be running before it's killed")
	Attribute("image", String, "Image that needs to be pulled to run the Check of this type")
	Required("id", "name", "image")
})

// BUG: the type was "application/adevinta.vulcan.api.checktype+json"
// but we had to change it because Gorma was not correctly generating some
// function calls. We should investigate more about it.
var ChecktypeMedia = MediaType("application/vnd.checktype+json", func() {
	Attributes(func() {
		Attribute("checktype", ChecktypeType)
		Required("checktype")
	})

	View("default", func() {
		Attribute("checktype")
	})
})

// BUG: the type was "application/adevinta.vulcan.api.checktypes+json"
// but we had to change it because Gorma was not correctly generating some
// function calls. We should investigate more about it.
var ChecktypeMediaCollection = MediaType("application/vnd.checktypes+json", func() {
	Attributes(func() {
		Attribute("checktypes", ArrayOf(ChecktypeType))
		Required("checktypes")
	})

	View("default", func() {
		Attribute("checktypes")
	})
})
