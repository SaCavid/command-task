package main

import (
	"testing"
)

var (
	args = []string{"apple/service1"}
	argsNoData = []string{""}
	argsErrorServiceName = []string{"apple"}
	argsErrorNotAllowedCharacter = []string{"-apple/service1"}
	//args = []string{"apple/service1"}
	//args = []string{"apple/service1"}
)

func TestRoot(t *testing.T) {

	// normal message
	err := Root(args)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	// not completed service name
	err = Root(argsErrorServiceName)
	if err == nil {
		t.Errorf("App error: service not exists command")
	}

	// not command exist
	err = Root(argsNoData)
	if err == nil {
		t.Errorf("App error: not fixed - no command entered ")
	}

	// not allowed character
	err = Root(argsErrorNotAllowedCharacter)
	if err == nil {
		t.Errorf("App error: not added character check")
	}
}
