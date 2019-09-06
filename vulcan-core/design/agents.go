package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("Agents", func() {
	BasePath("/agents")

	Action("index", func() {
		Routing(GET(""))
		Description("Get all agents")
		Params(func() {
			Param("status")
		})
		Response(OK, AgentMediaCollection)
		Response(InternalServerError)
		Response(BadRequest)
	})

	Action("show", func() {
		Routing(GET("/:id"))
		Params(func() {
			Param("id", UUID)
		})
		Description("Get an Agent by its ID")
		Response(OK, AgentMedia)
		Response(NotFound)
		Response(InternalServerError)
		Response(BadRequest)
	})

	Action("pause", func() {
		Routing(POST("/:id/pause"))
		Description("Pause an agent")
		Params(func() {
			Param("id", UUID)
		})
		Response(OK)
		Response(InternalServerError)
		Response(BadRequest)
	})
})

var AgentData = Type("AgentData", func() {
	Attribute("id", UUID)
	Attribute("status", String, func() {
		MinLength(1)
		Example("RUNNING")
		Pattern("^[[:word:]]+")
	})
	Required("id", "status")
})

var AgentMediaData = MediaType("application/vnd.agentdata+json", func() {
	Reference(AgentData)

	Attributes(func() {
		Attribute("id", UUID)
		Attribute("status")
		Attribute("version")
		Attribute("enabled", Boolean)
		Attribute("heartbet_at", DateTime)
		Attribute("jobqueue", JobqueueMedia)
		Required("id", "status", "version", "enabled", "heartbet_at", "jobqueue")
	})

	View("default", func() {
		Attribute("id")
		Attribute("status")
		Attribute("version")
		Attribute("enabled")
		Attribute("heartbet_at")
		Attribute("jobqueue")
	})
})

var AgentMedia = MediaType("application/vnd.agent+json", func() {
	Attributes(func() {
		Attribute("agent", "application/vnd.agentdata+json")
		Required("agent")
	})

	View("default", func() {
		Attribute("agent")
	})
})

var AgentMediaCollection = MediaType("application/vnd.agents+json", func() {
	Attributes(func() {
		Attribute("agents", ArrayOf(AgentData))
		Required("agents")
	})

	View("default", func() {
		Attribute("agents")
	})
})

var JobqueueMedia = MediaType("application/vnd.jobqueue+json", func() {
	Attributes(func() {
		Attribute("jobqueue", AgentJobqueueMediaData)
		Required("jobqueue")
	})

	View("default", func() {
		Attribute("jobqueue")
	})
})

var AgentJobqueueMediaData = MediaType("application/vnd.agentjobqueuedata+json", func() {
	Reference(AgentData)

	Attributes(func() {
		Attribute("id", UUID)
		Attribute("arn")
		Attribute("description")
		Attribute("enabled", Boolean)
		Required("id", "arn", "description", "enabled")
	})

	View("default", func() {
		Attribute("id")
		Attribute("arn")
		Attribute("description")
		Attribute("enabled")
	})
})
