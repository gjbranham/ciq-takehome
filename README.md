# CIQ Takehome Assessment

This repository implements a command-line utility for "querying" information contained in comma-separated server logs. 

### Clone

`git clone github.com/gjbranham/ciq-takehome`

### Build

`cd ciq-takehome && make build`

### Test

`make test`

## How to use

Run the command specifying a server log file. Must be comma-separated.

`./bin/server-info -f ./log/server_log.csv -all`

This will return a count of all server accesses contained in the log.

## Optional flags

**-u** filter by username

**-t** filter by date

**-gt** filter by file size greater than target

**-lt** filter by file size less than target