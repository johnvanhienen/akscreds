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

func retrieveSubscriptionNames(tenantId string) ([]string, error) {
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
		if subscription.HomeTenantID == tenantId {
			subscriptionNames = append(subscriptionNames, subscription.Name)
		}
	}

	subscriptionNames = removeBlacklistedSubscriptions(subscriptionNames)
	return subscriptionNames, nil
}

func removeBlacklistedSubscriptions(subscriptions []string) []string {
	// TODO: Set blacklist via parameter/config
	subscriptionBlacklist := []string{"Visual Studio Professional Subscription"}
	for _, entry := range subscriptionBlacklist {
		subscriptions = removeItemFromSlice(subscriptions, entry)
	}
	return subscriptions
}

func setActiveSubscription(subscriptionName string) error {
	cmd := exec.Command("az", "account", "set", "--subscription", fmt.Sprintf("%s", subscriptionName))
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("could not complete login, error: %s", err)
	}
	return nil
}
