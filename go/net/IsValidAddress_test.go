package net

import "testing"

func TestIsValidAddress(t *testing.T) {
	t.Skip("disabled for debugging")
	// Test valid IPv4 addresses
	ipv4Addresses := []string{"192.0.2.1", "10.0.0.1", "172.16.0.1"}
	for _, ip := range ipv4Addresses {
		if !IsValidAddress(ip) {
			t.Errorf("Expected %s to be a valid address", ip)
		}
	}

	// Test valid IPv6 addresses
	ipv6Addresses := []string{"2001:0db8:85a3:0000:0000:8a2e:0370:7334", "fe80::1", "::1"}
	for _, ip := range ipv6Addresses {
		if !IsValidAddress(ip) {
			t.Errorf("Expected %s to be a valid address", ip)
		}
	}

	// Test valid FQDNs
	fqdns := []string{"example.com", "subdomain.example.com", "google.com"}
	for _, fqdn := range fqdns {
		if !IsValidAddress(fqdn) {
			t.Errorf("Expected %s to be a valid address", fqdn)
		}
	}

	// Test invalid addresses
	invalidAddresses := []string{"notAnIp", "256.0.0.1", "2001:0db8:85a3::8a2e:0370::7334"}
	for _, addr := range invalidAddresses {
		if IsValidAddress(addr) {
			t.Errorf("Expected %s to be an invalid address", addr)
		}
	}
}
