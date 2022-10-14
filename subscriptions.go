package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
)

type Subscription struct {
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

func retrieveSubscriptionIds() ([]string, error) {
	cmd := exec.Command("az", "account", "list")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()

	if err != nil {
		return nil, fmt.Errorf("could not retrieve subscriptions, error: %s", err)
	}

	stdout := out.String()
	var subscriptions []Subscription
	var subscriptionIds []string

	json.Unmarshal([]byte(stdout), &subscriptions)

	for _, subscription := range subscriptions {
		subscriptionIds = append(subscriptionIds, subscription.Name)
	}

	return subscriptionIds, nil
}
