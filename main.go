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
	version    = "v0.1.0"
	goVersion  = runtime.Version()
	versionStr = fmt.Sprintf("Akscreds version %v\n%v", version, goVersion)
)

type Config struct {
	kubeConfigLocation string
}

func (c *Config) fillDefaults() {
	c.kubeConfigLocation = fmt.Sprintf("%s/.kube/config", os.Getenv("HOME"))
}

func main() {
	// TODO: Maybe rework all the shell commands to talk via the sdk. Removes the dependency on the 'az' binary.
	config := Config{}
	config.fillDefaults()

	versionFlag := flag.Bool("v", false, "Displays the version number of Akscreds and Go.")
	kubeConfigLocationOpt := flag.String("f", config.kubeConfigLocation, "Kubeconfig file to update.")
	flag.Parse()
	if *versionFlag {
		fmt.Println(versionStr)
		os.Exit(0)
	}

	tenantId, err := loginAccount()
	if err != nil {
		log.Fatalln(err)
	}

	subscriptionNames, err := retrieveSubscriptionNames(tenantId)
	if err != nil {
		log.Fatalln(err)
	}

	for _, subscriptionName := range subscriptionNames {
		err := setActiveSubscription(subscriptionName)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("Setting subscription name to: \"%s\"\n", subscriptionName)

		clusters, err := retrieveClusters()
		if err != nil {
			log.Fatalln(err)
		}

		for _, cluster := range clusters {
			saveKubeConfig(cluster.Name, cluster.ResourceGroup, *kubeConfigLocationOpt)
		}
	}
}

func saveKubeConfig(clusterName string, resourceGroup string, file string) error {
	cmd := exec.Command("az", "aks", "get-credentials",
		"--name", clusterName,
		"--resource-group", resourceGroup,
		"--file", file)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("could not save the credentials to the kubeconfig file %s, error: %s", file, err)
	}

	fmt.Printf("Succesfully saved credentials for %s to %s\n", clusterName, file)
	return nil
}
