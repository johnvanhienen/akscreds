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

type Subscriptions struct {
	subscriptions []Subscription
}

func retrieveAllSubscriptions() (Subscriptions, error) {
	cmd := exec.Command("az", "account", "list")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()

	if err != nil {
		return Subscriptions{}, fmt.Errorf("could not retrieve subscriptions, error: %s", err)
	}

	var subscriptions Subscriptions
	stdout := out.String()
	err = json.Unmarshal([]byte(stdout), &subscriptions.subscriptions)
	if err != nil {
		return Subscriptions{}, err
	}

	return subscriptions, nil
}

func (s *Subscriptions) getFirstTenant() string {
	return s.subscriptions[0].TenantID
}

func (s *Subscriptions) getAllSubscriptionNames() []string {
	var subscriptionNames []string

	for _, subscription := range s.subscriptions {
		subscriptionNames = append(subscriptionNames, subscription.Name)
	}
	subscriptionNames = removeBlacklistedSubscriptions(subscriptionNames)
	return subscriptionNames
}

func (s *Subscriptions) getAllSubscriptionNamesByTenantIds(tenantIds []string) []string {
	var subscriptionNames []string

	for _, tenantId := range tenantIds {
		for _, subscription := range s.subscriptions {
			if tenantId == subscription.TenantID {
				subscriptionNames = append(subscriptionNames, subscription.Name)
			}
		}
	}
	subscriptionNames = removeBlacklistedSubscriptions(subscriptionNames)
	return subscriptionNames
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
