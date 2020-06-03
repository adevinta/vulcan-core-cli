// Code generated by goagen v1.4.3, DO NOT EDIT.
//
// API "vulcan-persistence": Application User Types
//
// Command:
// $ main

package client

import (
	"github.com/goadesign/goa"
	uuid "github.com/goadesign/goa/uuid"
	"unicode/utf8"
)

// agentData user type.
type agentData struct {
	ID     *uuid.UUID `form:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty" xml:"id,omitempty"`
	Status *string    `form:"status,omitempty" json:"status,omitempty" yaml:"status,omitempty" xml:"status,omitempty"`
}

// Validate validates the agentData type instance.
func (ut *agentData) Validate() (err error) {
	if ut.ID == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "id"))
	}
	if ut.Status == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "status"))
	}
	if ut.Status != nil {
		if ok := goa.ValidatePattern(`^[[:word:]]+`, *ut.Status); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`request.status`, *ut.Status, `^[[:word:]]+`))
		}
	}
	if ut.Status != nil {
		if utf8.RuneCountInString(*ut.Status) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`request.status`, *ut.Status, utf8.RuneCountInString(*ut.Status), 1, true))
		}
	}
	return
}

// Publicize creates AgentData from agentData
func (ut *agentData) Publicize() *AgentData {
	var pub AgentData
	if ut.ID != nil {
		pub.ID = *ut.ID
	}
	if ut.Status != nil {
		pub.Status = *ut.Status
	}
	return &pub
}

// AgentData user type.
type AgentData struct {
	ID     uuid.UUID `form:"id" json:"id" yaml:"id" xml:"id"`
	Status string    `form:"status" json:"status" yaml:"status" xml:"status"`
}

// Validate validates the AgentData type instance.
func (ut *AgentData) Validate() (err error) {

	if ut.Status == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "status"))
	}
	if ok := goa.ValidatePattern(`^[[:word:]]+`, ut.Status); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`type.status`, ut.Status, `^[[:word:]]+`))
	}
	if utf8.RuneCountInString(ut.Status) < 1 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`type.status`, ut.Status, utf8.RuneCountInString(ut.Status), 1, true))
	}
	return
}

// assettypeType user type.
type assettypeType struct {
	Assettype *string  `form:"assettype,omitempty" json:"assettype,omitempty" yaml:"assettype,omitempty" xml:"assettype,omitempty"`
	Name      []string `form:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty" xml:"name,omitempty"`
}

// Publicize creates AssettypeType from assettypeType
func (ut *assettypeType) Publicize() *AssettypeType {
	var pub AssettypeType
	if ut.Assettype != nil {
		pub.Assettype = ut.Assettype
	}
	if ut.Name != nil {
		pub.Name = ut.Name
	}
	return &pub
}

// AssettypeType user type.
type AssettypeType struct {
	Assettype *string  `form:"assettype,omitempty" json:"assettype,omitempty" yaml:"assettype,omitempty" xml:"assettype,omitempty"`
	Name      []string `form:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty" xml:"name,omitempty"`
}

// checkData user type.
type checkData struct {
	// Name of the checktype this check is
	ChecktypeName *string `form:"checktype_name,omitempty" json:"checktype_name,omitempty" yaml:"checktype_name,omitempty" xml:"checktype_name,omitempty"`
	// ID of the specific queue this check must we enqueued.
	// 		The queue must already be created in vulcan core.
	JobqueueID *uuid.UUID `form:"jobqueue_id,omitempty" json:"jobqueue_id,omitempty" yaml:"jobqueue_id,omitempty" xml:"jobqueue_id,omitempty"`
	// Configuration options for the Check. It should be in JSON format
	Options *string `form:"options,omitempty" json:"options,omitempty" yaml:"options,omitempty" xml:"options,omitempty"`
	// a tag
	Tag *string `form:"tag,omitempty" json:"tag,omitempty" yaml:"tag,omitempty" xml:"tag,omitempty"`
	// Target to be scanned. Can be a domain, hostname, IP or URL
	Target *string `form:"target,omitempty" json:"target,omitempty" yaml:"target,omitempty" xml:"target,omitempty"`
	// Webhook to callback after the Check has finished
	Webhook *string `form:"webhook,omitempty" json:"webhook,omitempty" yaml:"webhook,omitempty" xml:"webhook,omitempty"`
}

