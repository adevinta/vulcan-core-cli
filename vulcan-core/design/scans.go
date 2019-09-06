package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("Scans", func() {
	BasePath("/scans")

	Action("index", func() {
		Routing(GET(""))
		Description("Get all scans")
		Response(OK, ScanMediaCollection)
		Response(InternalServerError)
		Response(BadRequest)
	})

	Action("show", func() {
		Routing(GET("/:id"))
		Params(func() {
			Param("id", UUID)
		})
		Description("Get a Scan by its ID")
		Response(OK, ScanMedia)
		Response(NotFound)
		Response(InternalServerError)
		Response(BadRequest)
	})

	Action("create", func() {
		Routing(POST("/"))
		Description("Create a new Scan")
		Payload(ScanPayload)
		Response(Created)
		Response(InternalServerError)
		Response(BadRequest)
	})

	Action("abort", func() {
		Routing(POST("/:id/abort"))
		Description("Abort a scan")
		Params(func() {
			Param("id", UUID)
		})
		Response(OK)
		Response(InternalServerError)
		Response(BadRequest)
	})

	Action("checks", func() {
		Routing(GET("/:id/checks"))
		Params(func() {
			Param("id", UUID)
		})
		Description("Get checks of the Scan")
		Response(OK, ChecksMediaCollection)
		Response(NotFound)
		Response(InternalServerError)
		Response(BadRequest)
	})

	Action("stats", func() {
		Routing(GET("/:id/stats"))
		Params(func() {
			Param("id", UUID)
		})
		Description("Get stats of the Scan")
		Response(OK, StatsMediaCollection)
		Response(NotFound)
		Response(InternalServerError)
		Response(BadRequest)
	})
})

var _ = Resource("FileScans", func() {
	BasePath("/filescan")
	Action("upload", func() {
		Routing(POST("/"))
		Payload(FileScanPayload)
		MultipartForm()
		Description("Create a scan by uploading a file using a multipart form with the scan definition")
		Response(Created, ScanMediaData)
		Response(InternalServerError)
		Response(BadRequest)
	})

})

var FileScanPayload = Type("FileScanPayload", func() {
	Attribute("upload", File, "Upload")
	Required("upload")
})

var ScanData = Type("ScanData", func() {
	Attribute("size", Integer, func() {
		Description("Number of checks of the scan")
	})
	Required("size")
})

var ScanPayload = Type("ScanPayload", func() {
	Attribute("scan", ScanChecksPayload)
})

var ScanChecksPayload = Type("ScanChecksPayload", func() {
	Attribute("checks", ArrayOf(CheckPayload))
	Required("checks")
})

// BUG: the type was "application/adevinta.vulcan.api.data+json"
// but we had to change it because Gorma was not correctly generating some
// function calls. We should investigate more about it.
var ScanMediaData = MediaType("application/vnd.scandata+json", func() {
	Reference(ScanData)

	Attributes(func() {
		Attribute("id", UUID)
		Attribute("size")
		Attribute("created_at", DateTime)
		Required("id", "size", "created_at")
	})

	View("default", func() {
		Attribute("id")
		Attribute("size")
	})
})

// BUG: the type was "application/adevinta.vulcan.api.check+json"
// but we had to change it because Gorma was not correctly generating some
// function calls. We should investigate more about it.
var ScanMedia = MediaType("application/vnd.scan+json", func() {
	Attributes(func() {
		Attribute("scan", "application/vnd.scandata+json")
		Required("scan")
	})

	View("default", func() {
		Attribute("scan")
	})
})

// BUG: the type was "application/adevinta.vulcan.api.scan+json"
// but we had to change it because Gorma was not correctly generating some
// function calls. We should investigate more about it.
var ScanMediaCollection = MediaType("application/vnd.scans+json", func() {
	Attributes(func() {
		Attribute("scans", ArrayOf(ScanData))
		Required("scans")
	})

	View("default", func() {
		Attribute("scans")
	})
})

// BUG: the type was "application/adevinta.vulcan.api.checks+json"
// but we had to change it because Gorma was not correctly generating some
// function calls. We should investigate more about it.
var ChecksMediaCollection = MediaType("application/vnd.checks+json", func() {
	Attributes(func() {
		Attribute("checks", ArrayOf(ScanCheckMediaData))
		Required("checks")
	})

	View("default", func() {
		Attribute("checks")
	})
})

// BUG: the type was "application/adevinta.vulcan.api.data+json"
// but we had to change it because Gorma was not correctly generating some
// function calls. We should investigate more about it.
var ScanCheckMediaData = MediaType("application/vnd.scancheckdata+json", func() {
	Reference(CheckData)

	Attributes(func() {
		Attribute("id", UUID)
		Attribute("checktype_name")
		Attribute("target")
		Attribute("status", String)
		Required("id", "checktype_name", "target", "status")
	})

	View("default", func() {
		Attribute("id")
		Attribute("checktype_name")
		Attribute("target")
		Attribute("status")
	})
})

// BUG: the type was "application/adevinta.vulcan.api.stats+json"
// but we had to change it because Gorma was not correctly generating some
// function calls. We should investigate more about it.
var StatsMediaCollection = MediaType("application/vnd.stats+json", func() {
	Attributes(func() {
		Attribute("checks", ArrayOf(Stat))
		Required("checks")
	})

	View("default", func() {
		Attribute("checks")
	})
})

// BUG: the type was "application/adevinta.vulcan.api.stat+json"
// but we had to change it because Gorma was not correctly generating some
// function calls. We should investigate more about it.
var Stat = MediaType("application/vnd.stat+json", func() {
	Attributes(func() {
		Attribute("status", String)
		Attribute("total", Integer)
		Required("status", "total")
	})

	View("default", func() {
		Attribute("status")
		Attribute("total")
	})
})
