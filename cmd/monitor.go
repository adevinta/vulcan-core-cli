/*
Copyright 2019 Adevinta
*/

package cmd

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"
	"time"

	cli "github.com/adevinta/vulcan-core-cli"

	"github.com/jroimartin/gocui"
	"github.com/spf13/cobra"
)

const spinnerSize = 20

var interval int

// monitorCmd represents the monitor command
var monitorCmd = &cobra.Command{
	Use:   "monitor scan_state_file",
	Short: "Monitors the status of the checks executed in a Scan",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("incorrect number of args, want 1, got %v", len(args))
		}

		t := time.Now()
		if verbose {
			log.Printf("start time: %v\n", t)
		}

		var err error
		c, err = cli.NewCLI(scheme, host, false)
		if err != nil {
			return err
		}
		if verbose {
			log.Printf("new cli created for %v://%v\n", scheme, host)
		}

		log.Printf("monitoring scan from %v\n", args[0])

		if err := monitorScan(args[0]); err != nil {
			return err
		}

		t2 := time.Since(t)
		log.Printf("time elapsed: %v\n", t2)
		if verbose {
			log.Printf("end time: %v\n", t)
		}

		return nil
	},
}

func init() {
	RootCmd.AddCommand(monitorCmd)

	monitorCmd.Flags().IntVarP(&interval, "interval", "i", 2, "time between refreshing the scan state")
}

func monitorScan(stateFile string) error {
	scan, err := cli.LoadState(stateFile)
	if err != nil {
		return err
	}

	g, err := gocui.NewGui(gocui.OutputNormal)

	if err != nil {
		return err
	}
	defer g.Close()

	errors := make(chan error)
	go runUI(g, errors)

	if err := updateAndPrint(scan, stateFile, g); err != nil {
		return err
	}

	if scan.IsFinished() {
		if verbose {
			log.Println("scan finished")
		}
		return nil
	}

	go spinner(g)

	for {
		select {
		case err := <-errors:
			if err != gocui.ErrQuit {
				return err
			}
			return nil
		case <-time.After(time.Duration(interval) * time.Second):
			if err := updateAndPrint(scan, stateFile, g); err != nil {
				return err
			}

			if scan.IsFinished() {
				return nil
			}
		}
	}
}

func runUI(g *gocui.Gui, errors chan error) {
	g.SetManagerFunc(layout)

	if err := keybindings(g); err != nil {
		errors <- err
		return
	}

	if err := g.MainLoop(); err != nil {
		errors <- err
		return
	}
}

func updateAndPrint(scan *cli.Scan, stateFile string, g *gocui.Gui) error {
	if err := c.UpdateStats(scan); err != nil {
		return err
	}

	if err := scan.SaveState(stateFile); err != nil {
		return err
	}

	printStatus(g, scan)

	dir := filepath.Dir(stateFile)
	fp := filepath.Join(dir, scan.ID.String()+"-stats.txt")

	return scan.Stats.PrintToFile(fp)
}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("ctr", maxX/2-35, maxY/2-6, maxX/2+35, maxY/2+6); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		fmt.Fprintf(v, "Status update pending to start...\n\n")
		for status, _ := range cli.ScanStatus {
			fmt.Fprintf(v, "%v: ?\n", status)
		}
	}

	if _, err := g.SetView("spinner", maxX/2-35, maxY/2+7, maxX/2+35, maxY/2+9); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
	}

	return nil
}

func keybindings(g *gocui.Gui) error {
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func spinner(g *gocui.Gui) {
	forward := false

	for ctr := 0; ; ctr++ {
		g.Update(func(g *gocui.Gui) error {
			v, err := g.View("spinner")
			if err != nil {
				return err
			}
			w, _ := v.Size()

			x := ctr % (w - spinnerSize)

			if x == 0 {
				forward = !forward
			}

			var padding int
			if forward {
				padding = x
			} else {
				padding = w - x - spinnerSize
			}
			bar := strings.Repeat(" ", padding) + strings.Repeat("█", spinnerSize)

			v.Clear()
			fmt.Fprint(v, bar)

			return nil
		})

		time.Sleep(50 * time.Millisecond)
	}
}

func printStatus(g *gocui.Gui, scan *cli.Scan) {
	g.Update(func(g *gocui.Gui) error {
		view, err := g.View("ctr")
		if err != nil {
			return err
		}

		view.Clear()

		scan.Stats.Print(view)

		return nil
	})
}
