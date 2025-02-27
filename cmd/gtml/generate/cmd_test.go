package generate

import (
	"embed"
	"log/slog"
	"strings"
	"testing"
)

//go:embed testdata
var testdata embed.FS

func TestStart(t *testing.T) {
	entries, err := testdata.ReadDir("testdata")
	if err != nil {
		t.Fatal(err)
	}
	for _, e := range entries {
		name := e.Name()
		if !strings.HasSuffix(name, ".html") {
			continue
		}
		name = strings.TrimSuffix(name, ".html")

		t.Run(name, func(t *testing.T) {
			in := readTestData(t, name+".html")
			out := readTestData(t, name+".go")

			r := strings.NewReader(in)
			var w strings.Builder
			err := Run(newLogger(t), r, &w)

			if err != nil {
				t.Error(err)
			}
			if out != w.String() {
				t.Errorf("expected %q, got %q", out, w.String())
			}
		})
	}
}

func readTestData(t *testing.T, path string) string {
	t.Helper()

	b, err := testdata.ReadFile("testdata/" + path)
	if err != nil {
		t.Fatal(err)
	}
	return string(b)
}

func newLogger(t *testing.T) *slog.Logger {
	t.Helper()
	return slog.New(slog.NewTextHandler(&testWriter{t}, nil))
}

type testWriter struct {
	t *testing.T
}

func (t *testWriter) Write(p []byte) (n int, err error) {
	t.t.Log(string(p))
	return len(p), nil
}
