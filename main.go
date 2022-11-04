package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
	"text/template"
)

const TEMPLATE_OPTIONS = "missingkey=error"

var templaterCmd = &cobra.Command{
	Use:   "templater",
	Short: "A tool to substitute variables in templates",
	Long:  "A command line tool that substitutes variables in a template file, with error handling",
	RunE: func(cmd *cobra.Command, args []string) error {
		flags := cmd.Flags()
		input, err := flags.GetString("input")
		if err != nil {
			return err
		}
		vars, err := flags.GetStringArray("var")
		if err != nil {
			return err
		}

		data, err := buildDataMap(vars)
		fmt.Printf("input: %v data: %v\n", input, data)
		doStuff()
		return nil
	},
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

func makeTemplate(templateFilename string) (*template.Template, error) {
	b, err := os.ReadFile(templateFilename)
	if err != nil {
		return nil, err
	}
	return template.New("").Option(TEMPLATE_OPTIONS).Parse(string(b))
}

func renderTemplate(inputFilename string, outputFilename string, data map[string]string) error {
	tmpl, err := ytmakeTemplate(inputFilename)
}

func doStuff() {
	data := make(map[string]string)
	data["a"] = "foo"
	data["b"] = "bar"
	data["c"] = "baz"

	tmpl, err := template.New("").Option("missingkey=error").Parse("The value of a is {{.a}} the value of b is {{.c}}\n")
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
	os.Stdout.Sync()
	os.Stdout.Close()

}

func run() error {
	err := configureFlags(templaterCmd)
	if err != nil {
		return err
	}
	err = templaterCmd.Execute()
	if err != nil {
		return err
	}

	return nil
}

func main() {
	err := run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}
}

func configureFlags(command *cobra.Command) error {
	flags := command.Flags()
	flags.StringP("input", "i", "", "input filename")
	err := command.MarkFlagRequired("input")
	if err != nil {
		return err
	}

	flags.StringArrayP("var", "v", []string{}, "key=value, can be repeated")

	flags.StringP("output", "o", "", "output filename")
	//err = command.MarkFlagRequired("output")
	//if err != nil {
	//	return err
	//}

	return nil
}
