package main

import (
	"testing"
)

var (
	args                         = []string{"apple/service3"}
	argsExpired                  = []string{"apple/service2"}
	argsCorruptedDate            = []string{"apple/service1"}
	argsNoData                   = []string{""}
	argsErrorServiceName         = []string{"apple"}
	argsErrorNotAllowedCharacter = []string{"-apple/service1"}
)

func TestRoot(t *testing.T) {

	// normal message and not expired license file
	err := Root(args)
	if err != nil {
		t.Errorf(err.Error())
	}

	// corrupted end date/notification date in license file
	err = Root(argsExpired)
	if err != nil {
		t.Errorf(err.Error())
	}

	// corrupted end date in license file
	err = Root(argsCorruptedDate)
	if err == nil {
		t.Errorf("App error: not implemented corrupted data check up")
	}

	// not completed service name
	err = Root(argsErrorServiceName)
	if err == nil {
		t.Errorf("App error: not implemented service not exists command")
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
