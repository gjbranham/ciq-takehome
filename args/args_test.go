package args

import (
	"os"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestInvalidCmdLineArgs(t *testing.T) {
	type test struct {
		args []string
	}

	tests := []test{
		{args: []string{"testapp"}},
		{args: []string{"testapp", "-f", ""}},
		{args: []string{"testapp", "-f", "./somefile", "-gt", "50", "-lt", "25"}},
		{args: []string{"testapp", "-f", "./somefile", "-d", "2020-04-12"}},
		{args: []string{"testapp", "-f", "./somefile", "-d", "2020/04/12"}},
	}

	for _, tt := range tests {
		t.Run(strings.Join(tt.args, " "), func(t *testing.T) {
			os.Args = tt.args
			_, err := ProcessArgs(os.Args[0], os.Args[1:])
			if err == nil {
				t.Errorf("Wanted err for invalid args, got nil")
			}
		})
	}
}
