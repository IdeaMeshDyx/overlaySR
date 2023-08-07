// Copyright 2020 Authors of Cilium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build go1.14
// +build go1.14

package ciliumAPI

import (
	"overlaysr/client/internal/pkg/data"
	"sort"
	"strings"

	"github.com/cilium/cilium/pkg/client"
)

// in fact , endpoints in cilium are usually pods
func GetEps() data.PodsMsg {
	// Connect to the default path /var/run/cilium/cilium.sock
	msg := data.PodsMsg{}
	c, err := client.NewDefaultClient()
	if err != nil {
		klog.Errorf("Create Cilium Client failed: %v", err)
		return nil
	}

	// List all endpoints
	eps, err := c.EndpointList()
	if err != nil {
		klog.Errorf("Create Cilium Client failed: %v", err)
		return nil
	}

	// Sort EPs per IDs
	sort.Slice(eps, func(i, j int) bool {
		return eps[i].ID < eps[j].ID
	})

	// Print the IPs of the endpoints
	for _, ep := range eps {

		var pod data.SinglePod
		pod.Id = ep.ID
		var v4s, v6s []string
		for _, ip := range ep.Status.Networking.Addressing {
			if ip.IPV4 != "" {
				v4s = append(v4s, ip.IPV4)
			}
			if ip.IPV6 != "" {
				v6s = append(v6s, ip.IPV6)
			}
		}
		ip4s := strings.Join(v4s, ", ")
		ip6s := strings.Join(v6s, ", ")
		if ip4s != "" {
			pod.IPv4 = ip4s
		} else {
			pod.IPv4 = "N/A"
		}
		if ip6s != "" {
			pod.IPv6 = ip6s
		} else {
			pod.IPv6 = "N/A"
		}

		pod.HostMac = ep.Status.Networking.HostMac
		pod.InterfaceName = ep.Status.Networking.InterfaceName
		pod.InterfaceIndex = ep.Status.Networking.InterfaceIndex
		pod.Mac = ep.Status.Networking.Mac

		pod.Indentity = ep.Status.Identity.ID
		pod.Labels = ep.Status.Identity.Labels
		pod.LabelsSHA256 = ep.Status.Identity.LabelsSHA256

		msg.Pods = append(msg.Pods, pod)

	}
	return msg

}
