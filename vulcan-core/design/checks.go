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
		Response(OK, CheckMedia)
		Response(NotFound)
		Response(InternalServerError)
		Response(BadRequest)
	})

	Action("create", func() {
		Routing(POST("/"))
		Description("Create a new Check")
		Payload(CheckPayload)
		Response(Created)
		Response(InternalServerError)
		Response(BadRequest)
	})

	Action("abort", func() {
		Routing(POST("/:id/abort"))
		Params(func() {
			Param("id", UUID)
		})
		Description("Abort a Check gracefully")
		Response(NoContent)
		Response(NotFound)
		Response(InternalServerError)
		Response(BadRequest)
	})

	Action("kill", func() {
		Routing(POST("/:id/kill"))
		Params(func() {
			Param("id", UUID)
		})
		Description("Kill a Check immediately")
		Response(NoContent)
		Response(NotFound)
		Response(InternalServerError)
		Response(BadRequest)
	})
})

var CheckData = Type("CheckData", func() {
	Attribute("checktype_name", String, func() {
		Description("Name of the checktype this check is")
	})
	Attribute("target", String, func() {
		MinLength(1)
		Example(`www.example.com`)
		Pattern("^[[:print:]]+")
		Description("Target to be scanned. Can be a domain, hostname, IP or URL")
	})
	Attribute("options", String, func() {
		MinLength(2)
		Example(`{ "policy" : "fullscan" }`)
		Pattern("^[[:print:]]+")
		Description("Configuration options for the Check. It should be in JSON format")
	})
	Attribute("webhook", String, func() {
		MinLength(3)
		Format("uri")
		Example("https://webhook.example.com/callback")
		Description("Webhook to callback after the Check has finished")
	})

	Attribute("tag", String, func() {
		MinLength(3)
		Example("sdrn")
		Description("a tag")
	})

	Attribute("jobqueue_id", UUID, func() {
		Description(`ID of the specific queue this check must we enqueued.
		The queue must already be created in vulcan core.`)
	})

	Attribute("assettype", String, func() {
		MinLength(1)
		Example(`Hostname`)
		Pattern("^[[:print:]]+")
		Description("Asset type of the target. Can be a DomainName, Hostname, etc.")
	})

	Required("target")
})

var CheckPayload = Type("CheckPayload", func() {
	Attribute("check", "CheckData")
	Required("check")
})

// BUG: the type was "application/adevinta.vulcan.api.data+json"
// but we had to change it because Gorma was not correctly generating some
// function calls. We should investigate more about it.
var CheckMediaData = MediaType("application/vnd.checkdata+json", func() {
	Reference(CheckData)

	Attributes(func() {
		Attribute("id", UUID)
		Attribute("checktype", ChecktypeType)
		Attribute("scan", ScanMediaData)
		Attribute("target")
		Attribute("status", String)
		Attribute("options")
		Attribute("webhook")
		Attribute("progress", Number)
		Attribute("raw", String)
		Attribute("report", String)
		Attribute("assettype", String)
		Required("id", "checktype", "target", "status")
	})

	View("default", func() {
		Attribute("id")
		Attribute("checktype")
		Attribute("target")
		Attribute("status")
		Attribute("options")
		Attribute("webhook")
		Attribute("progress")
		Attribute("raw")
		Attribute("report")
		Attribute("scan")
		Attribute("assettype")
	})
})

// BUG: the type was "application/adevinta.vulcan.api.check+json"
// but we had to change it because Gorma was not correctly generating some
// function calls. We should investigate more about it.
var CheckMedia = MediaType("application/vnd.check+json", func() {
	Attributes(func() {
		Attribute("check", "application/vnd.checkdata+json")
		Required("check")
	})

	View("default", func() {
		Attribute("check")
	})
})
