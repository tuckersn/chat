package api

import (
	"testing"

	"github.com/tuckersn/chatbackend/util"
)

func Test_Func_ValidateNotePath_NonExistant(t *testing.T) {
	notePath := "/nonexistant"
	notePathStatus := ValidateNotePath(notePath)
	if notePathStatus != util.FILE_NOT_FOUND {
		t.Errorf("Expected %d, got %d", util.FILE_NOT_FOUND, notePathStatus)
	}
}
