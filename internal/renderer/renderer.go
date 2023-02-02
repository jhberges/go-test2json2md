package renderer

import (
	"bufio"
	_ "embed"
	"fmt"
	"os"
	"path"
	"text/template"

	"github.com/jhberges/go-test2json2md/v2/internal/aggregator"
)

var (
	// NB No space between comments and "go:..." below!
	//go:embed template.tmpl
	report_template string
)

func To(out, tmpl string, agg *aggregator.TestEventAggregator) error {
	t, err := template.New(path.Base(tmpl)).Parse(report_template)
	if err != nil {
		fmt.Printf("No parsy\n")
		return err
	}
	f, err := os.Create(out)
	if err != nil {
		fmt.Printf("No creaty\n")
		return err
	}
	wr := bufio.NewWriter(f)
	defer wr.Flush()
	return t.Execute(wr, agg)
}
