package subnetgenerator

// incrementIP - increment the host octet
func (sa *SubnetAddresses) incrementIP() {
	for j := len(sa.ip) - 1; j >= 0; j-- {
		sa.ip[j]++
		if sa.ip[j] > 0 {
			break
		}
	}
}
