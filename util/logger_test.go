package util_test

import (
	"go-rest-echo/util"
	"testing"
)

func TestLogInfo(t *testing.T) {
	actual := util.LogInfo("info")

	if actual != nil {
		t.Error("Return must be nil")
	}
}

func TestLogError(t *testing.T) {
	actual := util.LogError("error")

	if actual != nil {
		t.Error("Return must be nil")
	}
}
