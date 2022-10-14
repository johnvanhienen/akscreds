package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	err := login()
	if err != nil {
		log.Fatalln(err)
	}

	subscriptionIds, err := retrieveSubscriptionIds()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(subscriptionIds)

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
