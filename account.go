package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

type Account struct {
	TenantID string `json:"tenantId"`
}

type Subscription struct {
	Name string `json:"name"`
}

func (a *Account) login() (string, error) {
	stdout, stderr := showAccount()

	if strings.Contains(stderr, "Please run 'az login' to setup account") {
		fmt.Println("Attempting to login...")
		cmd := exec.Command("az", "login")
		err := cmd.Run()
		if err != nil {
			return "", fmt.Errorf("could not complete login, error: %s", err)
		}
		stdout, stderr = showAccount()
	}
	err := json.Unmarshal([]byte(stdout), &a)
	if err != nil {
		return "", fmt.Errorf("could not parse the data to an Account struct, error: %s", err)
	}

	fmt.Printf("Logged in with tenant ID: %s\n", a.TenantID)
	return a.TenantID, nil
}

func (a *Account) getSubscriptions(config Config) ([]Subscription, error) {
	cmd := exec.Command("az", "account", "list")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()

	if err != nil {
		return []Subscription{}, fmt.Errorf("could not retrieve subscriptions, error: %s", err)
	}

	var subscriptions []Subscription
	stdout := out.String()
	err = json.Unmarshal([]byte(stdout), &subscriptions)
	if err != nil {
		return []Subscription{}, err
	}

	if config.subscriptionFilter != "" {
		tmp := []Subscription{}
		for _, subscription := range subscriptions {
			if regexp.MustCompile(config.subscriptionFilter).MatchString(subscription.Name) {
				tmp = append(tmp, subscription)
			}
		}
		subscriptions = tmp
	}

	return subscriptions, nil
}

func showAccount() (string, string) {
	cmd := exec.Command("az", "account", "show")
	var out bytes.Buffer
	var outErr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &outErr
	cmd.Run()

	return out.String(), outErr.String()
}
