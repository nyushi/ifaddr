package ifaddr

import (
	"fmt"
	"net"
)

// IPTypeQuery represents query for ip address type
type IPTypeQuery struct {
	IsLoopback  bool
	IsLinklocal bool
	IsMulticast bool
	IsIPv4      bool
	IsIPv6      bool
}

// ListInterfaces returns interface list
func ListInterfaces() ([]string, error) {
	ifs, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	names := make([]string, len(ifs))
	for i, v := range ifs {
		names[i] = v.Name
	}
	return names, nil
}

// ListAddresses returns address list
func ListAddresses(q IPTypeQuery) ([]string, error) {
	names, err := ListInterfaces()
	if err != nil {
		return nil, fmt.Errorf("failed toget interface list: %s", err)
	}
	list := []string{}
	for _, n := range names {
		addrs, err := GetAddress(n, q)
		if err != nil {
			return nil, fmt.Errorf("failed to get address for %s: %s", n, err)
		}
		list = append(list, addrs...)
	}
	return list, nil
}

// GetAddress returns address for interface
func GetAddress(ifname string, query IPTypeQuery) ([]string, error) {
	i, err := net.InterfaceByName(ifname)
	if err != nil {
		return nil, fmt.Errorf("failed to get interface %s: %s", ifname, err)
	}
	addrs, err := i.Addrs()
	if err != nil {
		return nil, fmt.Errorf("failed to get address for %s: %s", ifname, err)
	}
	list := []string{}
	for _, a := range addrs {
		ipnet := a.(*net.IPNet)
		ip := ipnet.IP
		if ip.IsLoopback() && !query.IsLoopback {
			continue
		}
		if ip.IsLinkLocalUnicast() && !query.IsLinklocal {
			continue
		}
		if ip.IsLinkLocalMulticast() && !query.IsLinklocal {
			continue
		}
		if ip.IsMulticast() && !query.IsMulticast {
			continue
		}
		if ip.To4() != nil && !query.IsIPv4 {
			continue
		}
		if ip.To4() == nil && !query.IsIPv6 {
			continue
		}
		list = append(list, ip.String())
	}
	return list, nil
}
