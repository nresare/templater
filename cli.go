package main

import (
	"fmt"
	"github.com/spf13/pflag"
	"os"
	"strings"
)

func failWithUsage(error string) {
	_, _ = fmt.Fprintf(os.Stderr, "%s\n\n", error)
	pflag.Usage()
	os.Exit(-1)
}

type config struct {
	inputFilename  string
	outputFilename string
	vars           map[string]string
}

func parseArgs() (*config, error) {
	inputFilename := pflag.StringP("input", "i", "", "input filename")
	vars := pflag.StringArrayP("var", "v", []string{}, "key=value, can be repeated")
	outputFilename := pflag.StringP("output", "o", "", "output filename")

	pflag.Parse()

	if *inputFilename == "" {
		failWithUsage("required parameter --input/-i is missing")
	}
	if *outputFilename == "" {
		failWithUsage("required parameter --output/-o is missing")
	}

	config := &config{}
	config.outputFilename = *outputFilename
	config.inputFilename = *inputFilename

	var err error
	config.vars, err = buildDataMap(*vars)
	if err != nil {
		return nil, err
	}
	return config, nil
}

func buildDataMap(vars []string) (map[string]string, error) {
	data := make(map[string]string)
	for _, s := range vars {
		parts := strings.Split(s, "=")
		if len(parts) != 2 {
			return nil, fmt.Errorf("failed to parse '%s' into a pair separated by '='", s)
		}
		data[parts[0]] = parts[1]
	}
	return data, nil
}
