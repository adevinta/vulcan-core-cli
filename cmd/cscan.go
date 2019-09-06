package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
	"time"

	cli "github.com/adevinta/vulcan-core-cli"

	"github.com/spf13/cobra"
)

var (
	targetsFile, blacklistFile, checktypesFile, continuousDir, outDir, dryRun string
	multiPart                                                                 bool
)

// cscanCmd represents the cscan command
var cscanCmd = &cobra.Command{
	Use:   "cscan",
	Short: "Executes a scan using the assets and checktypes defined for Vulcan Continuous",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 0 {
			return fmt.Errorf("incorrect number of args, want 0, got %v", len(args))
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
		checks, err := createChecks()
		if err != nil {
			return err
		}
		if verbose {
			log.Printf("%+v\n", checks.ToString())
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
	RootCmd.AddCommand(cscanCmd)

	cscanCmd.Flags().StringVarP(&targetsFile, "targets", "t", "", "a list of assets to scan, overriding the default info from continuous directory")
	cscanCmd.Flags().StringVarP(&blacklistFile, "blacklist", "b", "", "a blacklist of assets that will be ignored in the scan")
	cscanCmd.Flags().StringVarP(&checktypesFile, "checktypes", "c", "", "a list of checktypes to launch in the scan, overriding the default info from continuous directory")
	cscanCmd.Flags().StringVarP(&continuousDir, "cdir", "d", "/opt/crontinuous/assets", "directory where the vulcan continuous assets are defined")
	cscanCmd.Flags().StringVarP(&outDir, "odir", "o", "/opt/crontinuous/scans", "directory where the result state is saved")
	cscanCmd.Flags().StringVarP(&dryRun, "dry", "y", "", "file to store dry run results. Don't execute the scan but generate a json with the payload")
	cscanCmd.Flags().BoolVarP(&multiPart, "multipart", "m", false, "use multipart upload file api")
}

type Checktype struct {
	Name           string
	DefaultOptions string
	QueueID        string
}

// createChecks from targets file or checktypes file
func createChecks() (cli.Checks, error) {
	var (
		targetsFilter    []string
		blacklistFilter  []string
		checktypesFilter []Checktype
		err              error
	)

	// A filter for targets or a filter for checktypes can be specified,
	// but not both at the same time.
	if targetsFile != "" && checktypesFile != "" {
		return nil, errors.New("both targets and checktypes files provided")
	}

	if targetsFile != "" {
		if verbose {
			log.Printf("using %v to create a targets filter\n", targetsFile)
		}
		// Read the targets file to determine which will be the subset of
		// targets that will be scanned.
		targetsFilter, err = cli.ReadLines(targetsFile)
		if err != nil {
			return nil, err
		}
		if verbose {
			log.Printf("targets filter: %v\n", targetsFilter)
		}
	}

	if blacklistFile != "" {
		if verbose {
			log.Printf("using %v to create a blacklist filter\n", blacklistFile)
		}
		// Read the blacklist file to determine which will be the subset of
		// targets that will NOT be scanned.
		blacklistFilter, err = cli.ReadLines(blacklistFile)
		if err != nil {
			return nil, err
		}
		if verbose {
			log.Printf("blacklist filter: %v\n", blacklistFilter)
		}
	}

	if checktypesFile != "" {
		if verbose {
			log.Printf("using %v to create a checktypes filter\n", checktypesFile)
		}
		// Read the checktypes file to determine which will be the subset of
		// checktypes that will be executed.
		checktypesFilter, err = parseChecktypesFile(checktypesFile)
		if err != nil {
			return nil, err
		}
		if verbose {
			log.Printf("checktypes filter: %v\n", checktypesFilter)
		}
	}

	return parseVCDir(continuousDir, targetsFilter, blacklistFilter, checktypesFilter)
}

// parseChecktypesFile is used to read from a file a checktypes filter.
func parseChecktypesFile(file string) ([]Checktype, error) {
	checktypeLines, err := cli.ReadLines(file)
	if err != nil {
		return nil, err
	}

	var checktypes []Checktype

	for _, line := range checktypeLines {
		// Every line of the file is a checktype, and options for the checktype
		// might be specified too.
		params := strings.SplitN(line, ";", 3)
		if len(params) < 1 {
			return nil, fmt.Errorf("malformed checktype in line: %v", line)
		}

		checktype := Checktype{Name: params[0]}
		if len(params) > 2 {
			checktype.QueueID = params[2]
		}
		if len(params) > 1 {
			checktype.DefaultOptions = params[1]
		}
		checktypes = append(checktypes, checktype)
	}

	return checktypes, nil
}

// parseVCDir returns the assets to be scanned, by loading the information stored in the vcdir.
// It accepts filters to execute only some checktypes, to ignore some assets or to scan only
// some assets instead of using all those defined in the vcdir.
func parseVCDir(dir string, targetsFilter []string, blacklistFilter []string, checktypesFilter []Checktype) (cli.Checks, error) {
	// The vcdir contains files for every checktype.
	checktypeFiles, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	if verbose {
		log.Printf("checktypes in vcdir: %v\n", checktypeFiles)
	}

	checks := cli.Checks{}

	for _, f := range checktypeFiles {
		// The name of the file is the checktype name.
		checktypeFilename := f.Name()

		checktype, ok := findChecktype(checktypeFilename, checktypesFilter)
		if len(checktypesFilter) > 0 && !ok {
			if verbose {
				log.Printf("checktype %v not in filter\n", checktypeFilename)
			}
			// The checktypesFilter is not empty and the current checktype is not there,
			// so skip it.
			continue
		}

		// The contents of the checktype file are the assets and its options.
		checktypeFileContent, err := cli.ReadLines(filepath.Join(dir, checktypeFilename))
		if err != nil {
			return nil, err
		}

		// The first line of the checktype file, contains the default options
		// for the checktype.
		if len(checktypeFileContent) < 1 {
			return nil, fmt.Errorf("checktype %v doesn't contain at least its options", checktypeFilename)
		}

		// If no filters or no options specified, use the options defined in the checktype file.
		defaultOpts := checktypeFileContent[0]
		if verbose {
			log.Printf("default options for checktype %v in file: %v\n", checktypeFilename, defaultOpts)
		}

		if ok && checktype.DefaultOptions != "" {
			// Use the options defined in the checktypesFilter if they exist.
			defaultOpts = checktype.DefaultOptions
			if verbose {
				log.Printf("default options for checktype %v defined in filter: %v\n", checktypeFilename, defaultOpts)
			}
		}

		assetLines := checktypeFileContent[1:]
		if verbose {
			log.Printf("assets in vcdir for checktype %v: %v\n", checktypeFilename, assetLines)
		}

		var assets []cli.Asset

		for _, line := range assetLines {
			var target, opts string

			// An asset may have specific options.
			params := strings.SplitN(line, " ", 2)

			if len(params) < 1 {
				return nil, fmt.Errorf("malformed target in line: %v", line)
			}

			target = params[0]

			if len(blacklistFilter) > 0 && findTarget(target, blacklistFilter) {
				if verbose {
					log.Printf("asset %v for checktype %v in blacklist filter\n", target, checktypeFilename)
				}
				// The blacklistFilter is not empty and the current target IS there,
				// so skip it.
				continue
			}

			if len(targetsFilter) > 0 && !findTarget(target, targetsFilter) {
				if verbose {
					log.Printf("asset %v for checktype %v not in filter\n", target, checktypeFilename)
				}
				// The targetsFilter is not empty and the current target is not there,
				// so skip it.
				continue
			}

			// By default, the previously defined options of the checktype will be used.
			opts = defaultOpts

			if len(params) >= 2 && (!ok || checktype.DefaultOptions == "") {
				// Use the options specified for the asset only when they exist in
				// the file and, or when a checktypesFilter has not been specified,
				// or when even being specified there were no options in the filter.
				opts = params[1]
				if verbose {
					log.Printf("options for checktype %v and asset %v defined in file: %v\n", checktypeFilename, target, opts)
				}
			}
			if verbose {
				log.Printf("options for checktype %v and asset %v that will be used: %v\n", checktypeFilename, target, opts)
			}

			asset := cli.Asset{Target: target, Options: opts}
			assets = append(assets, asset)

		}
		checks[checktypeFilename] = assets
	}

	return checks, nil
}

// findChecktype find a Checktype by its name in a slice of Checktype.
// In addition to the checktype, it returns a bool indicating whether a
// checktype has been found or not.
func findChecktype(name string, checktypes []Checktype) (Checktype, bool) {
	for _, checktype := range checktypes {
		if checktype.Name == name {
			return checktype, true
		}
	}
	return Checktype{}, false
}

func findTarget(name string, targets []string) bool {
	for _, target := range targets {
		if target == name {
			return true
		}
	}
	return false
}
