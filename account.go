package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
)

func loginAccount() (string, error) {
	cmd := exec.Command("az", "login")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("could not complete login, error: %s", err)
	}

	stdout := out.String()
	var subscriptions []Subscription
	err = json.Unmarshal([]byte(stdout), &subscriptions)
	if err != nil {
		return "", fmt.Errorf("could not parse the data to an Subscription struct, error: %s", err)
	}

	// When logging in, you should only retrieve subscriptions from a single tenant. Just grabbing the first tenantId
	// should be sufficient.
	firstTenantId := subscriptions[0].TenantID
	return firstTenantId, nil
}
