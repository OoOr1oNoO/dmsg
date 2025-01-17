package skyenv

// Constants for new default services.
const (
	ServiceConfAddr     = "http://conf.skywire.skycoin.com"
	TpDiscAddr          = "http://tpd.skywire.skycoin.com"
	DmsgDiscAddr        = "http://dmsgd.skywire.skycoin.com"
	ServiceDiscAddr     = "http://sd.skycoin.com"
	RouteFinderAddr     = "http://rf.skywire.skycoin.com"
	UptimeTrackerAddr   = "http://ut.skywire.skycoin.com"
	AddressResolverAddr = "http://ar.skywire.skycoin.com"
	SetupPK             = "0324579f003e6b4048bae2def4365e634d8e0e3054a20fc7af49daf2a179658557"
	NetworkMonitorPKs   = ""
)

// Constants for testing deployment.
const (
	TestServiceConfAddr     = "http://conf.skywire.dev"
	TestTpDiscAddr          = "http://tpd.skywire.dev"
	TestDmsgDiscAddr        = "http://dmsgd.skywire.dev"
	TestServiceDiscAddr     = "http://sd.skywire.dev"
	TestRouteFinderAddr     = "http://rf.skywire.dev"
	TestUptimeTrackerAddr   = "http://ut.skywire.dev"
	TestAddressResolverAddr = "http://ar.skywire.dev"
	TestSetupPK             = "026c2a3e92d6253c5abd71a42628db6fca9dd9aa037ab6f4e3a31108558dfd87cf"
	TestNetworkMonitorPKs   = "0218905f5d9079bab0b62985a05bd162623b193e948e17e7b719133f2c60b92093"
)

// GetStunServers gives back deafault Stun Servers
func GetStunServers() []string {
	return []string{
		"172.104.188.139:3478",
		"172.104.59.235:3478",
		"172.104.183.187:3478",
		"139.162.54.63:3478",
		"172.105.115.97:3478",
		"172.104.188.39:3478",
		"172.104.188.140:3478",
		"172.104.40.88:3478",
	}
}
