/*
Copyright 2019 Adevinta
*/

package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/goadesign/goa/uuid"
	"github.com/spf13/cobra"

	cli "github.com/adevinta/vulcan-core-cli"
)

func init() {
	RootCmd.AddCommand(abortCmd)
}

var abortCmd = &cobra.Command{
	Use:   "abort scan_id",
	Short: "Aborts an existing scan",
	Long:  `Signals vulcan core to abort a scan`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("incorrect number of args, want 1, got %v", len(args))
		}
		id, err := uuid.FromString(args[0])
		if err != nil {
			return err
		}
		t := time.Now()
		if verbose {
			log.Printf("start time: %v\n", t)
		}

		c, err = cli.NewCLI(scheme, host, false)
		if err != nil {
			return err
		}
		if verbose {
			log.Printf("new cli created for %v://%v\n", scheme, host)
		}
		err = c.AbortScan(id)
		if err != nil {
			fmt.Printf("error: %v\n", err)
			return nil
		}
		fmt.Println("abort scan command accepted")
		return nil
	},
}
