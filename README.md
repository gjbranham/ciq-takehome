# CIQ Takehome Assessment



This repository implements a command-line utility for filtering server access data contained in comma-separated server logs. The data should be structured like so:

| timestamp                    | username | operation | size |
| ---------------------------- | -------- | --------- | ---- |
| Sun Apr 12 22:10:38 UTC 2020 | sarah94  | download  | 34   |
| Sun Apr 12 22:35:06 UTC 2020 | Maia86   | download  | 75   |

## Usage
**Note**: go must be installed on the host machine for building and testing

### Clone

`git clone https://github.com/gjbranham/ciq-takehome.git`

### Build

`cd ciq-takehome/ && make build`

### Test

`make test`

## How to use

Run the executable in your terminal of choice after building as follows. Providing a comma-separated source file is mandatory:

`./bin/server-info -f ./log/server_log.csv -all`

This will return a count of all server accesses contained in the log.

## Optional flags

**-v** verbose output. Prints all matching lines

**-u** filter by username

**-d** filter by date

**-gt** filter by file size greater than target

**-lt** filter by file size less than target

## Notes and ideas for improvements

- More robust timestamp handling
- Filter by upload or download (easy to implement)
- More robust Makefile (fully qualified paths)
- Get test coverage near 100% by checking more error branches
- Expand username matching

Current test coverage is pretty good though!

```
coverage: 91.9% of statements
ok      github.com/gjbranham/ciq-takehome/filter        0.765s  coverage: 91.9% of statements
github.com/gjbranham/ciq-takehome/args/args.go:21:      ProcessArgs             68.4%
github.com/gjbranham/ciq-takehome/args/args.go:51:      validateArgs            88.9%
github.com/gjbranham/ciq-takehome/csv/csv.go:16:        ReadCsv                 92.3%
github.com/gjbranham/ciq-takehome/filter/filter.go:10:  FilterData              86.7%
github.com/gjbranham/ciq-takehome/filter/filter.go:38:  filterByUsername        100.0%
github.com/gjbranham/ciq-takehome/filter/filter.go:48:  filterByDate            88.9%
github.com/gjbranham/ciq-takehome/filter/filter.go:65:  filterBySize            100.0%
total:                                                  (statements)            85.9%
```