// Validate validates the checkData type instance.
func (ut *checkData) Validate() (err error) {
	if ut.Target == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "target"))
	}
	if ut.Options != nil {
		if ok := goa.ValidatePattern(`^[[:print:]]+`, *ut.Options); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`request.options`, *ut.Options, `^[[:print:]]+`))
		}
	}
	if ut.Options != nil {
		if utf8.RuneCountInString(*ut.Options) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`request.options`, *ut.Options, utf8.RuneCountInString(*ut.Options), 2, true))
		}
	}
	if ut.Tag != nil {
		if utf8.RuneCountInString(*ut.Tag) < 3 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`request.tag`, *ut.Tag, utf8.RuneCountInString(*ut.Tag), 3, true))
		}
	}
	if ut.Target != nil {
		if ok := goa.ValidatePattern(`^[[:print:]]+`, *ut.Target); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`request.target`, *ut.Target, `^[[:print:]]+`))
		}
	}
	if ut.Target != nil {
		if utf8.RuneCountInString(*ut.Target) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`request.target`, *ut.Target, utf8.RuneCountInString(*ut.Target), 1, true))
		}
	}
	if ut.Webhook != nil {
		if err2 := goa.ValidateFormat(goa.FormatURI, *ut.Webhook); err2 != nil {
			err = goa.MergeErrors(err, goa.InvalidFormatError(`request.webhook`, *ut.Webhook, goa.FormatURI, err2))
		}
	}
	if ut.Webhook != nil {
		if utf8.RuneCountInString(*ut.Webhook) < 3 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`request.webhook`, *ut.Webhook, utf8.RuneCountInString(*ut.Webhook), 3, true))
		}
	}
	return
}

// Publicize creates CheckData from checkData
func (ut *checkData) Publicize() *CheckData {
	var pub CheckData
	if ut.ChecktypeName != nil {
		pub.ChecktypeName = ut.ChecktypeName
	}
	if ut.JobqueueID != nil {
		pub.JobqueueID = ut.JobqueueID
	}
	if ut.Options != nil {
		pub.Options = ut.Options
	}
	if ut.Tag != nil {
		pub.Tag = ut.Tag
	}
	if ut.Target != nil {
		pub.Target = *ut.Target
	}
	if ut.Webhook != nil {
		pub.Webhook = ut.Webhook
	}
	return &pub
}

// CheckData user type.
type CheckData struct {
	// Name of the checktype this check is
	ChecktypeName *string `form:"checktype_name,omitempty" json:"checktype_name,omitempty" yaml:"checktype_name,omitempty" xml:"checktype_name,omitempty"`
	// ID of the specific queue this check must we enqueued.
	// 		The queue must already be created in vulcan core.
	JobqueueID *uuid.UUID `form:"jobqueue_id,omitempty" json:"jobqueue_id,omitempty" yaml:"jobqueue_id,omitempty" xml:"jobqueue_id,omitempty"`
	// Configuration options for the Check. It should be in JSON format
	Options *string `form:"options,omitempty" json:"options,omitempty" yaml:"options,omitempty" xml:"options,omitempty"`
	// a tag
	Tag *string `form:"tag,omitempty" json:"tag,omitempty" yaml:"tag,omitempty" xml:"tag,omitempty"`
	// Target to be scanned. Can be a domain, hostname, IP or URL
	Target string `form:"target" json:"target" yaml:"target" xml:"target"`
	// Webhook to callback after the Check has finished
	Webhook *string `form:"webhook,omitempty" json:"webhook,omitempty" yaml:"webhook,omitempty" xml:"webhook,omitempty"`
}

