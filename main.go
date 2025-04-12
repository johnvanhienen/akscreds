package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
)

var (
	version    = "0.3.0"
	goVersion  = runtime.Version()
	versionStr = fmt.Sprintf("Akscreds version %v\n%v", version, goVersion)
)

func (c *Config) fillDefaults() {
	c.kubeConfigLocation = fmt.Sprintf("%s/.kube/configs/", os.Getenv("HOME"))
	c.subscriptionFilter = "kaas"
	c.version = false
}

func main() {
	config := Config{}
	config.fillDefaults()

	flag.Bool("v", config.version, "Displays the version number of Akscreds and Go.")
	flag.String("s", config.subscriptionFilter, "Filter clusters by subscription name (regEx).")
	flag.String("f", config.kubeConfigLocation, "Kubeconfig file to update.")

	flag.Parse()
	if config.version {
		fmt.Println(versionStr)
		os.Exit(0)
	}

	account := Account{}
	_, err := account.login()
	if err != nil {
		log.Fatalln(err)
	}

	cleanupCredentials(config.kubeConfigLocation)

	subscriptions, _ := account.getSubscriptions(config)
	clusters := getClusters(subscriptions)
	getCredentialsClusters(clusters, config.kubeConfigLocation)

	cmd := exec.Command("kubelogin", "convert-kubeconfig", "-l", "azurecli")
	err = cmd.Run()
}
