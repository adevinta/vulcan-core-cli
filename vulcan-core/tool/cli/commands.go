// Code generated by goagen v1.4.3, DO NOT EDIT.
//
// API "vulcan-core": CLI Commands
//
// Command:
// $ main

package cli

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/adevinta/vulcan-core-cli/vulcan-core/client"
	"github.com/goadesign/goa"
	goaclient "github.com/goadesign/goa/client"
	uuid "github.com/goadesign/goa/uuid"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type (
	// IndexAssettypesCommand is the command line data structure for the index action of Assettypes
	IndexAssettypesCommand struct {
		PrettyPrint bool
	}

	// ShowChecksCommand is the command line data structure for the show action of Checks
	ShowChecksCommand struct {
		ID          string
		PrettyPrint bool
	}

	// IndexChecktypesCommand is the command line data structure for the index action of Checktypes
	IndexChecktypesCommand struct {
		Enabled     string
		Name        string
		PrettyPrint bool
	}

	// ShowChecktypesCommand is the command line data structure for the show action of Checktypes
	ShowChecktypesCommand struct {
		ID          string
		PrettyPrint bool
	}

	// AbortScansCommand is the command line data structure for the abort action of Scans
	AbortScansCommand struct {
		ID          string
		PrettyPrint bool
	}

	// ChecksScansCommand is the command line data structure for the checks action of Scans
	ChecksScansCommand struct {
		ID          string
		PrettyPrint bool
	}

	// CreateScansCommand is the command line data structure for the create action of Scans
	CreateScansCommand struct {
		Payload     string
		ContentType string
		PrettyPrint bool
	}

	// IndexScansCommand is the command line data structure for the index action of Scans
	IndexScansCommand struct {
		PrettyPrint bool
	}

	// ShowScansCommand is the command line data structure for the show action of Scans
	ShowScansCommand struct {
		ID          string
		PrettyPrint bool
	}

	// StatsScansCommand is the command line data structure for the stats action of Scans
	StatsScansCommand struct {
		ID          string
		PrettyPrint bool
	}

	// ShowHealthcheckCommand is the command line data structure for the show action of healthcheck
	ShowHealthcheckCommand struct {
		PrettyPrint bool
	}

	// DownloadCommand is the command line data structure for the download command.
	DownloadCommand struct {
		// OutFile is the path to the download output file.
		OutFile string
	}
)

