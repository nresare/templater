package main

import (
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
