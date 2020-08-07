/*
Copyright 2020 Rancher Labs, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by main. DO NOT EDIT.

package v1alpha3

import (
	"github.com/rancher/harvester/pkg/generated/clientset/versioned/scheme"
	rest "k8s.io/client-go/rest"
	v1alpha3 "kubevirt.io/client-go/api/v1alpha3"
)

type KubevirtV1alpha3Interface interface {
	RESTClient() rest.Interface
	KubeVirtsGetter
	VirtualMachinesGetter
	VirtualMachineInstancesGetter
	VirtualMachineInstanceMigrationsGetter
	VirtualMachineInstancePresetsGetter
	VirtualMachineInstanceReplicaSetsGetter
}

// KubevirtV1alpha3Client is used to interact with features provided by the kubevirt.io group.
type KubevirtV1alpha3Client struct {
	restClient rest.Interface
}

func (c *KubevirtV1alpha3Client) KubeVirts(namespace string) KubeVirtInterface {
	return newKubeVirts(c, namespace)
}

func (c *KubevirtV1alpha3Client) VirtualMachines(namespace string) VirtualMachineInterface {
	return newVirtualMachines(c, namespace)
}

func (c *KubevirtV1alpha3Client) VirtualMachineInstances(namespace string) VirtualMachineInstanceInterface {
	return newVirtualMachineInstances(c, namespace)
}

func (c *KubevirtV1alpha3Client) VirtualMachineInstanceMigrations(namespace string) VirtualMachineInstanceMigrationInterface {
	return newVirtualMachineInstanceMigrations(c, namespace)
}

func (c *KubevirtV1alpha3Client) VirtualMachineInstancePresets(namespace string) VirtualMachineInstancePresetInterface {
	return newVirtualMachineInstancePresets(c, namespace)
}

func (c *KubevirtV1alpha3Client) VirtualMachineInstanceReplicaSets(namespace string) VirtualMachineInstanceReplicaSetInterface {
	return newVirtualMachineInstanceReplicaSets(c, namespace)
}

// NewForConfig creates a new KubevirtV1alpha3Client for the given config.
func NewForConfig(c *rest.Config) (*KubevirtV1alpha3Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &KubevirtV1alpha3Client{client}, nil
}

// NewForConfigOrDie creates a new KubevirtV1alpha3Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *KubevirtV1alpha3Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new KubevirtV1alpha3Client for the given RESTClient.
func New(c rest.Interface) *KubevirtV1alpha3Client {
	return &KubevirtV1alpha3Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha3.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *KubevirtV1alpha3Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
