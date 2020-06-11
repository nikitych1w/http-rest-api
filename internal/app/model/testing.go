package model

import "testing"

// TestUser ...
func TestUser(t *testing.T) *User {
	return &User{
		Email:    "user@examole.org",
		Password: "password",
	}
}
