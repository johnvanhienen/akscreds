package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
)

type Cluster struct {
	AadProfile struct {
		AdminGroupObjectIDs []string    `json:"adminGroupObjectIDs"`
		AdminUsers          interface{} `json:"adminUsers"`
		ClientAppID         interface{} `json:"clientAppId"`
		EnableAzureRbac     bool        `json:"enableAzureRbac"`
		Managed             bool        `json:"managed"`
		ServerAppID         interface{} `json:"serverAppId"`
		ServerAppSecret     interface{} `json:"serverAppSecret"`
		TenantID            string      `json:"tenantId"`
	} `json:"aadProfile"`
	AddonProfiles struct {
		Azurepolicy struct {
			Config   interface{} `json:"config"`
			Enabled  bool        `json:"enabled"`
			Identity interface{} `json:"identity"`
		} `json:"azurepolicy"`
		HTTPApplicationRouting struct {
			Config   interface{} `json:"config"`
			Enabled  bool        `json:"enabled"`
			Identity interface{} `json:"identity"`
		} `json:"httpApplicationRouting"`
		KubeDashboard struct {
			Config   interface{} `json:"config"`
			Enabled  bool        `json:"enabled"`
			Identity interface{} `json:"identity"`
		} `json:"kubeDashboard"`
		Omsagent struct {
			Config struct {
				LogAnalyticsWorkspaceResourceID string `json:"logAnalyticsWorkspaceResourceID"`
			} `json:"config"`
			Enabled  bool        `json:"enabled"`
			Identity interface{} `json:"identity"`
		} `json:"omsagent"`
	} `json:"addonProfiles"`
	AgentPoolProfiles []struct {
		AvailabilityZones      []string    `json:"availabilityZones"`
		Count                  int         `json:"count"`
		CreationData           interface{} `json:"creationData"`
		EnableAutoScaling      bool        `json:"enableAutoScaling"`
		EnableEncryptionAtHost bool        `json:"enableEncryptionAtHost"`
		EnableFips             bool        `json:"enableFips"`
		EnableNodePublicIP     bool        `json:"enableNodePublicIp"`
		EnableUltraSsd         bool        `json:"enableUltraSsd"`
		GpuInstanceProfile     interface{} `json:"gpuInstanceProfile"`
		KubeletConfig          interface{} `json:"kubeletConfig"`
		KubeletDiskType        string      `json:"kubeletDiskType"`
		LinuxOsConfig          interface{} `json:"linuxOsConfig"`
		MaxCount               interface{} `json:"maxCount"`
		MaxPods                int         `json:"maxPods"`
		MinCount               interface{} `json:"minCount"`
		Mode                   string      `json:"mode"`
		Name                   string      `json:"name"`
		NodeImageVersion       string      `json:"nodeImageVersion"`
		NodeLabels             interface{} `json:"nodeLabels"`
		NodePublicIPPrefixID   interface{} `json:"nodePublicIpPrefixId"`
		NodeTaints             interface{} `json:"nodeTaints"`
		OrchestratorVersion    string      `json:"orchestratorVersion"`
		OsDiskSizeGb           int         `json:"osDiskSizeGb"`
		OsDiskType             string      `json:"osDiskType"`
		OsSku                  string      `json:"osSku"`
		OsType                 string      `json:"osType"`
		PodSubnetID            interface{} `json:"podSubnetId"`
		PowerState             struct {
			Code string `json:"code"`
		} `json:"powerState"`
		ProvisioningState         string      `json:"provisioningState"`
		ProximityPlacementGroupID interface{} `json:"proximityPlacementGroupId"`
		ScaleDownMode             interface{} `json:"scaleDownMode"`
		ScaleSetEvictionPolicy    interface{} `json:"scaleSetEvictionPolicy"`
		ScaleSetPriority          interface{} `json:"scaleSetPriority"`
		SpotMaxPrice              interface{} `json:"spotMaxPrice"`
		Tags                      interface{} `json:"tags"`
		Type                      string      `json:"type"`
		UpgradeSettings           struct {
			MaxSurge interface{} `json:"maxSurge"`
		} `json:"upgradeSettings"`
		VMSize          string      `json:"vmSize"`
		VnetSubnetID    string      `json:"vnetSubnetId"`
		WorkloadRuntime interface{} `json:"workloadRuntime"`
	} `json:"agentPoolProfiles"`
	APIServerAccessProfile struct {
		AuthorizedIPRanges             interface{} `json:"authorizedIpRanges"`
		DisableRunCommand              interface{} `json:"disableRunCommand"`
		EnablePrivateCluster           bool        `json:"enablePrivateCluster"`
		EnablePrivateClusterPublicFqdn bool        `json:"enablePrivateClusterPublicFqdn"`
		PrivateDNSZone                 string      `json:"privateDnsZone"`
	} `json:"apiServerAccessProfile"`
	AutoScalerProfile  interface{} `json:"autoScalerProfile"`
	AutoUpgradeProfile struct {
		UpgradeChannel string `json:"upgradeChannel"`
	} `json:"autoUpgradeProfile"`
	AzurePortalFqdn         string      `json:"azurePortalFqdn"`
	DisableLocalAccounts    bool        `json:"disableLocalAccounts"`
	DiskEncryptionSetID     string      `json:"diskEncryptionSetId"`
	DNSPrefix               string      `json:"dnsPrefix"`
	EnablePodSecurityPolicy interface{} `json:"enablePodSecurityPolicy"`
	EnableRbac              bool        `json:"enableRbac"`
	ExtendedLocation        interface{} `json:"extendedLocation"`
	Fqdn                    interface{} `json:"fqdn"`
	FqdnSubdomain           interface{} `json:"fqdnSubdomain"`
	HTTPProxyConfig         interface{} `json:"httpProxyConfig"`
	ID                      string      `json:"id"`
	Identity                interface{} `json:"identity"`
	IdentityProfile         interface{} `json:"identityProfile"`
	KubernetesVersion       string      `json:"kubernetesVersion"`
	LinuxProfile            struct {
		AdminUsername string `json:"adminUsername"`
		SSH           struct {
			PublicKeys []struct {
				KeyData string `json:"keyData"`
			} `json:"publicKeys"`
		} `json:"ssh"`
	} `json:"linuxProfile"`
	Location       string `json:"location"`
	MaxAgentPools  int    `json:"maxAgentPools"`
	Name           string `json:"name"`
	NetworkProfile struct {
		DNSServiceIP        string   `json:"dnsServiceIp"`
		DockerBridgeCidr    string   `json:"dockerBridgeCidr"`
		IPFamilies          []string `json:"ipFamilies"`
		LoadBalancerProfile struct {
			AllocatedOutboundPorts interface{} `json:"allocatedOutboundPorts"`
			EffectiveOutboundIPs   []struct {
				ID            string `json:"id"`
				ResourceGroup string `json:"resourceGroup"`
			} `json:"effectiveOutboundIPs"`
			EnableMultipleStandardLoadBalancers interface{} `json:"enableMultipleStandardLoadBalancers"`
			IdleTimeoutInMinutes                interface{} `json:"idleTimeoutInMinutes"`
			ManagedOutboundIPs                  struct {
				Count     int         `json:"count"`
				CountIpv6 interface{} `json:"countIpv6"`
			} `json:"managedOutboundIPs"`
			OutboundIPs        interface{} `json:"outboundIPs"`
			OutboundIPPrefixes interface{} `json:"outboundIpPrefixes"`
		} `json:"loadBalancerProfile"`
		LoadBalancerSku   string      `json:"loadBalancerSku"`
		NatGatewayProfile interface{} `json:"natGatewayProfile"`
		NetworkMode       interface{} `json:"networkMode"`
		NetworkPlugin     string      `json:"networkPlugin"`
		NetworkPolicy     string      `json:"networkPolicy"`
		OutboundType      string      `json:"outboundType"`
		PodCidr           string      `json:"podCidr"`
		PodCidrs          []string    `json:"podCidrs"`
		ServiceCidr       string      `json:"serviceCidr"`
		ServiceCidrs      []string    `json:"serviceCidrs"`
	} `json:"networkProfile"`
	NodeResourceGroup  string      `json:"nodeResourceGroup"`
	PodIdentityProfile interface{} `json:"podIdentityProfile"`
	PowerState         struct {
		Code string `json:"code"`
	} `json:"powerState"`
	PrivateFqdn          string `json:"privateFqdn"`
	PrivateLinkResources []struct {
		GroupID              string      `json:"groupId"`
		ID                   interface{} `json:"id"`
		Name                 string      `json:"name"`
		PrivateLinkServiceID string      `json:"privateLinkServiceId"`
		RequiredMembers      []string    `json:"requiredMembers"`
		Type                 string      `json:"type"`
	} `json:"privateLinkResources"`
	ProvisioningState   string `json:"provisioningState"`
	PublicNetworkAccess string `json:"publicNetworkAccess"`
	ResourceGroup       string `json:"resourceGroup"`
	SecurityProfile     struct {
		AzureDefender interface{} `json:"azureDefender"`
	} `json:"securityProfile"`
	ServicePrincipalProfile struct {
		ClientID string `json:"clientId"`
	} `json:"servicePrincipalProfile"`
	Sku struct {
		Name string `json:"name"`
		Tier string `json:"tier"`
	} `json:"sku"`
	SystemData     interface{} `json:"systemData"`
	Type           string      `json:"type"`
	WindowsProfile interface{} `json:"windowsProfile"`
}

func retrieveClusters() ([]Cluster, error) {
	cmd := exec.Command("az", "aks", "list")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()

	if err != nil {
		return []Cluster{}, fmt.Errorf("could not retrieve clusters, error: %s", err)
	}

	stdout := out.String()
	var clusters []Cluster

	json.Unmarshal([]byte(stdout), &clusters)

	return clusters, nil
}
