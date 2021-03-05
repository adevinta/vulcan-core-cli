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
		Params(func() {
			Param("offset", Integer)
			Param("limit", Integer)
			Param("external_id", String)
		})
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
		Response(OK, ScanMediaData)
		Response(NotFound)
		Response(InternalServerError)
		Response(BadRequest)
	})

	Action("create", func() {
		Routing(POST("/"))
		Description("Create a new Scan")
		Payload(ScanPayload)
		Response(Created, CreateScanMediaData)
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
		Response(Conflict)
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

var ScanPayload = Type("ScanPayload", func() {
	Attribute("external_id", UUID)
	Attribute("scheduled_time", DateTime)
	Attribute("trigger", String)
	Attribute("tag", String)
	Attribute("target_groups", ArrayOf(ScanTargetsGroup))
})

var ScanTargetsGroup = Type("ScanTargetsGroup", func() {
	Attribute("target_group", TargetGroup)
	Attribute("checktypes_group", ChecktypesGroup)
})

var TargetGroup = Type("TargetGroup", func() {
	Attribute("name", String)
	Attribute("options", String)
	Attribute("targets", ArrayOf(Target))
})

var Target = Type("Target", func() {
	Attribute("identifier", String)
	Attribute("type", String)
	Attribute("options", String)

})

var ChecktypesGroup = Type("ChecktypesGroup", func() {
	Attribute("name", String)
	Attribute("checktypes", ArrayOf(ScanChecktype))
})

var ScanChecktype = Type("ScanChecktype", func() {
	Attribute("name", String)
	Attribute("options", String)
})

var CreateScanMediaData = MediaType("application/vnd.createscandata+json", func() {
	Attribute("scan_id", UUID)
	View("default", func() {
		Attribute("scan_id")
	})
})

var ScanMediaData = MediaType("application/vnd.scandata+json", func() {
	Attributes(func() {
		Attribute("id", UUID)
		Attribute("external_id", UUID)
		Attribute("status", String)
		Attribute("trigger", String)
		Attribute("scheduled_time", DateTime)
		Attribute("start_time", DateTime)
		Attribute("end_time", DateTime)
		Attribute("aborted_at", DateTime)
		Attribute("progress", Number)
		Attribute("check_count", Integer)
		Attribute("checks_created", Integer)
		Required("id", "status")
	})

	View("default", func() {
		Attribute("id")
		Attribute("status")
		Attribute("check_count", Integer)
	})
})

var ScanMediaCollection = MediaType("application/vnd.scans+json", func() {
	Attributes(func() {
		Attribute("scans", ArrayOf(ScanMediaData))
		Required("scans")
	})

	View("default", func() {
		Attribute("scans")
	})
})

var ChecksMediaCollection = MediaType("application/vnd.checks+json", func() {
	Attributes(func() {
		Attribute("checks", ArrayOf(CheckMediaData))
		Required("checks")
	})

	View("default", func() {
		Attribute("checks")
	})
})

var StatsMediaCollection = MediaType("application/vnd.stats+json", func() {
	Attributes(func() {
		Attribute("checks", ArrayOf(Stat))
		Required("checks")
	})

	View("default", func() {
		Attribute("checks")
	})
})

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