// Validate validates the CheckData type instance.
func (ut *CheckData) Validate() (err error) {
	if ut.Target == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "target"))
	}
	if ut.Options != nil {
		if ok := goa.ValidatePattern(`^[[:print:]]+`, *ut.Options); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`type.options`, *ut.Options, `^[[:print:]]+`))
		}
	}
	if ut.Options != nil {
		if utf8.RuneCountInString(*ut.Options) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`type.options`, *ut.Options, utf8.RuneCountInString(*ut.Options), 2, true))
		}
	}
	if ut.Tag != nil {
		if utf8.RuneCountInString(*ut.Tag) < 3 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`type.tag`, *ut.Tag, utf8.RuneCountInString(*ut.Tag), 3, true))
		}
	}
	if ok := goa.ValidatePattern(`^[[:print:]]+`, ut.Target); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`type.target`, ut.Target, `^[[:print:]]+`))
	}
	if utf8.RuneCountInString(ut.Target) < 1 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`type.target`, ut.Target, utf8.RuneCountInString(ut.Target), 1, true))
	}
	if ut.Webhook != nil {
		if err2 := goa.ValidateFormat(goa.FormatURI, *ut.Webhook); err2 != nil {
			err = goa.MergeErrors(err, goa.InvalidFormatError(`type.webhook`, *ut.Webhook, goa.FormatURI, err2))
		}
	}
	if ut.Webhook != nil {
		if utf8.RuneCountInString(*ut.Webhook) < 3 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`type.webhook`, *ut.Webhook, utf8.RuneCountInString(*ut.Webhook), 3, true))
		}
	}
	return
}

// checkPayload user type.
type checkPayload struct {
	Check *checkData `form:"check,omitempty" json:"check,omitempty" yaml:"check,omitempty" xml:"check,omitempty"`
}

// Validate validates the checkPayload type instance.
func (ut *checkPayload) Validate() (err error) {
	if ut.Check == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "check"))
	}
	if ut.Check != nil {
		if err2 := ut.Check.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// Publicize creates CheckPayload from checkPayload
func (ut *checkPayload) Publicize() *CheckPayload {
	var pub CheckPayload
	if ut.Check != nil {
		pub.Check = ut.Check.Publicize()
	}
	return &pub
}

// CheckPayload user type.
type CheckPayload struct {
	Check *CheckData `form:"check" json:"check" yaml:"check" xml:"check"`
}

// Validate validates the CheckPayload type instance.
func (ut *CheckPayload) Validate() (err error) {
	if ut.Check == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "check"))
	}
	if ut.Check != nil {
		if err2 := ut.Check.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// checktypeType user type.
type checktypeType struct {
	Description *string    `form:"description,omitempty" json:"description,omitempty" yaml:"description,omitempty" xml:"description,omitempty"`
	Enabled     *bool      `form:"enabled,omitempty" json:"enabled,omitempty" yaml:"enabled,omitempty" xml:"enabled,omitempty"`
	ID          *uuid.UUID `form:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty" xml:"id,omitempty"`
	// Image that needs to be pulled to run the Check of this type
	Image *string `form:"image,omitempty" json:"image,omitempty" yaml:"image,omitempty" xml:"image,omitempty"`
	Name  *string `form:"name,omitempty" json:"name,omitempty" yaml:"name,omitempty" xml:"name,omitempty"`
	// Default configuration options for the Checktype. It should be in JSON format
	Options *string `form:"options,omitempty" json:"options,omitempty" yaml:"options,omitempty" xml:"options,omitempty"`
	// Specifies the maximum amount of time that the check should be running before it's killed
	Timeout *int `form:"timeout,omitempty" json:"timeout,omitempty" yaml:"timeout,omitempty" xml:"timeout,omitempty"`
}

// Validate validates the checktypeType type instance.
func (ut *checktypeType) Validate() (err error) {
	if ut.ID == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "id"))
	}
	if ut.Name == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "name"))
	}
	if ut.Image == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "image"))
	}
	if ut.Description != nil {
		if ok := goa.ValidatePattern(`^[[:print:]]+`, *ut.Description); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`request.description`, *ut.Description, `^[[:print:]]+`))
		}
	}
	if ut.Description != nil {
		if utf8.RuneCountInString(*ut.Description) > 255 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`request.description`, *ut.Description, utf8.RuneCountInString(*ut.Description), 255, false))
		}
	}
	if ut.Name != nil {
		if ok := goa.ValidatePattern(`^[[:word:]]+`, *ut.Name); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`request.name`, *ut.Name, `^[[:word:]]+`))
		}
	}
	if ut.Name != nil {
		if utf8.RuneCountInString(*ut.Name) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`request.name`, *ut.Name, utf8.RuneCountInString(*ut.Name), 1, true))
		}
	}
	return
}

