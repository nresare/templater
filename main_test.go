package main

import (
	"bytes"
	"testing"
)

func Test_render_template(t *testing.T) {
	data := map[string]string{
		"a": "x",
	}
	output := bytes.Buffer{}
	err := renderTemplate("value of a: {{.a}}", "filename.tmpl", data, &output)

	if err != nil {
		t.Error(err)
	}
	expected := "value of a: x"
	if output.String() != expected {
		t.Errorf("expected '%s' but found '%s'", expected, output.String())
	}
}

func Test_render_template_unknown_key(t *testing.T) {
	empty := map[string]string{}
	err := renderTemplate("{{.nonexistent}}", "filename.tmpl", empty, &bytes.Buffer{})
	expected := "template: filename.tmpl:1:2: executing \"filename.tmpl\" at <.nonexistent>: map has no entry for key \"nonexistent\""
	if err.Error() != expected {
		t.Errorf("expected '%s' but found '%s'", expected, err.Error())
	}
}
