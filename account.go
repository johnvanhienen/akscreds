package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
)

type Account struct {
	CloudName        string        `json:"cloudName"`
	HomeTenantID     string        `json:"homeTenantId"`
	ID               string        `json:"id"`
	IsDefault        bool          `json:"isDefault"`
	ManagedByTenants []interface{} `json:"managedByTenants"`
	Name             string        `json:"name"`
	State            string        `json:"state"`
	TenantID         string        `json:"tenantId"`
	User             struct {
		Name string `json:"name"`
		Type string `json:"type"`
	} `json:"user"`
}

func loginAccount() (string, error) {
	cmd := exec.Command("az", "login")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("could not complete login, error: %s", err)
	}

	stdout := out.String()
	var accounts []Account
	err = json.Unmarshal([]byte(stdout), &accounts)
	if err != nil {
		fmt.Errorf("could not parse the data to an Account struct, error: %s", err)
	}

	// When logging in, you should only retrieve subscriptions from a single tenant. Just grabbing the first tenantId
	// should be sufficient.
	firstTenantId := accounts[0].HomeTenantID
	return firstTenantId, nil
}
