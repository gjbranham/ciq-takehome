package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gjbranham/ciq-takehome/args"
	c "github.com/gjbranham/ciq-takehome/csv"
	f "github.com/gjbranham/ciq-takehome/filter"
)

func main() {
	// process cmd-line args
	args, err := args.ProcessArgs(os.Args[0], os.Args[1:])
	if err == flag.ErrHelp {
		os.Exit(2)
	} else if err != nil {
		fmt.Printf("Failed to parse command-line arguments: %v\n", err)
		os.Exit(1)
	}

	// open target log file
	fo, err := os.Open(args.SourceFile)
	if err != nil {
		fmt.Printf("Failed to open provided log file '%v': %v\n", args.SourceFile, err)
		os.Exit(1)
	}
	defer fo.Close()

	// read csv data, passing in file object (dependency injection - we can substitute a buffer in unit tests)
	allData, err := c.ReadCsv(fo)
	if err != nil {
		fmt.Printf("Failed to read provided log file '%v': %v\n", args.SourceFile, err)
		os.Exit(1)
	} else if len(allData) == 0 {
		fmt.Println("No data found in provided log file")
		os.Exit(0)
	}

	// apply filters
	filteredData, err := f.FilterData(allData, args)
	if err != nil {
		fmt.Printf("Failed to filter server data: %v\n", err)
		os.Exit(1)
	}

	// done
	if args.Verbose {
		for _, item := range filteredData {
			fmt.Printf("%v,%v,%v,%v\n", item.Timestamp, item.Username, item.Operation, item.Size)
		}
	}
	fmt.Printf("Number of server access entries after filtering: %v\n", len(filteredData))
}
