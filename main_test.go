package main

import (
	"bytes"
	"reflect"
	"testing"
)

func Test_buildDataMap_happy(t *testing.T) {
	vars := []string{"foo=bar"}
	data, err := buildDataMap(vars)
	if err != nil {
		t.Errorf("%v", err)
	}
	expected := map[string]string{"foo": "bar"}
	if !reflect.DeepEqual(expected, data) {
		t.Errorf("expected %v but received %v", expected, data)
	}
}

func Test_buildDataMap_malformed(t *testing.T) {
	vars := []string{"foobar"}
	_, err := buildDataMap(vars)
	if err == nil {
		t.Errorf("expected error")
	}
	expectedError := "failed to parse 'foobar' into a pair separated by '='"
	if err.Error() != expectedError {
		t.Errorf("Expected error '%s' but received '%s'", expectedError, err.Error())
	}
}

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
