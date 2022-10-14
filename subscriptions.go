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

func retrieveSubscriptionNames() ([]string, error) {
	cmd := exec.Command("az", "account", "list")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()

	if err != nil {
		return nil, fmt.Errorf("could not retrieve subscriptions, error: %s", err)
	}

	stdout := out.String()
	var subscriptions []Subscription
	var subscriptionNames []string

	json.Unmarshal([]byte(stdout), &subscriptions)

	for _, subscription := range subscriptions {
		subscriptionNames = append(subscriptionNames, subscription.Name)
	}
	fmt.Println(subscriptionNames)
	subscriptionNames = scrubBlacklist(subscriptionNames)
	fmt.Println(subscriptionNames)
	return subscriptionNames, nil
}

func scrubBlacklist(subscriptions []string) []string {
	// TODO: Set blacklist via parameter/config
	subscriptionBlacklist := []string{"Visual Studio Professional Subscription"}
	for _, entry := range subscriptionBlacklist {
		subscriptions = removeItemFromSlice(subscriptions, entry)
	}
	return subscriptions
}

func removeItemFromSlice(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}
