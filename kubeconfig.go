package main

import (
    "fmt"
    "gopkg.in/yaml.v3"
    "io/ioutil"
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

func rewriteClusterEndpoint(clusterName string, kubeconfigLocation string, newServerEndpoint string) error {
    kubeConfig := Kubeconfig{}
    kubeConfig, err := readKubeConfig(kubeconfigLocation)
    if err != nil {
        return fmt.Errorf("could not read kubeConfig file, error: %s", err)
    }

    for i, cluster := range kubeConfig.Clusters {
        if cluster.Name == clusterName {
            kubeConfig.Clusters[i].Cluster.Server = newServerEndpoint
        }
    }
    d, err := yaml.Marshal(&kubeConfig)
    if err != nil {
        fmt.Errorf("error: %v", err)
    }

    err = ioutil.WriteFile(kubeconfigLocation, d, 0644)
    if err != nil {
        return fmt.Errorf("could not write to kubeConfig file, error: %s", err)
    }

    return nil
}

func readKubeConfig(file string) (Kubeconfig, error) {
    kubeconfig := Kubeconfig{}
    data, err := ioutil.ReadFile(file)
    if err != nil {
        return Kubeconfig{}, fmt.Errorf("could not read kubeconfig file, error: %s", err)
    }

    err = yaml.Unmarshal(data, &kubeconfig)
    if err != nil {
        return Kubeconfig{}, fmt.Errorf("could not parse yaml from kubeconfig, error: %s", err)
    }

    return kubeconfig, nil
}
