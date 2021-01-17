package helper_test

import (
	"go-rest-echo/helper"
	"testing"
)

func TestGenerateUUID(t *testing.T) {
	uuid := helper.GenerateUUID()

	if "" == uuid {
		t.Errorf("Expected `%v`, but got `%v`", "uuid string", uuid)
	}

	if len(uuid) != 36 {
		t.Errorf("Expected length = `%v`, but got `%v`", 36, len(uuid))
	}
}
