# templater

Templater is a small stand-alone tool that executes go templates, substituting variables provided on the command line.
The purpose is to provide an easier to use alternative to creative sed invocations with better error handling.
One context when this is useful is in scripts that execute when a VM instance or similar is started, and 
config files will need to be created automatically that contains instance specific values.

##  Installation

Templater can be installed as other go mod based software: `go install github.com/nresare/templater@latest` which
will compile and install the binary in `$GOROOT/bin`. On my system this is `$HOME/go/bin`, so for the tool to
be available in your shell you will need to make sure that your `$PATH` contains this directory.

## Usage

```bash
$ echo "The value of a is {{.a}}" > test.tmpl
$ templater --input test.tmpl --output test.txt --var a=42
$ cat test.txt
The value of a is 42
```
