package scandaloriantypes

type Scan interface {
	SetDefaults(scan *ScanMetaData)
	GetStream() string
}

//ScanRequest object instructing system on how to scan.
type ScanRequest struct {
	Host            string           `json:"host,omitempty"`
	PortScan        *PortScan        `json:"port_scan,omitempty"`
	ApplicationScan *ApplicationScan `json:"application_scan,omitempty"`
}

//Scan structure to send to message queue for scanning
type ScanMetaData struct {
	IP        string `json:"ip"`
	ScanID    string `json:"scan_id"`
	RequestID string `json:"request_id"`
	Stream    string `json:"-"`
}

//Top level object to define a port scan
type PortScan struct {
	ScanMetaData
	Run                    bool     `json:"run"`
	PPS                    int      `json:"pps,omitempty"` //Set rate limiter value
	Ports                  []int    `json:"ports,omitempty"`
	TopTen                 bool     `json:"top_ten,omitempty"`
	TopHundred             bool     `json:"top_hundred,omitempty"`
	TopThousand            bool     `json:"top_thousand,omitempty"`
	HostScanTimeoutSeconds int      `json:"host_scan_timeout_seconds"`
	Errors                 []string `json:"errors,omitempty"`
}

func (ps *PortScan) SetDefaults(scan *ScanMetaData) {
	ps.IP = scan.IP
	ps.ScanID = scan.ScanID
	ps.RequestID = scan.RequestID
}

func (ps *PortScan) GetStream() string {
	return discoveryStream
}

//Top Level Object to define application level scanning
type ApplicationScan struct {
	ScanMetaData
	Run         bool      `json:"run"`
	Flags       BaseFlags `json:"flags"`
	UDPFlags    UDPFlags  `json:"udp_flags,omitempty"`
	HttpOptions *Http     `json:"http_options,omitempty"`
}

func (ps *ApplicationScan) SetDefaults(scan *ScanMetaData) {
	ps.IP = scan.IP
	ps.ScanID = scan.ScanID
	ps.RequestID = scan.RequestID
}

func (ps *ApplicationScan) GetStream() string {
	return applicationStream
}
