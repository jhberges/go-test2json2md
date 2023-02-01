package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/jhberges/go-test2json2md/v2/internal/aggregator"
	"github.com/jhberges/go-test2json2md/v2/internal/renderer"
)

func main() {
	flag.Parse()

	source := flag.Arg(0)
	if len(source) > 0 {
		destination := flag.Arg(1)
		if len(destination) == 0 {
			destination = source + ".md"
		}
		bytes, err := ioutil.ReadFile(source)
		if err != nil {
			fmt.Println(err.Error())
		}
		data := string(bytes)
		lines := strings.Split(data, "\n")
		agg := &aggregator.TestEventAggregator{}
		for _, v := range lines {
			if len(strings.Trim(v, "\n\t ")) > 0 {
				ev := &aggregator.TestEvent{}
				err = json.Unmarshal([]byte(v), ev)
				if err != nil {
					fmt.Printf("Un parseable line: %s\n", err.Error())
				} else {
					agg.Add(ev)
				}
			}
		}
		template := "./template.tmpl"
		err = renderer.To(destination, template, agg)
		if err != nil {
			fmt.Printf("Failed to write output %s\n", err.Error())
		}
	}
}
