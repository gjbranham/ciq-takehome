# CIQ Takehome Assessment

This repository implements a command-line utility for filtering server access data contained in comma-separated server logs.

### Clone

`git clone https://github.com/gjbranham/ciq-takehome.git`

### Build

`cd ciq-takehome && make build`

### Test

`make test`

## How to use

Run the executable in your terminal of choice after building. Providing a comma-separated source file is mandatory:

`./bin/server-info -f ./log/server_log.csv -all`

This will return a count of all server accesses contained in the log.

## Optional flags

**-v** verbose output. Prints all matching lines

**-u** filter by username

**-t** filter by date

**-gt** filter by file size greater than target

**-lt** filter by file size less than target
