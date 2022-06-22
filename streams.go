package scandaloriantypes

const (
	discoveryStream      = "discovery.requests"
	reverseDNSStream     = "reversedns.requests"
	applicationStream    = "application.requests"
	hostEntrichmenStream = "hostinfo.requests"
)

func GetDiscoveryStream() string {
	return discoveryStream
}

func GetReverseDNSStream() string {
	return reverseDNSStream
}

func GetApplicationStream() string {
	return applicationStream
}

func GetHostInfoStream() string {
	return hostEntrichmenStream
}
