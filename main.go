package main

import (
    "bytes"
    "fmt"
	"github.com/spf13/cobra"
    "io/ioutil"
    "os"
	"strings"
	"text/template"
)

const TemplateOptions = "missingkey=error"

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
        outputFilename, err := flags.GetString("output")
        if err != nil {
            return err
        }


        data, err := buildDataMap(vars)
        if err != nil {
            return err
        }
        inputContents, err := ioutil.ReadFile(input)
        if err != nil {
            return err
        }
        outputContents, err := renderTemplate(string(inputContents), data)
        if err != nil {
            return err
        }
        os.WriteFile(outputFilename, outputContents.Bytes(), 0644)
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

func makeTemplate(input string) (*template.Template, error) {
	return template.New("").Option(TemplateOptions).Parse(input)
}

func renderTemplate(input string, data map[string]string) (*bytes.Buffer, error) {
    tmpl, err := makeTemplate(input)
    if err != nil {
        return nil, err
    }
    buffer := &bytes.Buffer{}
    err = tmpl.Execute(buffer, data)
    if err != nil {
        return nil, err
    }
    //zs := buffer.String()
    return buffer, nil
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
	err = command.MarkFlagRequired("output")
	if err != nil {
		return err
	}

	return nil
}
