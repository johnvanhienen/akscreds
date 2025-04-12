package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

type Kubeconfig struct {
	APIVersion string `yaml:"apiVersion"`
	Clusters   []struct {
		Cluster struct {
			CertificateAuthorityData string `yaml:"certificate-authority-data"`
			Server                   string `yaml:"server"`
		} `yaml:"cluster"`
		Name string `yaml:"name"`
	} `yaml:"clusters"`
	Contexts []struct {
		Context struct {
			Cluster string `yaml:"cluster"`
			User    string `yaml:"user"`
		} `yaml:"context"`
		Name string `yaml:"name"`
	} `yaml:"contexts"`
	CurrentContext string `yaml:"current-context"`
	Kind           string `yaml:"kind"`
	Preferences    struct {
	} `yaml:"preferences"`
	Users []struct {
		Name string `yaml:"name"`
		User struct {
			AuthProvider struct {
				Config struct {
					ApiserverID string `yaml:"apiserver-id"`
					ClientID    string `yaml:"client-id"`
					ConfigMode  string `yaml:"config-mode"`
					Environment string `yaml:"environment"`
					TenantID    string `yaml:"tenant-id"`
				} `yaml:"config"`
				Name string `yaml:"name"`
			} `yaml:"auth-provider"`
		} `yaml:"user"`
	} `yaml:"users"`
}

func getCredential(clusterName string, resourceGroup string, subscriptionName string, file string, ch chan<- error) {
	uniqueFileName := fmt.Sprintf("%sah-%s.yaml", file, clusterName)
	cmd := exec.Command("az", "aks", "get-credentials",
		"--name", clusterName,
		"--resource-group", resourceGroup,
		"--file", uniqueFileName,
		"--subscription", subscriptionName,
		"--overwrite-existing",
	)
	err := cmd.Run()
	if err != nil {
		fmt.Printf("az aks get-credentials --name %s --resource-group %s --file %s --subscription %s\n", clusterName, resourceGroup, uniqueFileName, subscriptionName)
		ch <- fmt.Errorf("could not save the credentials of %s clusterName to the kubeconfig file %s, error: %s", clusterName, uniqueFileName, err)
	} else {
		fmt.Printf("Succesfully saved credentials for %s to %s\n", clusterName, uniqueFileName)
		ch <- nil
	}
}

func cleanupCredentials(kubeConfigLocation string) {
	command := fmt.Sprintf("rm %sah*.yaml", kubeConfigLocation)
	cmd := exec.Command("bash", "-c", command)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}
