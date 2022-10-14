package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	//err := login()
	//if err != nil {
	//	log.Fatalln(err)
	//}
	subscriptionNames, err := retrieveSubscriptionNames()
	if err != nil {
		log.Fatalln(err)
	}

	for _, subscriptionName := range subscriptionNames {
		err := setSubscription(subscriptionName)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("Setting subscription name to: \"%s\"\n", subscriptionName)

		//clusters, err := retrieveClusters()
		//if err != nil {
		//	log.Fatalln(err)
		//}
		//
		//for _, cluster := range clusters {
		//	fmt.Println(cluster.Name)
		//	fmt.Println(cluster.ResourceGroup)
		//}

	}
}

func login() error {
	// Maybe rework to login via the azure sdk instead of relying on the azure cli.
	cmd := exec.Command("az", "login")
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("could not complete login, error: %s", err)
	}
	return nil
}

func setSubscription(subscriptionName string) error {
	cmd := exec.Command("az", "account", "set", "--subscription", fmt.Sprintf("%s", subscriptionName))
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("could not complete login, error: %s", err)
	}
	return nil
}
