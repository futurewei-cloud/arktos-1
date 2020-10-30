/*
Copyright The Kubernetes Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package versioned

import (
	discovery "k8s.io/client-go/discovery"
	rest "k8s.io/client-go/rest"
	flowcontrol "k8s.io/client-go/util/flowcontrol"
	globalschedulerv1 "k8s.io/kubernetes/globalscheduler/pkg/apis/cluster/client/clientset/versioned/typed/cluster/v1"
)

type Interface interface {
	Discovery() discovery.DiscoveryInterface
	GlobalschedulerV1() globalschedulerv1.GlobalschedulerV1Interface
}

// Clientset contains the clients for groups. Each group has exactly one
// version included in a Clientset.
type Clientset struct {
	*discovery.DiscoveryClient
	globalschedulerV1 *globalschedulerv1.GlobalschedulerV1Client
}

// GlobalschedulerV1 retrieves the GlobalschedulerV1Client
func (c *Clientset) GlobalschedulerV1() globalschedulerv1.GlobalschedulerV1Interface {
	return c.globalschedulerV1
}

// Discovery retrieves the DiscoveryClient
func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	if c == nil {
		return nil
	}
	return c.DiscoveryClient
}

// NewForConfig creates a new Clientset for the given config.
func NewForConfig(c *rest.Config) (*Clientset, error) {
	configShallowCopy := *c
	for _, configCopy := range configShallowCopy.GetAllConfigs() {
		if configCopy.RateLimiter == nil && configCopy.QPS > 0 {
			configCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configCopy.QPS, configCopy.Burst)
		}
	}
	var cs Clientset
	var err error
	cs.globalschedulerV1, err = globalschedulerv1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	cs.DiscoveryClient, err = discovery.NewDiscoveryClientForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	return &cs, nil
}

// NewForConfigOrDie creates a new Clientset for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *Clientset {
	var cs Clientset
	cs.globalschedulerV1 = globalschedulerv1.NewForConfigOrDie(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClientForConfigOrDie(c)
	return &cs
}

// New creates a new Clientset for the given RESTClient.
func New(c rest.Interface) *Clientset {
	var cs Clientset
	cs.globalschedulerV1 = globalschedulerv1.New(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClient(c)
	return &cs
}
