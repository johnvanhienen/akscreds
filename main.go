package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
)

var (
	version    = "0.3.1"
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

	// Define flags pointing to the config struct fields
	versionFlag := flag.Bool("v", config.version, "Displays the version number of Akscreds and Go.")
	subFilterFlag := flag.String("s", config.subscriptionFilter, "Filter clusters by subscription name (regEx).")
	kubeConfigFlag := flag.String("f", config.kubeConfigLocation, "Directory to store kubeconfig files.")

	// Add custom help text
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "\n%s is a tool to retrieve credentials for all AKS clusters.\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nExample usage:\n")
		fmt.Fprintf(os.Stderr, "  %s -s \"dev|prod\" -f ~/.kube/configs/\n", os.Args[0])
	}

	flag.Parse()
	
	// Update config with flag values after parsing
	config.version = *versionFlag
	config.subscriptionFilter = *subFilterFlag
	config.kubeConfigLocation = *kubeConfigFlag

	// Make sure kubeConfigLocation ends with slash if provided
	if config.kubeConfigLocation != "" && !strings.HasSuffix(config.kubeConfigLocation, "/") {
		config.kubeConfigLocation += "/"
	}

	if config.version {
		fmt.Println(versionStr)
		os.Exit(0)
	}

	account := Account{}
	_, err := account.login()
	if err != nil {
		log.Fatalln(err)
	}

	// TODO: Execute cleanup after retrieving all the credentials. Should also only cleanup unreachable clusters, unless specified via arg.
	// eg: generate a list of clusters that are not already in the kubeconfig location and subtract those from the clusters being retrieved
	// with getClusters()
	cleanupCredentials(config.kubeConfigLocation)

	subscriptions, _ := account.getSubscriptions(config)
	clusters := getClusters(subscriptions)
	// TODO: Only retrieve clusters that are not already in the kubeconfig location
	getCredentialsClusters(clusters, config.kubeConfigLocation)
}
