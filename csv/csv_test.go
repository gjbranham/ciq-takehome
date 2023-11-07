package csv

import (
	"bytes"
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestCsvIsReadProperly(t *testing.T) {
	records := [][]string{
		{"Sun Apr 12 22:10:38 UTC 2020", "sarah94", "download", "34"},
		{"Sun Apr 13 22:35:06 UTC 2020", "Maia86", "download", "75"},
		{"Sun Apr 14 22:49:47 UTC 2020", "Maia86", "upload", "9"},
	}

	var buf bytes.Buffer
	buf.WriteString("timestamp,username,operation,size\n")
	for _, r := range records {
		row := strings.Join(r, ",")
		buf.WriteString(row + "\n")
	}

	actualInfo, err := ReadCsv(&buf)
	if err != nil {
		t.Fatalf("Unable to read csv data from buffer: %v\n", err)
	}

	expectedInfo := []AccessInfo{
		{"Sun Apr 12 22:10:38 UTC 2020", "sarah94", "download", 34},
		{"Sun Apr 13 22:35:06 UTC 2020", "Maia86", "download", 75},
		{"Sun Apr 14 22:49:47 UTC 2020", "Maia86", "upload", 9},
	}

	if !reflect.DeepEqual(expectedInfo, actualInfo) {
		t.Fatalf("Expected and actual differ:\nExpected: %v\nActual  : %v\n", expectedInfo, actualInfo)
	}
}
