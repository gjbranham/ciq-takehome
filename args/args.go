package args

import (
	"errors"
	"flag"
	"os"
	"path/filepath"
	"time"
)

type Arguments struct {
	AllResults      bool
	Verbose         bool
	SourceFile      string
	Username        string
	Date            string
	GreaterThanSize int
	LessThanSize    int
}

func ProcessArgs(exeName string, sysArgs []string) (*Arguments, error) {
	flags := flag.NewFlagSet(exeName, flag.ContinueOnError)
	flag.CommandLine.SetOutput(os.Stderr)

	var args Arguments
	flags.StringVar(&args.SourceFile, "f", "", "Source csv server log file. Required argument")
	flags.BoolVar(&args.Verbose, "v", false, "Show verbose output")
	flags.BoolVar(&args.AllResults, "all", false, "Return all server accesses. Supersedes all other flags")
	flags.StringVar(&args.Username, "u", "", "Filter results by username")
	flags.StringVar(&args.Date, "d", "", "Filter results by date. DD/MM/YYYY format required")
	flags.IntVar(&args.GreaterThanSize, "gt", 0, "Filter results by size greater than target")
	flags.IntVar(&args.LessThanSize, "lt", 0, "Filter results by size less than target")

	if err := flags.Parse(sysArgs); err != nil {
		return nil, err
	}

	if err := validateArgs(args); err != nil {
		return nil, err
	}

	absPath, err := filepath.Abs(args.SourceFile)
	if err != nil {
		return nil, err
	}
	args.SourceFile = absPath

	return &args, nil
}

func validateArgs(args Arguments) error {
	if args.SourceFile == "" {
		return errors.New("source file cannot be empty")
	}

	if (args.GreaterThanSize > 0 && args.LessThanSize > 0) && (args.GreaterThanSize > args.LessThanSize) {
		return errors.New("invalid combination of greaterThan and lessThan values")
	}

	if args.Date != "" {
		_, err := time.Parse("02/01/2006", args.Date)
		if err != nil {
			return err
		}
	}
	return nil
}
