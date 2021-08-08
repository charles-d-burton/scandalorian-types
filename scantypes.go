package scandaloriantypes

//ScanRequest object instructing system on how to scan.
type ScanRequest struct {
	Address         string           `json:"address,omitempty"`
	Host            string           `json:"host,omitempty"`
	PortScan        *PortScan        `json:"port_scan,omitempty"`
	ApplicationScan *ApplicationScan `json:"application_scan,omitempty"`
}

//Scan structure to send to message queue for scanning
type Scan struct {
	IP        string `json:"ip"`
	ScanID    string `json:"scan_id"`
	RequestID string `json:"request_id"`
	Stream    string `json:"-"`
}

//Top level object to define a port scan
type PortScan struct {
	Scan
	Options *PortScanOptions `json:"options,omitempty"`
}

//PortScanOptions optional parameters to set for a scan
type PortScanOptions struct {
	TopTen      bool     `json:"top_ten,omitempty"`
	TopHundred  bool     `json:"top_hundred,omitempty"`
	TopThousand bool     `json:"top_thousand,omitempty"`
	PPS         int      `json:"pps,omitempty"` //Set rate limiter value
	Ports       []string `json:"ports,omitempty"`
}

func (ps *PortScan) SetDefaults(scan *Scan) {
	ps.IP = scan.IP
	ps.ScanID = scan.ScanID
	ps.RequestID = scan.RequestID
}

//Top Level Object to define application level scanning
type ApplicationScan struct {
	Scan
	Options *ApplicationScanOptions `json:"options,omitempty"`
}

type ApplicationScanOptions struct {
	Flags       BaseFlags `json:"flags"`
	UDPFlags    UDPFlags  `json:"udp_flags,omitempty"`
	HttpOptions *Http     `json:"http_options,omitempty"`
}

func (ps *ApplicationScan) SetDefaults(scan *Scan) {
	ps.IP = scan.IP
	ps.ScanID = scan.ScanID
	ps.RequestID = scan.RequestID
}
