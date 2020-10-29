package cmd

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"
	"time"

	cli "github.com/adevinta/vulcan-core-cli"

	"github.com/spf13/cobra"
)

// scanCmd represents the scan command
var scanCmd = &cobra.Command{
	Use:   "scan targets_file checktypes_file",
	Short: "Executes a scan for the specified checktypes and assets",
	Long: `Executes a scan for the specified checktypes and assets.
	The command takes two params.
	The first parameter is a path to a file containing a list of targets one per line e.g:
	example.com
	anotherexample.com;DomainName
	Each line must have the following format:
	target;assettype
	Only the first field name is mandatory.
	The fields are separated by a semicolon ";".
	The second parameter is a path to a file containing a list of checktypes to execute.
	Each line must have the following format:
	name;options;queueid
	Only the first field name is mandatory.
	The fields are separated by a semicolon ";".`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			return fmt.Errorf("incorrect number of args, want 2, got %v", len(args))
		}

		t := time.Now()
		if verbose {
			log.Printf("start time: %v\n", t)
		}

		var err error
		c, err = cli.NewCLI(scheme, host)
		if err != nil {
			return err
		}
		if verbose {
			log.Printf("new cli created for %v://%v\n", scheme, host)
		}

		log.Println("parsing the targets and checktypes files")
		// Determine the checktypes and assets for the checks that will be
		// executed.
		checks, err := createScanChecks(args[0], args[1])
		if err != nil {
			return err
		}
		if verbose {
			log.Printf("%v\n", checks.ToString())
		}

		log.Println("executing the scan")
		s, err := c.RunScan(checks, dryRun, multiPart)
		if err != nil {
			return err
		}

		if dryRun != "" {
			log.Printf("Dry run executed, payload writted to file:%s", dryRun)
			return nil
		}

		if verbose {
			log.Printf("%v\n", s.ToString())
		}

		resultsFile := filepath.Join(outDir, s.ID.String()+".gob")

		log.Println("saving the state")
		if err := s.SaveState(resultsFile); err != nil {
			return err
		}

		t2 := time.Since(t)
		log.Printf("time elapsed: %v\n", t2)
		if verbose {
			log.Printf("end time: %v\n", time.Now())
		}

		log.Println("scan successfully launched and state stored:")
		fmt.Println(resultsFile)

		return nil
	},
}

func init() {
	RootCmd.AddCommand(scanCmd)

	scanCmd.Flags().StringVarP(&outDir, "odir", "o", "/tmp", "directory where the result state is saved")
	scanCmd.Flags().StringVarP(&dryRun, "dry", "y", "", "file to store dry run output payload. Don't execute the scan but generate a json with the payload")
	scanCmd.Flags().BoolVarP(&multiPart, "multipart", "m", false, "use multipart upload file api")
}

type asset struct {
	target    string
	assetType string
}

func createScanChecks(targetsFile, checktypesFile string) (cli.Checks, error) {
	assets, err := parseTargetsFile(targetsFile)
	if err != nil {
		return nil, err
	}
	log.Printf("num targets: %v\n", len(assets))
	if verbose {
		log.Printf("targets %v\n", assets)
	}

	checktypes, err := parseChecktypesFile(checktypesFile)
	if err != nil {
		return nil, err
	}
	log.Printf("num checktypes: %v\n", len(checktypes))
	if verbose {
		log.Printf("checktypes %v\n", checktypes)
	}

	checks := cli.Checks{}

	for _, checktype := range checktypes {
		var cliAssets []cli.Asset

		for _, asset := range assets {
			cliAsset := cli.Asset{
				Target:    asset.target,
				AssetType: asset.assetType,
				Options:   checktype.DefaultOptions,
				QueueID:   checktype.QueueID,
			}
			cliAssets = append(cliAssets, cliAsset)
		}
		checks[checktype.Name] = cliAssets
	}

	return checks, nil
}

// parseTargetsFile is used to read assets from a file.
func parseTargetsFile(file string) ([]asset, error) {
	targetLines, err := cli.ReadLines(file)
	if err != nil {
		return nil, err
	}

	var assets []asset

	for _, line := range targetLines {
		// Every line of the file is an asset: a target and optionally the asset type.
		params := strings.SplitN(line, ";", 2)
		if len(params) < 1 {
			return nil, fmt.Errorf("malformed asset in line: %v", line)
		}

		asset := asset{target: params[0]}
		if len(params) > 1 {
			asset.assetType = params[1]
		}
		assets = append(assets, asset)
	}

	return assets, nil
}
