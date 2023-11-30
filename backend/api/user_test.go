package api

import (
	"testing"

	_ "github.com/lib/pq"
	"github.com/tuckersn/chatbackend/db"
	"github.com/tuckersn/chatbackend/util"
)

func Test_AccountLifeCycle(t *testing.T) {

	logger, err := util.TestSetup("../.env")
	if err != nil {
		t.Errorf("Error setting up test: %s", err)
		return
	}

	logger.Println("Creating user")
	createdUser, err := UserCreate("test_lifecycle", "test_lifecycle")
	if err != nil {
		panic(err)
	}
	if createdUser.Username != "test_lifecycle" {
		t.Errorf("Expected username test_lifecycle, got %s", createdUser.Username)
		return
	}
	if createdUser.DisplayName != "test_lifecycle" {
		t.Errorf("Expected display name test_lifecycle, got %s", createdUser.DisplayName)
		return
	}
	if createdUser.Admin != false {
		t.Errorf("Expected admin false, got %t", createdUser.Admin)
		return
	}
	logger.Println("Getting user")
	getUser, err := db.GetUser(createdUser.Username)
	if err != nil {
		panic(err)
	}
	if getUser.Username != "test_lifecycle" {
		t.Errorf("Expected username test_lifecycle, got %s", getUser.Username)
		return
	}
	if getUser.DisplayName != "test_lifecycle" {
		t.Errorf("Expected display name test_lifecycle, got %s", getUser.DisplayName)
		return
	}
	if getUser.Admin != false {
		t.Errorf("Expected admin false, got %t", getUser.Admin)
		return
	}
	logger.Println("Deleting user")
	err = db.DeleteUser(createdUser.Username)
	if err != nil {
		t.Errorf("Error deleting user: %s", err)
		return
	}
}
