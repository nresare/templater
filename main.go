package main

import (
	"fmt"
	"io"
	"os"
	"text/template"
)

const TemplateOptions = "missingkey=error"

func execute(config *config) error {
	inputContents, err := os.ReadFile(config.inputFilename)
	if err != nil {
		return err
	}
	output, err := os.Create(config.outputFilename)
	if err != nil {
		return err
	}
	err = renderTemplate(string(inputContents), config.inputFilename, config.vars, output)
	if err != nil {
		return err
	}
	return nil
}

func renderTemplate(input string, templateName string, data map[string]string, output io.Writer) error {
	tmpl, err := template.New(templateName).Option(TemplateOptions).Parse(input)
	if err != nil {
		return err
	}

	err = tmpl.Execute(output, data)
	if err != nil {
		return err
	}
	return nil
}

func run() error {
	args, err := parseArgs()
	if err != nil {
		return err
	}
	err = execute(args)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	err := run()
	if err != nil {
		// explicitly ignore error writing to Stderr, as there is nowhere else to signal the error
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}
}
