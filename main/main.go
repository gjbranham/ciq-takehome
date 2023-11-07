package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/gjbranham/ciq-takehome/args"
	c "github.com/gjbranham/ciq-takehome/csv"
	f "github.com/gjbranham/ciq-takehome/filter"
)

func main() {
	args, err := args.ProcessArgs(os.Args[0], os.Args[1:])
	if err == flag.ErrHelp {
		os.Exit(2)
	} else if err != nil {
		fmt.Printf("Failed to parse command-line arguments: %v\n", err)
		os.Exit(1)
	}

	absPath, err := filepath.Abs(args.SourceFile)
	if err != nil {
		fmt.Printf("Failed to resolve absolute path for '%v': %v\n", args.SourceFile, err)
		os.Exit(1)
	}

	fo, err := os.Open(absPath)
	if err != nil {
		fmt.Printf("Failed to open provided log file '%v': %v\n", args.SourceFile, err)
	}
	defer fo.Close()

	allData, err := c.ReadCsv(fo)
	if err != nil {
		fmt.Printf("Failed to read provided log file '%v': %v\n", args.SourceFile, err)
		os.Exit(1)
	} else if len(allData) == 0 {
		fmt.Println("No data found in provided log file")
		os.Exit(0)
	}

	finalData, err := f.FilterData(allData, *args)
	if err != nil {
		fmt.Printf("Failed to filter server data: %v\n", err)
		os.Exit(1)
	}

	if args.Verbose {
		for _, item := range finalData {
			fmt.Printf("%v,%v,%v,%v\n", item.Timestamp, item.Username, item.Operation, item.Size)
		}
	}
	fmt.Printf("Number of server access entries after filtering: %v\n", len(finalData))

}
