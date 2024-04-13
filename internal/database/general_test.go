package database

import (
	"fmt"
	"testing"
)

func TestGetUsersOrExecutorsLastId(t *testing.T) {
	isUser := false
	id, _ := GetUsersOrExecutorsLastId(isUser)
	if id < 0 {
		t.Fatalf("Expected id >= 0, but got %d", id)
	} else {
		fmt.Println(id)
	}
}

func TestChangeUsersOrExecutorsLastId(t *testing.T) {
	isUser := true
	success := ChangeUsersOrExecutorsLastId(1, isUser)
	if success != nil {
		t.Fatal("Could not update metadata", success)
	} else {
		fmt.Println("OK")
	}
}
