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

// +build go1.14

package ciliumAPI

import (
	"fmt"
	"sort"
	"strings"

	"github.com/cilium/cilium/pkg/client"
)

func main() {
	// Connect to the default path /var/run/cilium/cilium.sock
	c, err := client.NewDefaultClient()
	if err != nil {
		panic(err)
	}

	// List all endpoints
	eps, err := c.EndpointList()
	if err != nil {
		panic(err)
	}

	// Sort EPs per IDs
	sort.Slice(eps, func(i, j int) bool {
		return eps[i].ID < eps[j].ID
	})

	// Print the IPs of the endpoints
	for _, ep := range eps {
		var v4s, v6s []string
		for _, ip := range ep.Status.Networking.Addressing {
			if ip.IPV4 != "" {
				v4s = append(v4s, ip.IPV4)
			}
			if ip.IPV4 != "" {
				v6s = append(v6s, ip.IPV6)
			}
		}
		ips := strings.Join(v4s, ", ")
		if ips != "" {
			fmt.Printf("EP ID %d has IP addresses: %s\n", ep.ID, ips)
		} else {
			fmt.Printf("EP ID %d does not have an IP address\n", ep.ID)
		}

		// test below

		// No.1 print networking issues
		fmt.Printf("Addressing: %v\n", ep.Status.Networking.Addressing)
		fmt.Printf("HostAddressing: %v\n", ep.Status.Networking.HostAddressing)
		fmt.Printf("HostMac: %s\n", ep.Status.Networking.HostMac)
		fmt.Printf("InterfaceIndex: %d\n", ep.Status.Networking.InterfaceIndex)
		fmt.Printf("InterfaceName: %s\n", ep.Status.Networking.InterfaceName)
		fmt.Printf("Mac: %s\n", ep.Status.Networking.Mac)

		for _, ip := range ep.Status.Networking.Addressing {
			fmt.Printf("ip: %v\n", ip)
			if ip.IPV4 != "" {
				fmt.Printf("IPv4: %s\n", ip.IPV4)
			} else {
				fmt.Printf("No IPv4\n")
			}
			if ip.IPV6 != "" {
				fmt.Printf("IPv6: %s\n", ip.IPV6)
			} else {
				fmt.Printf("No IPv6\n")
			}
			// print uuid
			fmt.Printf("UUID: %s\n", ip.IPV4ExpirationUUID)
		}
		fmt.Printf("\n=====================\n")

		// No.2 print identity
		// fmt.Printf("Identity: %v\n", ep.Status.Identity)

	}

}
