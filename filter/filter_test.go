package filter

import (
	"os"
	"strings"
	"testing"

	"github.com/gjbranham/ciq-takehome/args"
	"github.com/gjbranham/ciq-takehome/csv"
)

func TestMain(m *testing.M) {
	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestVariousArgCombinations(t *testing.T) {
	type test struct {
		args       []string
		inputData  []csv.AccessInfo
		matchCount int
	}

	tests := []test{
		{
			args: []string{"testapp", "-f", "dummyCsv"}, inputData: []csv.AccessInfo{
				{"Sun Apr 26 03:49:17 UTC 2020", "jordonGriff", "upload", 60},
				{"Sun Apr 12 22:10:38 UTC 2020", "sarah94", "download", 34},
				{"Mon Apr 13 01:44:53 UTC 2020", "rosannaM", "download", 83},
			}, matchCount: 3,
		},
		{
			args: []string{"testapp", "-f", "dummyCsv", "-all"}, inputData: []csv.AccessInfo{
				{"Sun Apr 26 03:49:17 UTC 2020", "jordonGriff", "upload", 60},
				{"Sun Apr 12 22:10:38 UTC 2020", "sarah94", "download", 34},
				{"Mon Apr 13 01:44:53 UTC 2020", "rosannaM", "download", 83},
			}, matchCount: 3,
		},
		{
			args: []string{"testapp", "-f", "dummyCsv", "-u", "jordonGriff"}, inputData: []csv.AccessInfo{
				{"Sun Apr 26 03:49:17 UTC 2020", "jordonGriff", "upload", 60},
				{"Sun Apr 12 22:10:38 UTC 2020", "sarah94", "download", 34},
				{"Mon Apr 13 01:44:53 UTC 2020", "rosannaM", "download", 83},
			}, matchCount: 1,
		},
		{
			args: []string{"testapp", "-f", "dummyCsv", "-d", "26/04/2020"}, inputData: []csv.AccessInfo{
				{"Sun Apr 26 03:49:17 UTC 2020", "jordonGriff", "upload", 60},
				{"Sun Apr 12 22:10:38 UTC 2020", "sarah94", "download", 34},
				{"Mon Apr 13 01:44:53 UTC 2020", "rosannaM", "download", 83},
			}, matchCount: 1,
		},
		{
			args: []string{"testapp", "-f", "dummyCsv", "-gt", "82"}, inputData: []csv.AccessInfo{
				{"Sun Apr 26 03:49:17 UTC 2020", "jordonGriff", "upload", 60},
				{"Sun Apr 12 22:10:38 UTC 2020", "sarah94", "download", 34},
				{"Mon Apr 13 01:44:53 UTC 2020", "rosannaM", "download", 83},
			}, matchCount: 1,
		},
		{
			args: []string{"testapp", "-f", "dummyCsv", "-lt", "35"}, inputData: []csv.AccessInfo{
				{"Sun Apr 26 03:49:17 UTC 2020", "jordonGriff", "upload", 60},
				{"Sun Apr 12 22:10:38 UTC 2020", "sarah94", "download", 34},
				{"Mon Apr 13 01:44:53 UTC 2020", "rosannaM", "download", 83},
			}, matchCount: 1,
		},
		{
			args: []string{"testapp", "-f", "dummyCsv", "-gt", "59", "-lt", "61"}, inputData: []csv.AccessInfo{
				{"Sun Apr 26 03:49:17 UTC 2020", "jordonGriff", "upload", 60},
				{"Sun Apr 12 22:10:38 UTC 2020", "sarah94", "download", 34},
				{"Mon Apr 13 01:44:53 UTC 2020", "rosannaM", "download", 83},
			}, matchCount: 1,
		},
		{
			args: []string{"testapp", "-f", "dummyCsv", "-u", "jordonGriff", "-gt", "59", "-lt", "61", "-all"}, inputData: []csv.AccessInfo{
				{"Sun Apr 26 03:49:17 UTC 2020", "jordonGriff", "upload", 60},
				{"Sun Apr 12 22:10:38 UTC 2020", "sarah94", "download", 34},
				{"Mon Apr 13 01:44:53 UTC 2020", "rosannaM", "download", 83},
			}, matchCount: 3,
		},
	}

	for _, tt := range tests {
		t.Run(strings.Join(tt.args, " "), func(t *testing.T) {
			os.Args = tt.args
			args, _ := args.ProcessArgs(os.Args[0], os.Args[1:])
			filtered, _ := FilterData(tt.inputData, *args)

			if len(filtered) != tt.matchCount {
				t.Errorf("Number of matches did not align\nWanted: %v\nGot   : %v\n", tt.matchCount, len(filtered))
			}
		})
	}
}
