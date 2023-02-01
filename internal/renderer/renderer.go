package renderer

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"text/template"

	"github.com/jhberges/go-test2json2md/v2/internal/aggregator"
)

func To(out, tmpl string, agg *aggregator.TestEventAggregator) error {
	fmt.Printf("Using template %s to write %s\n", tmpl, out)
	t, err := template.New(path.Base(tmpl)).ParseFiles(tmpl)
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
