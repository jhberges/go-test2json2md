# go-test2json2md

Small utility to produce test report in markdown format, primarily intended for GitHub Actions.

Uses the test2json event output from `go test --json` and generates a summary.

## Usage

```shell
 $ gotest2json2md <test2json.json> 
```

Example:
```shell
$ go test ./... -json > my-test-report.json
$ gotest2jsonmd my-test-report.json
$ head my-test-report.json.md
# Test report

Of 122 tests run there were 120 passes, 1 fails, 1 skips.



# PACKAGE github.com/octocat/super-software/v2/cmd

<details>
  <summary>Console output</summary>

```

See output in action <https://github.com/jhberges/go-test2json2md/actions/workflows/build.yaml> for example.

Inspired by the awesome [gotestsum](https://github.com/gotestyourself/gotestsum) tool.
