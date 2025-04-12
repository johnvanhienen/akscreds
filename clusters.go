package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
	"sync"
)

type Cluster struct {
	Name             string `json:"name"`
	ResourceGroup    string `json:"resourceGroup"`
	SubscriptionName string
}

func getClustersPerSubscription(subscriptionName string, ch chan<- []Cluster) {
	cmd := exec.Command("az", "aks", "list", "--subscription", subscriptionName)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()

	if err != nil {
		fmt.Errorf("could not retrieve clusters, error: %s", err)
		ch <- []Cluster{}
		return
	}

	var clusters []Cluster
	json.Unmarshal([]byte(out.String()), &clusters)
	for i := range clusters {
		clusters[i].SubscriptionName = subscriptionName
	}
	ch <- clusters
}

func getClusters(subscriptions []Subscription) (result []Cluster) {
	clusters := make(chan []Cluster)
	var wg sync.WaitGroup

	for _, subscription := range subscriptions {
		wg.Add(1)
		go func(subName string) {
			defer wg.Done()
			getClustersPerSubscription(subName, clusters)
		}(subscription.Name)
	}
	go func() {
		wg.Wait()
		close(clusters)
	}()

	for cluster := range clusters {
		result = append(result, cluster...)
	}

	return
}

func getCredentialsClusters(clusters []Cluster, kubeConfigLocation string) {
	var wg sync.WaitGroup
	credentialsCh := make(chan error)

	for _, cluster := range clusters {
		wg.Add(1)
		go func(cluster Cluster) {
			defer wg.Done()
			getCredential(cluster.Name, cluster.ResourceGroup, cluster.SubscriptionName, kubeConfigLocation, credentialsCh)
		}(cluster)
	}

	go func() {
		wg.Wait()
		close(credentialsCh)
	}()

	for err := range credentialsCh {
		if err != nil {
			fmt.Println(err)
		}
	}
}
