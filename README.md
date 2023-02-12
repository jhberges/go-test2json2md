# go-test2json2md

---

Small utility to produce test report in markdown format, primarily intended for GitHub Actions.

Uses the test2json event output from `go test --json` and generates a summary.

## Usage

```shell
 $ go-test2json2md <test2json.json> 
```

See output in action <https://github.com/jhberges/go-test2json2md/actions/workflows/build.yaml> for example.

Inspired by the awesome [gotestsum](https://github.com/gotestyourself/gotestsum) tool.
