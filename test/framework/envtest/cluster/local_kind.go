// +build test

package cluster

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"path/filepath"
	"time"

	"sigs.k8s.io/kind/pkg/cluster"

	"github.com/rancher/harvester/test/framework/envtest/cluster/finder"
	"github.com/rancher/harvester/test/framework/envtest/cluster/logs"
	"github.com/rancher/harvester/test/framework/fuzz"
)

// LocalKindCluster specifies the configurable parameters to launch a local kubernetes-sigs/kind cluster.
type LocalKindCluster struct {
	// Specify the exported ingress http port for running cluster,
	// configure in "KIND_EXPORT_INGRESS_HTTP_PORT" env,
	// default is created randomly.
	ExportIngressHTTPPort int
	// Specify the exported ingress https port for running cluster,
	// configure in "KIND_EXPORT_INGRESS_HTTPS_PORT" env,
	// default is created randomly.
	ExportIngressHTTPSPort int
	// Specify the image for running cluster,
	// configure in "KIND_IMAGE" env,
	// default is "kindest/node:v1.18.2".
	Image string
	// Specify the name of cluster,
	// configure in "KIND_CLUSTER_NAME" env,
	// default is "harvester".
	ClusterName string
	// Specify the amount of control-plane nodes,
	// configure in "KIND_CONTROL_PLANES" env,
	// default is "1".
	ControlPlanes int
	// Specify the amount of worker nodes,
	// configure in "KIND_WORKERS" env,
	// default is "3".
	Workers int
	// Specify the wait timeout for bringing up cluster,
	// configure in "KIND_WAIT_TIMEOUT" env,
	// default is "10m".
	WaitTimeout time.Duration
	// Specify the path of preset cluster configuration,
	// configure in "KIND_CLUSTER_CONFIG_PATH" env.
	ClusterConfigPath string
}

func (c LocalKindCluster) Startup(output io.Writer) error {
	var logger = logs.NewLogger(output, 0)
	var provider = cluster.NewProvider(
		cluster.ProviderWithLogger(logger),
	)

	// check if the cluster is existed
	var existed, err = isClusterExisted(provider, c.ClusterName)
	if err != nil {
		return err
	}

	// remove the existed cluster
	if existed {
		err = provider.Delete(c.ClusterName, "")
		if err != nil {
			return fmt.Errorf("failed to clean the previous cluster, %v", err)
		}
	}

	// create configuration
	var configOption cluster.CreateOption
	if c.ClusterConfigPath == "" {
		var config, err = c.generateConfiguration()
		if err != nil {
			return err
		}
		logger.V(0).Info(string(config))
		configOption = cluster.CreateWithRawConfig(config)
	} else {
		var configPath, err = filepath.Abs(c.ClusterConfigPath)
		if err != nil {
			return fmt.Errorf("failed to load cluster config from path %s, %v", c.ClusterConfigPath, err)
		}
		configOption = cluster.CreateWithConfigFile(configPath)
	}

	// create cluster
	err = provider.Create(
		c.ClusterName,
		configOption,
		cluster.CreateWithNodeImage(c.Image),
		cluster.CreateWithWaitForReady(c.WaitTimeout),
	)
	if err != nil {
		return fmt.Errorf("failed to startup, %v", err)
	}
	return nil
}

func (c LocalKindCluster) Cleanup(output io.Writer) error {
	var logger = logs.NewLogger(output, 0)
	var provider = cluster.NewProvider(
		cluster.ProviderWithLogger(logger),
	)

	// check if the cluster is existed
	var existed, err = isClusterExisted(provider, c.ClusterName)
	if err != nil {
		return err
	}

	if !existed {
		return nil
	}

	err = provider.Delete(c.ClusterName, "")
	if err != nil {
		return fmt.Errorf("failed to clean, %v", err)
	}
	return nil
}

func (c LocalKindCluster) GetKind() string {
	return "kind"
}

func (c LocalKindCluster) String() string {
	return fmt.Sprintf("Name: %s, Kind: %s, Image: %s", c.ClusterName, c.GetKind(), c.Image)
}

func (c LocalKindCluster) generateConfiguration() ([]byte, error) {
	var tpText = `---
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
networking:
  apiServerAddress: "0.0.0.0"
nodes:
  - role: control-plane
    kubeadmConfigPatches:
    - |
      kind: InitConfiguration
      nodeRegistration:
        kubeletExtraArgs:
          node-labels: "ingress-ready=true"
    extraPortMappings:
    - containerPort: 80
      hostPort: {{ .ExportIngressHTTPPort }}
      protocol: TCP
    - containerPort: 443
      hostPort: {{ .ExportIngressHTTPSPort }}
      protocol: TCP
{{- range (intRange .ControlPlanes) }}
  - role: control-plane
{{- end }}
{{- range (intRange .Workers) }}
  - role: worker
{{- end }}
---
`
	var tpFuncs = template.FuncMap{
		"intRange": func(size int) []int {
			return make([]int, size)
		},
	}
	var tp, err = template.New("harvester").Funcs(tpFuncs).Parse(tpText)
	if err != nil {
		return nil, fmt.Errorf("failed to parse configuration template, %v", err)
	}

	var cp = c
	if cp.ExportIngressHTTPSPort == 0 || cp.ExportIngressHTTPPort == 0 {
		var ports, err = fuzz.FreePorts(2)
		if err != nil {
			return nil, fmt.Errorf("failed to generate free ports in local, %v", err)
		}
		cp.ExportIngressHTTPPort = ports[0]
		cp.ExportIngressHTTPSPort = ports[1]
	}
	cp.ControlPlanes--
	var output bytes.Buffer
	err = tp.Execute(&output, cp)
	if err != nil {
		return nil, fmt.Errorf("failed to generate configuration, %v", err)
	}
	return output.Bytes(), nil
}

func isClusterExisted(provider *cluster.Provider, clusterName string) (bool, error) {
	var clusters, err = provider.List()
	if err != nil {
		return false, fmt.Errorf("failed to list all local clusters, %v", err)
	}

	for _, cls := range clusters {
		if cls == clusterName {
			return true, nil
		}
	}
	return false, nil
}

func DefaultLocalKindCluster() LocalKindCluster {
	var envFinder = finder.NewEnvFinder("kind")
	return LocalKindCluster{
		ExportIngressHTTPPort:  envFinder.GetInt("exportIngressHttpPort", 0),
		ExportIngressHTTPSPort: envFinder.GetInt("exportIngressHttpsPort", 0),
		Image:                  envFinder.Get("image", "kindest/node:v1.18.2"),
		ClusterName:            envFinder.Get("clusterName", "harvester"),
		ControlPlanes:          envFinder.GetInt("controlPlanes", 1),
		Workers:                envFinder.GetInt("workers", 3),
		WaitTimeout:            envFinder.GetDuration("waitTimeout", 10*time.Minute),
		ClusterConfigPath:      envFinder.Get("clusterConfigPath", ""),
	}
}