// Publicize creates ChecktypeType from checktypeType
func (ut *checktypeType) Publicize() *ChecktypeType {
	var pub ChecktypeType
	if ut.Description != nil {
		pub.Description = ut.Description
	}
	if ut.Enabled != nil {
		pub.Enabled = ut.Enabled
	}
	if ut.ID != nil {
		pub.ID = *ut.ID
	}
	if ut.Image != nil {
		pub.Image = *ut.Image
	}
	if ut.Name != nil {
		pub.Name = *ut.Name
	}
	if ut.Options != nil {
		pub.Options = ut.Options
	}
	if ut.Timeout != nil {
		pub.Timeout = ut.Timeout
	}
	return &pub
}

// ChecktypeType user type.
type ChecktypeType struct {
	Description *string   `form:"description,omitempty" json:"description,omitempty" yaml:"description,omitempty" xml:"description,omitempty"`
	Enabled     *bool     `form:"enabled,omitempty" json:"enabled,omitempty" yaml:"enabled,omitempty" xml:"enabled,omitempty"`
	ID          uuid.UUID `form:"id" json:"id" yaml:"id" xml:"id"`
	// Image that needs to be pulled to run the Check of this type
	Image string `form:"image" json:"image" yaml:"image" xml:"image"`
	Name  string `form:"name" json:"name" yaml:"name" xml:"name"`
	// Default configuration options for the Checktype. It should be in JSON format
	Options *string `form:"options,omitempty" json:"options,omitempty" yaml:"options,omitempty" xml:"options,omitempty"`
	// Specifies the maximum amount of time that the check should be running before it's killed
	Timeout *int `form:"timeout,omitempty" json:"timeout,omitempty" yaml:"timeout,omitempty" xml:"timeout,omitempty"`
}

// Validate validates the ChecktypeType type instance.
func (ut *ChecktypeType) Validate() (err error) {

	if ut.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "name"))
	}
	if ut.Image == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "image"))
	}
	if ut.Description != nil {
		if ok := goa.ValidatePattern(`^[[:print:]]+`, *ut.Description); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`type.description`, *ut.Description, `^[[:print:]]+`))
		}
	}
	if ut.Description != nil {
		if utf8.RuneCountInString(*ut.Description) > 255 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`type.description`, *ut.Description, utf8.RuneCountInString(*ut.Description), 255, false))
		}
	}
	if ok := goa.ValidatePattern(`^[[:word:]]+`, ut.Name); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`type.name`, ut.Name, `^[[:word:]]+`))
	}
	if utf8.RuneCountInString(ut.Name) < 1 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`type.name`, ut.Name, utf8.RuneCountInString(ut.Name), 1, true))
	}
	return
}

// fileScanPayload user type.
type fileScanPayload struct {
	// Program ID
	ProgramID *string `form:"program_id,omitempty" json:"program_id,omitempty" yaml:"program_id,omitempty" xml:"program_id,omitempty"`
	// Tag associated with the scan
	Tag *string `form:"tag,omitempty" json:"tag,omitempty" yaml:"tag,omitempty" xml:"tag,omitempty"`
	// Upload
	Upload *string `form:"upload,omitempty" json:"upload,omitempty" yaml:"upload,omitempty" xml:"upload,omitempty"`
}

// Validate validates the fileScanPayload type instance.
func (ut *fileScanPayload) Validate() (err error) {
	if ut.Upload == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "upload"))
	}
	return
}

