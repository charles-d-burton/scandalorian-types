package scandaloriantypes

type CustomMeta map[string]string

type Scan interface {
	SetDefaults(scan *ScanMetaData)
	GetStream() string
}

// ScanRequest object instructing system on how to scan.
type ScanRequest struct {
	Host            string           `json:"host"`
	Company         string           `json:"company"`
	FQDN            string           `json:"fqdn,omitempty"`
	PortScan        *PortScan        `json:"port_scan,omitempty"`
	ApplicationScan *ApplicationScan `json:"application_scan,omitempty"`
	CustomMeta      CustomMeta       `json:"custom_meta,omitempty"`
}

// Scan structure to send to message queue for scanning
type ScanMetaData struct {
	IP         string     `json:"ip"`
	Domain     string     `json:"domain"`
	Requestor  string     `json:"requestor"`
	RequestID  string     `json:"request_id"`
	CustomMeta CustomMeta `json:"custom_meta,omitempty"`
	Stream     string     `json:"-"`
}

func (sr *ScanRequest) GenerateScanMeta(requestID string) ScanMetaData {
	var meta ScanMetaData
	meta.IP = sr.Host
	meta.Domain = sr.FQDN
	meta.RequestID = requestID
	meta.CustomMeta = sr.CustomMeta
	return meta
}

// Top level object to define a port scan
type PortScan struct {
	ScanMetaData
	Run                    bool     `json:"run"`
	PPS                    int      `json:"pps,omitempty"` // Set rate limiter value
	Ports                  []int    `json:"ports,omitempty"`
	TopTen                 bool     `json:"top_ten,omitempty"`
	TopHundred             bool     `json:"top_hundred,omitempty"`
	TopThousand            bool     `json:"top_thousand,omitempty"`
	HostScanTimeoutSeconds int      `json:"host_scan_timeout_seconds"`
	Errors                 []string `json:"errors,omitempty"`
}

func (ps *PortScan) GetStream() string {
	return discoveryStream
}

// Top Level Object to define application level scanning
type ApplicationScan struct {
	ScanMetaData
	Run         bool      `json:"run"`
	Flags       BaseFlags `json:"flags"`
	UDPFlags    UDPFlags  `json:"udp_flags,omitempty"`
	HttpOptions *Http     `json:"http_options,omitempty"`
}

func (ps *ApplicationScan) GetStream() string {
	return applicationStream
}

//Top level struct to define RSA Cracker attempt