// RegisterCommands registers the resource action CLI commands.
func RegisterCommands(app *cobra.Command, c *client.Client) {
	var command, sub *cobra.Command
	command = &cobra.Command{
		Use:   "abort",
		Short: `Abort a scan`,
	}
	tmp1 := new(AbortScansCommand)
	sub = &cobra.Command{
		Use:   `scans ["/v1/scans/ID/abort"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp1.Run(c, args) },
	}
	tmp1.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp1.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "checks",
		Short: `Get checks of the Scan`,
	}
	tmp2 := new(ChecksScansCommand)
	sub = &cobra.Command{
		Use:   `scans ["/v1/scans/ID/checks"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp2.Run(c, args) },
	}
	tmp2.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp2.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "create",
		Short: `Create a new Scan`,
	}
	tmp3 := new(CreateScansCommand)
	sub = &cobra.Command{
		Use:   `scans ["/v1/scans/"]`,
		Short: ``,
		Long: `

Payload example:

{
   "external_id": "cfb906e1-a99b-4613-849a-f2bbce19db63",
   "scheduled_time": "2014-02-20T01:06:30Z",
   "tag": "Et sunt sapiente nostrum aliquid natus quae.",
   "target_groups": [
      {
         "checktypes_group": {
            "checktypes": [
               {
                  "name": "Quis doloribus.",
                  "options": "Voluptates maxime sunt."
               },
               {
                  "name": "Quis doloribus.",
                  "options": "Voluptates maxime sunt."
               }
            ],
            "name": "At necessitatibus qui nulla qui ea."
         },
         "target_group": {
            "name": "Ut consequatur porro et dolorum repellat ut.",
            "options": "Quis neque tenetur.",
            "targets": [
               {
                  "identifier": "Sequi quas voluptatem.",
                  "options": "Praesentium sint sit fugiat.",
                  "type": "Voluptatem et velit et iusto aut voluptas."
               },
               {
                  "identifier": "Sequi quas voluptatem.",
                  "options": "Praesentium sint sit fugiat.",
                  "type": "Voluptatem et velit et iusto aut voluptas."
               },
               {
                  "identifier": "Sequi quas voluptatem.",
                  "options": "Praesentium sint sit fugiat.",
                  "type": "Voluptatem et velit et iusto aut voluptas."
               }
            ]
         }
      },
      {
         "checktypes_group": {
            "checktypes": [
               {
                  "name": "Quis doloribus.",
                  "options": "Voluptates maxime sunt."
               },
               {
                  "name": "Quis doloribus.",
                  "options": "Voluptates maxime sunt."
               }
            ],
            "name": "At necessitatibus qui nulla qui ea."
         },
         "target_group": {
            "name": "Ut consequatur porro et dolorum repellat ut.",
            "options": "Quis neque tenetur.",
            "targets": [
               {
                  "identifier": "Sequi quas voluptatem.",
                  "options": "Praesentium sint sit fugiat.",
                  "type": "Voluptatem et velit et iusto aut voluptas."
               },
               {
                  "identifier": "Sequi quas voluptatem.",
                  "options": "Praesentium sint sit fugiat.",
                  "type": "Voluptatem et velit et iusto aut voluptas."
               },
               {
                  "identifier": "Sequi quas voluptatem.",
                  "options": "Praesentium sint sit fugiat.",
                  "type": "Voluptatem et velit et iusto aut voluptas."
               }
            ]
         }
      }
   ],
   "trigger": "Vero maiores sit vel."
}`,
		RunE: func(cmd *cobra.Command, args []string) error { return tmp3.Run(c, args) },
	}
	tmp3.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp3.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "index",
		Short: `index action`,
	}
	tmp4 := new(IndexAssettypesCommand)
	sub = &cobra.Command{
		Use:   `assettypes ["/v1/assettypes"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp4.Run(c, args) },
	}
	tmp4.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp4.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp5 := new(IndexChecktypesCommand)
	sub = &cobra.Command{
		Use:   `checktypes ["/v1/checktypes"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp5.Run(c, args) },
	}
	tmp5.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp5.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp6 := new(IndexScansCommand)
	sub = &cobra.Command{
		Use:   `scans ["/v1/scans"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp6.Run(c, args) },
	}
	tmp6.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp6.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "show",
		Short: `show action`,
	}
	tmp7 := new(ShowChecksCommand)
	sub = &cobra.Command{
		Use:   `checks ["/v1/checks/ID"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp7.Run(c, args) },
	}
	tmp7.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp7.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp8 := new(ShowChecktypesCommand)
	sub = &cobra.Command{
		Use:   `checktypes ["/v1/checktypes/ID"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp8.Run(c, args) },
	}
	tmp8.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp8.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp9 := new(ShowScansCommand)
	sub = &cobra.Command{
		Use:   `scans ["/v1/scans/ID"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp9.Run(c, args) },
	}
	tmp9.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp9.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp10 := new(ShowHealthcheckCommand)
	sub = &cobra.Command{
		Use:   `healthcheck ["/v1/healthcheck/"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp10.Run(c, args) },
	}
	tmp10.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp10.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "stats",
		Short: `Get stats of the Scan`,
	}
	tmp11 := new(StatsScansCommand)
	sub = &cobra.Command{
		Use:   `scans ["/v1/scans/ID/stats"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp11.Run(c, args) },
	}
	tmp11.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp11.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)

	dl := new(DownloadCommand)
	dlc := &cobra.Command{
		Use:   "download [PATH]",
		Short: "Download file with given path",
		RunE: func(cmd *cobra.Command, args []string) error {
			return dl.Run(c, args)
		},
	}
	dlc.Flags().StringVar(&dl.OutFile, "out", "", "Output file")
	app.AddCommand(dlc)
}