// Publicize creates FileScanPayload from fileScanPayload
func (ut *fileScanPayload) Publicize() *FileScanPayload {
	var pub FileScanPayload
	if ut.ProgramID != nil {
		pub.ProgramID = ut.ProgramID
	}
	if ut.Tag != nil {
		pub.Tag = ut.Tag
	}
	if ut.Upload != nil {
		pub.Upload = *ut.Upload
	}
	return &pub
}

// FileScanPayload user type.
type FileScanPayload struct {
	// Program ID
	ProgramID *string `form:"program_id,omitempty" json:"program_id,omitempty" yaml:"program_id,omitempty" xml:"program_id,omitempty"`
	// Tag associated with the scan
	Tag *string `form:"tag,omitempty" json:"tag,omitempty" yaml:"tag,omitempty" xml:"tag,omitempty"`
	// Upload
	Upload string `form:"upload" json:"upload" yaml:"upload" xml:"upload"`
}

// Validate validates the FileScanPayload type instance.
func (ut *FileScanPayload) Validate() (err error) {
	if ut.Upload == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "upload"))
	}
	return
}

// scanChecksPayload user type.
type scanChecksPayload struct {
	Checks []*checkPayload `form:"checks,omitempty" json:"checks,omitempty" yaml:"checks,omitempty" xml:"checks,omitempty"`
}

// Validate validates the scanChecksPayload type instance.
func (ut *scanChecksPayload) Validate() (err error) {
	if ut.Checks == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "checks"))
	}
	for _, e := range ut.Checks {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// Publicize creates ScanChecksPayload from scanChecksPayload
func (ut *scanChecksPayload) Publicize() *ScanChecksPayload {
	var pub ScanChecksPayload
	if ut.Checks != nil {
		pub.Checks = make([]*CheckPayload, len(ut.Checks))
		for i2, elem2 := range ut.Checks {
			pub.Checks[i2] = elem2.Publicize()
		}
	}
	return &pub
}

// ScanChecksPayload user type.
type ScanChecksPayload struct {
	Checks []*CheckPayload `form:"checks" json:"checks" yaml:"checks" xml:"checks"`
}

// Validate validates the ScanChecksPayload type instance.
func (ut *ScanChecksPayload) Validate() (err error) {
	if ut.Checks == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "checks"))
	}
	for _, e := range ut.Checks {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// scanData user type.
type scanData struct {
	// Number of checks of the scan
	Size *int `form:"size,omitempty" json:"size,omitempty" yaml:"size,omitempty" xml:"size,omitempty"`
}

// Validate validates the scanData type instance.
func (ut *scanData) Validate() (err error) {
	if ut.Size == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "size"))
	}
	return
}

// Publicize creates ScanData from scanData
func (ut *scanData) Publicize() *ScanData {
	var pub ScanData
	if ut.Size != nil {
		pub.Size = *ut.Size
	}
	return &pub
}

// ScanData user type.
type ScanData struct {
	// Number of checks of the scan
	Size int `form:"size" json:"size" yaml:"size" xml:"size"`
}

// scanPayload user type.
type scanPayload struct {
	Scan *scanChecksPayload `form:"scan,omitempty" json:"scan,omitempty" yaml:"scan,omitempty" xml:"scan,omitempty"`
}

// Validate validates the scanPayload type instance.
func (ut *scanPayload) Validate() (err error) {
	if ut.Scan != nil {
		if err2 := ut.Scan.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// Publicize creates ScanPayload from scanPayload
func (ut *scanPayload) Publicize() *ScanPayload {
	var pub ScanPayload
	if ut.Scan != nil {
		pub.Scan = ut.Scan.Publicize()
	}
	return &pub
}

// ScanPayload user type.
type ScanPayload struct {
	Scan *ScanChecksPayload `form:"scan,omitempty" json:"scan,omitempty" yaml:"scan,omitempty" xml:"scan,omitempty"`
}

// Validate validates the ScanPayload type instance.
func (ut *ScanPayload) Validate() (err error) {
	if ut.Scan != nil {
		if err2 := ut.Scan.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}
