package main

import (
    "flag"
    "fmt"
    "log"
    "os"
    "runtime"
)

var (
    version    = "0.2.0"
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
    allTenantFlag := flag.Bool("A", false, "Retrieve credentials from all tenants which are shown in 'az account list'")
    rewriteServerEndpointOpt := flag.String("r", "", "Rewrite server endpoint in from saved config to specified value. ")
    kubeConfigLocationOpt := flag.String("f", config.kubeConfigLocation, "Kubeconfig file to update.")
    flag.Parse()
    if *versionFlag {
        fmt.Println(versionStr)
        os.Exit(0)
    }

    loggedInTenantId, err := loginAccount()
    if err != nil {
        log.Fatalln(err)
    }

    subscriptions := Subscriptions{}
    subscriptions, err = retrieveAllSubscriptions()
    if err != nil {
        log.Fatalln(err)
    }

    var subscriptionNames []string
    if *allTenantFlag {
        subscriptionNames = subscriptions.getAllSubscriptionNames()
    } else {
        subscriptionNames = subscriptions.getAllSubscriptionNamesByTenantIds([]string{loggedInTenantId})
    }

    for _, subscriptionName := range subscriptionNames {
        err := setActiveSubscription(subscriptionName)
        if err != nil {
            log.Fatalln(err)
        }
        fmt.Printf("Setting subscription name to: \"%s\"\n", subscriptionName)

        clusters, err := retrieveClusters()
        if err != nil {
            log.Println(err)
        }

        for _, cluster := range clusters {
            saveKubeConfig(cluster.Name, cluster.ResourceGroup, *kubeConfigLocationOpt)
            if *rewriteServerEndpointOpt != "" {
                rewriteClusterEndpoint(cluster.Name, *kubeConfigLocationOpt, *rewriteServerEndpointOpt)
            }
        }
    }
}