func intFlagVal(name string, parsed int) *int {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func float64FlagVal(name string, parsed float64) *float64 {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func boolFlagVal(name string, parsed bool) *bool {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func stringFlagVal(name string, parsed string) *string {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func hasFlag(name string) bool {
	for _, arg := range os.Args[1:] {
		if strings.HasPrefix(arg, "--"+name) {
			return true
		}
	}
	return false
}

func jsonVal(val string) (*interface{}, error) {
	var t interface{}
	err := json.Unmarshal([]byte(val), &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func jsonArray(ins []string) ([]interface{}, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []interface{}
	for _, id := range ins {
		val, err := jsonVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, val)
	}
	return vals, nil
}

func timeVal(val string) (*time.Time, error) {
	t, err := time.Parse(time.RFC3339, val)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func timeArray(ins []string) ([]time.Time, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []time.Time
	for _, id := range ins {
		val, err := timeVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

func uuidVal(val string) (*uuid.UUID, error) {
	t, err := uuid.FromString(val)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func uuidArray(ins []string) ([]uuid.UUID, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []uuid.UUID
	for _, id := range ins {
		val, err := uuidVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

func float64Val(val string) (*float64, error) {
	t, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func float64Array(ins []string) ([]float64, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []float64
	for _, id := range ins {
		val, err := float64Val(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

func boolVal(val string) (*bool, error) {
	t, err := strconv.ParseBool(val)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func boolArray(ins []string) ([]bool, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []bool
	for _, id := range ins {
		val, err := boolVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

// Run downloads files with given paths.
func (cmd *DownloadCommand) Run(c *client.Client, args []string) error {
	var (
		fnf func(context.Context, string) (int64, error)
		fnd func(context.Context, string, string) (int64, error)

		rpath   = args[0]
		outfile = cmd.OutFile
		logger  = goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
		ctx     = goa.WithLogger(context.Background(), logger)
		err     error
	)

	if rpath[0] != '/' {
		rpath = "/" + rpath
	}
	if rpath == "/v1/swagger.json" {
		fnf = c.DownloadSwaggerJSON
		if outfile == "" {
			outfile = "swagger.json"
		}
		goto found
	}
	return fmt.Errorf("don't know how to download %s", rpath)
found:
	ctx = goa.WithLogContext(ctx, "file", outfile)
	if fnf != nil {
		_, err = fnf(ctx, outfile)
	} else {
		_, err = fnd(ctx, rpath, outfile)
	}
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	return nil
}

// Run makes the HTTP request corresponding to the IndexAssettypesCommand command.
func (cmd *IndexAssettypesCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/v1/assettypes"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.IndexAssettypes(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *IndexAssettypesCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
}

// Run makes the HTTP request corresponding to the ShowChecksCommand command.
func (cmd *ShowChecksCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/v1/checks/%v", cmd.ID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ShowChecks(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ShowChecksCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var id string
	cc.Flags().StringVar(&cmd.ID, "id", id, ``)
}

// Run makes the HTTP request corresponding to the IndexChecktypesCommand command.
func (cmd *IndexChecktypesCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/v1/checktypes"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.IndexChecktypes(ctx, path, stringFlagVal("enabled", cmd.Enabled), stringFlagVal("name", cmd.Name))
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *IndexChecktypesCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Enabled, "enabled", "true", ``)
	var name string
	cc.Flags().StringVar(&cmd.Name, "name", name, ``)
}

// Run makes the HTTP request corresponding to the ShowChecktypesCommand command.
func (cmd *ShowChecktypesCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/v1/checktypes/%v", cmd.ID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ShowChecktypes(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ShowChecktypesCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var id string
	cc.Flags().StringVar(&cmd.ID, "id", id, ``)
}

// Run makes the HTTP request corresponding to the AbortScansCommand command.
func (cmd *AbortScansCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/v1/scans/%v/abort", cmd.ID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.AbortScans(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *AbortScansCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var id string
	cc.Flags().StringVar(&cmd.ID, "id", id, ``)
}

// Run makes the HTTP request corresponding to the ChecksScansCommand command.
func (cmd *ChecksScansCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/v1/scans/%v/checks", cmd.ID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ChecksScans(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ChecksScansCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var id string
	cc.Flags().StringVar(&cmd.ID, "id", id, ``)
}

// Run makes the HTTP request corresponding to the CreateScansCommand command.
func (cmd *CreateScansCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/v1/scans/"
	}
	var payload client.ScanPayload
	if cmd.Payload != "" {
		err := json.Unmarshal([]byte(cmd.Payload), &payload)
		if err != nil {
			return fmt.Errorf("failed to deserialize payload: %s", err)
		}
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.CreateScans(ctx, path, &payload)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *CreateScansCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	cc.Flags().StringVar(&cmd.Payload, "payload", "", "Request body encoded in JSON")
	cc.Flags().StringVar(&cmd.ContentType, "content", "", "Request content type override, e.g. 'application/x-www-form-urlencoded'")
}

// Run makes the HTTP request corresponding to the IndexScansCommand command.
func (cmd *IndexScansCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/v1/scans"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.IndexScans(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *IndexScansCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
}

// Run makes the HTTP request corresponding to the ShowScansCommand command.
func (cmd *ShowScansCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/v1/scans/%v", cmd.ID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ShowScans(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ShowScansCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var id string
	cc.Flags().StringVar(&cmd.ID, "id", id, ``)
}

// Run makes the HTTP request corresponding to the StatsScansCommand command.
func (cmd *StatsScansCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/v1/scans/%v/stats", cmd.ID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.StatsScans(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *StatsScansCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var id string
	cc.Flags().StringVar(&cmd.ID, "id", id, ``)
}

// Run makes the HTTP request corresponding to the ShowHealthcheckCommand command.
func (cmd *ShowHealthcheckCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/v1/healthcheck/"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ShowHealthcheck(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ShowHealthcheckCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
}
